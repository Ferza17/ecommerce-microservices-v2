mod cmd;
mod config;
mod infrastructure;
mod interceptor;
mod model;
mod module;
mod package;
mod transport;
mod util;

#[tokio::main]
async fn main() {
    cmd::root::execute().await;
}
