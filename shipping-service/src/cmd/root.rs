use crate::cmd::{migration, run};
use clap::{Parser, Subcommand};

// Main CLI struct, derive Parser to parse arguments from the command line
#[derive(Parser, Debug)]
#[command(author, version, about = "SHIPPING Service CLI", long_about = None)]
pub struct Cli {
    #[command(subcommand)]
    pub command: Commands,
}

// Subcommands for the CLI
#[derive(Subcommand, Debug)]
pub enum Commands {
    Run(run::RunArgs),

    /// Perform database migrations (up or down)
    Migration(migration::MigrationArgs),
}

pub async fn execute() {
    let cli = Cli::parse();
    println!("Parsed CLI: {:?}", cli);

    match cli.command {
        Commands::Run(args) => {
            run::handle_run_command(args).await;
        }
        Commands::Migration(args) => {
            migration::handle_migration_command(args);
        }
    }
}
