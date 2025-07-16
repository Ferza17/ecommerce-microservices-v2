use crate::config::config::AppConfig;
use clap::Args;

#[derive(Args, Debug)]
pub struct RunArgs {
    #[arg(short, long, help = "run direction: 'local' or 'production'")]
    pub direction: String,
}
pub fn handle_run_command(args: RunArgs) {
    //TODO: RUN GRPC, HTTP, RABBITMQ

    let cfg = AppConfig::set_config(&*args.direction);

    println!("{:?}", cfg);
    println!("run");
}
