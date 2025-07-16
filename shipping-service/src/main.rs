mod cmd;
mod config;
mod infrastructure;
mod model;
mod module;

mod pkg;
mod util;

fn main() {
    cmd::root::execute()
}
