use std::sync::Arc;
use tokio::sync::Semaphore;
use tokio::task::JoinHandle;

pub struct WorkerPool {
    semaphore: Arc<Semaphore>,
    // max_workers: usize,
}

impl WorkerPool {
    pub fn new(max_workers: usize) -> Self {
        Self {
            semaphore: Arc::new(Semaphore::new(max_workers)),
            // max_workers,
        }
    }

    pub async fn spawn<F, Fut, T>(&self, task: F) -> JoinHandle<T>
    where
        F: FnOnce() -> Fut + Send + 'static,
        Fut: std::future::Future<Output = T> + Send + 'static,
        T: Send + 'static,
    {
        let permit = self.semaphore.clone().acquire_owned().await.unwrap();

        tokio::spawn(async move {
            let _permit = permit; // Keep permit alive for task duration
            task().await
        })
    }

    // pub fn available_workers(&self) -> usize {
    //     self.semaphore.available_permits()
    // }

    // pub fn max_workers(&self) -> usize {
    //     self.max_workers
    // }
}
