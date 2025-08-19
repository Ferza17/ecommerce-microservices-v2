use tokio::task::JoinHandle;
use crate::package::worker_pool::worker_pool::{PoolStat, PoolStats, WorkerPool, WorkerPoolError, WorkerType};

#[derive(Clone)]
pub struct TypedWorkerPool {
    pub http_pool: WorkerPool,
    pub grpc_pool: WorkerPool,
    pub messaging_pool: WorkerPool,
    pub metrics_pool: WorkerPool,
}


impl TypedWorkerPool {
    pub fn new(
        http_workers: usize,
        grpc_workers: usize,
        messaging_workers: usize,
        metrics_workers: usize,
    ) -> Self {
        Self {
            http_pool: WorkerPool::new(http_workers, WorkerType::Http),
            grpc_pool: WorkerPool::new(grpc_workers, WorkerType::Grpc),
            messaging_pool: WorkerPool::new(messaging_workers, WorkerType::Messaging),
            metrics_pool: WorkerPool::new(metrics_workers, WorkerType::Metrics),
        }
    }

    pub fn with_defaults() -> Self {
        Self::new(
            50,  // HTTP workers
            20,  // gRPC workers
            10,  // Messaging workers
            1,   // Metrics workers
        )
    }

    /// Spawn a task on the most appropriate pool based on the task type
    pub async fn spawn_task<F, Fut, T>(
        &self,
        worker_type: WorkerType,
        task: F,
    ) -> Result<JoinHandle<T>, WorkerPoolError>
    where
        F: FnOnce() -> Fut + Send + 'static,
        Fut: std::future::Future<Output = T> + Send + 'static,
        T: Send + 'static,
    {
        match worker_type {
            WorkerType::Http => self.spawn_http_task(task).await,
            WorkerType::Grpc => self.spawn_grpc_task(task).await,
            WorkerType::Messaging => self.spawn_messaging_task(task).await,
            WorkerType::Metrics => self.spawn_metrics_task(task).await,
        }
    }

    // HTTP-specific methods
    pub async fn spawn_http_task<F, Fut, T>(&self, task: F) -> Result<JoinHandle<T>, WorkerPoolError>
    where
        F: FnOnce() -> Fut + Send + 'static,
        Fut: std::future::Future<Output = T> + Send + 'static,
        T: Send + 'static,
    {
        self.http_pool.spawn(task).await
    }

    // gRPC-specific methods
    pub async fn spawn_grpc_task<F, Fut, T>(&self, task: F) -> Result<JoinHandle<T>, WorkerPoolError>
    where
        F: FnOnce() -> Fut + Send + 'static,
        Fut: std::future::Future<Output = T> + Send + 'static,
        T: Send + 'static,
    {
        self.grpc_pool.spawn(task).await
    }

    // Messaging-specific methods (RabbitMQ/Kafka)
    pub async fn spawn_messaging_task<F, Fut, T>(&self, task: F) -> Result<JoinHandle<T>, WorkerPoolError>
    where
        F: FnOnce() -> Fut + Send + 'static,
        Fut: std::future::Future<Output = T> + Send + 'static,
        T: Send + 'static,
    {
        self.messaging_pool.spawn(task).await
    }

    pub async fn spawn_rabbitmq_task<F, Fut, T>(&self, task: F) -> Result<JoinHandle<T>, WorkerPoolError>
    where
        F: FnOnce() -> Fut + Send + 'static,
        Fut: std::future::Future<Output = T> + Send + 'static,
        T: Send + 'static,
    {
        self.spawn_messaging_task(task).await
    }

    pub async fn spawn_kafka_task<F, Fut, T>(&self, task: F) -> Result<JoinHandle<T>, WorkerPoolError>
    where
        F: FnOnce() -> Fut + Send + 'static,
        Fut: std::future::Future<Output = T> + Send + 'static,
        T: Send + 'static,
    {
        self.spawn_messaging_task(task).await
    }

    // Metrics-specific methods
    pub async fn spawn_metrics_task<F, Fut, T>(&self, task: F) -> Result<JoinHandle<T>, WorkerPoolError>
    where
        F: FnOnce() -> Fut + Send + 'static,
        Fut: std::future::Future<Output = T> + Send + 'static,
        T: Send + 'static,
    {
        self.metrics_pool.spawn(task).await
    }

    // Utility methods
    pub fn get_pool_stats(&self) -> PoolStats {
        PoolStats {
            http: PoolStat {
                available: self.http_pool.available_workers(),
                max: self.http_pool.max_workers(),
                utilization: self.http_pool.utilization(),
            },
            grpc: PoolStat {
                available: self.grpc_pool.available_workers(),
                max: self.grpc_pool.max_workers(),
                utilization: self.grpc_pool.utilization(),
            },
            messaging: PoolStat {
                available: self.messaging_pool.available_workers(),
                max: self.messaging_pool.max_workers(),
                utilization: self.messaging_pool.utilization(),
            },
            metrics: PoolStat {
                available: self.metrics_pool.available_workers(),
                max: self.metrics_pool.max_workers(),
                utilization: self.metrics_pool.utilization(),
            },
        }
    }

    pub fn total_available_workers(&self) -> usize {
        self.http_pool.available_workers()
            + self.grpc_pool.available_workers()
            + self.messaging_pool.available_workers()
            + self.metrics_pool.available_workers()
    }

    pub fn total_max_workers(&self) -> usize {
        self.http_pool.max_workers()
            + self.grpc_pool.max_workers()
            + self.messaging_pool.max_workers()
            + self.metrics_pool.max_workers()
    }
}
