use crate::package::worker_pool::worker_pool::WorkerPool;

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
            http_pool: WorkerPool::new(http_workers),
            grpc_pool: WorkerPool::new(grpc_workers),
            messaging_pool: WorkerPool::new(messaging_workers),
            metrics_pool: WorkerPool::new(metrics_workers),
        }
    }
}
