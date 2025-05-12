package cmd

import (
	"database/sql"
	"github.com/pressly/goose/v3"
	"github.com/spf13/cobra"
	"log"
)

var migrationCommand = &cobra.Command{
	Use:   "migration",
	Short: "Migration database",
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) == 0 {
			log.Fatalf("please insert argument up or down")
			return
		} else if args[0] == "up" {
			if err := Up(dependency.PostgreSQLInfrastructure.SqlDB()); err != nil {
				log.Fatalf(err.Error())
				return
			}
		} else if args[0] == "down" {
			if err := Down(dependency.PostgreSQLInfrastructure.SqlDB()); err != nil {
				log.Fatalf(err.Error())
				return
			}
		} else {
			log.Fatalf("migration argument not found")
			return
		}
	},
}

func Up(db *sql.DB) error {
	if err := goose.SetDialect("postgres"); err != nil {
		return err
	}

	if err := goose.Up(db, "dbMigration/postgres"); err != nil {
		return err
	}

	return nil
}

func Down(db *sql.DB) error {
	if err := goose.SetDialect("postgres"); err != nil {
		return err
	}

	if err := goose.Down(db, "dbMigration/postgres"); err != nil {
		return err
	}

	return nil
}
