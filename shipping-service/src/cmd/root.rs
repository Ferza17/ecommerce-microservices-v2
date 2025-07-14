use crate::cmd::run::run;
use clap::{Parser, Subcommand};

#[derive(Parser)]
#[command(name = "shipping-service", version = "1.0", author = "You")]
#[command(about = "A CLI app similar to Cobra")]
struct Cli {
    #[command(subcommand)]
    pub command: Commands,
}

#[derive(Subcommand)]
enum Commands {
    Run,
}

pub fn run_command() {
    let cli = Cli::parse();

    match &cli.command {
        Commands::Run {} => run(),
    }
}
