package cmd

import (
	"database/sql"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/config"
	"github.com/pressly/goose/v3"
	"github.com/spf13/cobra"
	"log"
)

var rootCommand = &cobra.Command{
	Use: "root",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("running root command...")
	},
}

func init() {
	config.SetConfig(".")
}

func Run() {
	cmd := &cobra.Command{}
	cmd.AddCommand(rootCommand, migrationCommand, runCommand)
	if err := cmd.Execute(); err != nil {
		log.Panic(err)
	}
}

func Up(db *sql.DB) error {
	if err := goose.SetDialect("postgres"); err != nil {
		return err
	}

	if err := goose.Up(db, "dbMigration"); err != nil {
		return err
	}

	return nil
}

func Down(db *sql.DB) error {
	if err := goose.SetDialect("postgres"); err != nil {
		return err
	}

	if err := goose.Down(db, "dbMigration"); err != nil {
		return err
	}

	return nil
}
