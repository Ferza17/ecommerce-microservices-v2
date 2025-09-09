mod cmd;
mod config;
mod infrastructure;

mod interceptor;
mod module;
mod package;
mod transport;
mod util;

mod model;

#[tokio::main]
async fn main() {
    cmd::root::execute().await;
}
