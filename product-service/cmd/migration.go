package cmd

import (
	"database/sql"
	pgsql "github.com/ferza17/ecommerce-microservices-v2/product-service/connector/postgresql"
	"github.com/pressly/goose/v3"
	"github.com/spf13/cobra"
	"log"
)

var migrationCommand = &cobra.Command{
	Use:   "migration",
	Short: "Migration database",
	Run: func(cmd *cobra.Command, args []string) {

		pqConn := pgsql.NewPostgres()
		if len(args) == 0 {
			log.Fatalf("please insert argument up or down")
			return
		} else if args[0] == "up" {
			if err := Up(pqConn.SqlDB); err != nil {
				log.Fatalf(err.Error())
				return
			}
		} else if args[0] == "down" {
			if err := Down(pqConn.SqlDB); err != nil {
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

func init() {
	rootCommand.AddCommand(migrationCommand)
}
