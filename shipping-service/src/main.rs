mod cmd;
mod config;
mod infrastructure;
mod interceptor;
mod model;
mod module;
mod package;
mod transport;
mod util;
use tracing_subscriber;

#[tokio::main]
async fn main() {
    tracing_subscriber::fmt::init();

    cmd::root::execute().await;
}
