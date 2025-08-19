use std::sync::Arc;
use tokio::sync::Semaphore;
use tokio::task::JoinHandle;

#[derive(Debug, Clone)]
pub enum WorkerType {
    Http,
    Grpc,
    Messaging,
    Metrics,
}

#[derive(Debug, thiserror::Error)]
pub enum WorkerPoolError {
    #[error("Failed to acquire semaphore permit")]
    SemaphoreAcquisitionFailed,
    #[error("Worker pool is at capacity")]
    PoolAtCapacity,
}

#[derive(Clone)]
pub struct WorkerPool {
    semaphore: Arc<Semaphore>,
    max_workers: usize,
    worker_type: WorkerType,
}

#[derive(Debug)]
pub struct PoolStats {
    pub http: PoolStat,
    pub grpc: PoolStat,
    pub messaging: PoolStat,
    pub metrics: PoolStat,
}

impl std::fmt::Display for PoolStats {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        write!(
            f,
            "Pool Stats:\n\
             HTTP: {}/{} ({:.1}% utilized)\n\
             gRPC: {}/{} ({:.1}% utilized)\n\
             Messaging: {}/{} ({:.1}% utilized)\n\
             Metrics: {}/{} ({:.1}% utilized)",
            self.http.max - self.http.available,
            self.http.max,
            self.http.utilization * 100.0,
            self.grpc.max - self.grpc.available,
            self.grpc.max,
            self.grpc.utilization * 100.0,
            self.messaging.max - self.messaging.available,
            self.messaging.max,
            self.messaging.utilization * 100.0,
            self.metrics.max - self.metrics.available,
            self.metrics.max,
            self.metrics.utilization * 100.0
        )
    }
}

#[derive(Debug)]
pub struct PoolStat {
    pub available: usize,
    pub max: usize,
    pub utilization: f64,
}

impl WorkerPool {
    pub fn new(max_workers: usize, worker_type: WorkerType) -> Self {
        Self {
            semaphore: Arc::new(Semaphore::new(max_workers)),
            max_workers,
            worker_type,
        }
    }

    pub async fn spawn<F, Fut, T>(&self, task: F) -> Result<JoinHandle<T>, WorkerPoolError>
    where
        F: FnOnce() -> Fut + Send + 'static,
        Fut: std::future::Future<Output = T> + Send + 'static,
        T: Send + 'static,
    {
        let permit = self
            .semaphore
            .clone()
            .acquire_owned()
            .await
            .map_err(|_| WorkerPoolError::SemaphoreAcquisitionFailed)?;

        let worker_type = self.worker_type.clone();
        let handle = tokio::spawn(async move {
            tracing::debug!("Starting task on {:?} worker", worker_type);
            let result = task().await;
            tracing::debug!("Completed task on {:?} worker", worker_type);
            if matches!(worker_type, WorkerType::Http)
                || matches!(worker_type, WorkerType::Grpc)
                || matches!(worker_type, WorkerType::Metrics)
            {
                permit.semaphore().add_permits(1);
            }

            if matches!(worker_type, WorkerType::Messaging) {
                permit.semaphore().add_permits(1);
                permit.forget();
            }

            result
        });

        Ok(handle)
    }

    pub fn available_workers(&self) -> usize {
        self.semaphore.available_permits()
    }

    pub fn max_workers(&self) -> usize {
        self.max_workers
    }

    pub fn utilization(&self) -> f64 {
        let available = self.available_workers();
        let used = self.max_workers - available;
        (used as f64) / (self.max_workers as f64)
    }

    pub fn worker_type(&self) -> &WorkerType {
        &self.worker_type
    }
}
