use clap::Args;

#[derive(Args, Debug)]
pub struct MigrationArgs {
    #[arg(short, long, help = "Migration direction: 'up' or 'down'")]
    pub direction: String,
}

pub fn handle_migration_command(args: MigrationArgs) {
    if args.direction == "up" {
        println!("Running UP migrations...");
    } else if args.direction == "down" {
        println!("Running DOWN migrations...");
    } else {
        eprintln!("Invalid migration direction. Use 'up' or 'down'.");
        std::process::exit(1);
    }
}
