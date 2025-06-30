package cmd

import (
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/notification-service/infrastructure/mongodb"
	logger2 "github.com/ferza17/ecommerce-microservices-v2/notification-service/pkg/logger"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mongodb"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"path/filepath"
)

var migrationCommand = &cobra.Command{
	Use:   "migration",
	Short: "Migration database",
	Run: func(cmd *cobra.Command, args []string) {
		migrationPath := filepath.Join("dbMigration")
		migrationURL := fmt.Sprintf("file://%s", migrationPath)

		mongoDB := mongodb.ProvideMongoDBInfrastructure()
		mongoURL := mongoDB.GetConnectionString()
		logger := logger2.ProvideLogger()

		instance, err := migrate.New(migrationURL, mongoURL)
		if err != nil {
			logger.Error("CMD.MigrationCommand", zap.Error(fmt.Errorf("failed to create migration : %v", err)))
			panic(err)
		}

		if len(args) == 0 {
			logger.Error("CMD.MigrationCommand", zap.Error(fmt.Errorf("no migration version specified, insert argument up or down")))
			return
		} else if args[0] == "up" {
			if err := instance.Up(); err != nil {
				logger.Error("CMD.MigrationCommand", zap.Error(fmt.Errorf("failed to migrate up : %v", err)))
				return
			}
		} else if args[0] == "down" {
			if err := instance.Down(); err != nil {
				logger.Error("CMD.MigrationCommand", zap.Error(fmt.Errorf("failed to create down : %v", err)))
				return
			}
		} else {
			logger.Error("CMD.MigrationCommand", zap.Error(fmt.Errorf("invalid migration version, insert argument up or down")))
			return
		}

	},
}
