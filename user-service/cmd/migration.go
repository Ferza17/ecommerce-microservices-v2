package cmd

import (
	"database/sql"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/postgres"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/pkg/logger"
	"github.com/pressly/goose/v3"
	"github.com/spf13/cobra"
)

var migrationCommand = &cobra.Command{
	Use:   "migration",
	Short: "Migration database",
	Run: func(cmd *cobra.Command, args []string) {
		logger := logger.ProvideLogger()
		postgres := postgres.ProvidePostgresInfrastructure()

		if len(args) == 0 {
			logger.Error("please insert argument up or down")
			return
		} else if args[0] == "up" {
			if err := Up(postgres.SqlDB); err != nil {
				logger.Error(err.Error())
				return
			}
		} else if args[0] == "down" {
			if err := Down(postgres.SqlDB); err != nil {
				logger.Error(err.Error())
				return
			}
		} else {
			logger.Error("migration argument not found")
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
