package cmd

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mongodb"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/spf13/cobra"
	"path/filepath"
)

var migrationCommand = &cobra.Command{
	Use:   "migration",
	Short: "Migration database",
	Run: func(cmd *cobra.Command, args []string) {
		migrationPath := filepath.Join("dbMigration")
		migrationURL := fmt.Sprintf("file://%s", migrationPath)

		mongoURL := dependency.MongoDBInfrastructure.GetConnectionString()

		instance, err := migrate.New(migrationURL, mongoURL)
		if err != nil {
			dependency.Logger.Error(fmt.Sprintf("failed to create migration : %v", err))
			panic(err)
		}

		if len(args) == 0 {
			dependency.Logger.Info("No migration version specified, insert argument up or down")
			return
		} else if args[0] == "up" {
			if err := instance.Up(); err != nil {
				dependency.Logger.Error(fmt.Sprintf("failed to migrate up : %v", err))
				return
			}
		} else if args[0] == "down" {
			if err := instance.Down(); err != nil {
				dependency.Logger.Error(fmt.Sprintf("failed to migrate down : %v", err))
				return
			}
		} else {
			dependency.Logger.Info("Invalid migration version, insert argument up or down")
			return
		}

	},
}
