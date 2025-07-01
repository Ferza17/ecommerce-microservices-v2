package cmd

import (
	"context"
	"database/sql"
	"errors"
	postgres2 "github.com/ferza17/ecommerce-microservices-v2/product-service/infrastructure/postgres"
	productEsRepo "github.com/ferza17/ecommerce-microservices-v2/product-service/module/product/repository/elasticsearch"
	productPgRepo "github.com/ferza17/ecommerce-microservices-v2/product-service/module/product/repository/postgres"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/pkg/logger"
	"github.com/pressly/goose/v3"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var migrationCommand = &cobra.Command{
	Use:   "migration",
	Short: "Migration database",
	Run: func(cmd *cobra.Command, args []string) {
		logger := logger.ProvideLogger()
		postgres := postgres2.ProvidePostgresSQL()

		if len(args) == 0 {
			logger.Error("cmd.migrationCommand", zap.Error(errors.New("please insert argument up, down, elasticsearch")))
			return
		} else if args[0] == "up" {
			if err := Up(postgres.SqlDB, logger); err != nil {
				logger.Error("cmd.migrationCommand", zap.Error(err), zap.String("command", "migration up"))
				return
			}
		} else if args[0] == "down" {
			if err := Down(postgres.SqlDB, logger); err != nil {
				logger.Error("cmd.migrationCommand", zap.Error(err), zap.String("command", "migration down"))
				return
			}
		} else if args[0] == "elasticsearch" {
			if err := elasticsearch(context.Background(), logger, postgres.GormDB); err != nil {
				logger.Error("cmd.migrationCommand", zap.Error(err), zap.String("command", "migration elasticsearch"))
				return
			}
		} else {
			logger.Error("cmd.migrationCommand", zap.Error(errors.New("migration argument not found")), zap.String("command", "migration up"))
			return
		}
	},
}

func Up(db *sql.DB, logger logger.IZapLogger) error {
	if err := goose.SetDialect("postgres"); err != nil {
		logger.Error("cmd.migrationCommand", zap.Error(err), zap.String("command", "migration up"))
		return err
	}

	if err := goose.Up(db, "dbMigration/postgres"); err != nil {
		logger.Error("cmd.migrationCommand", zap.Error(err), zap.String("command", "migration up"))
		return err
	}

	return nil
}

func Down(db *sql.DB, logger logger.IZapLogger) error {
	if err := goose.SetDialect("postgres"); err != nil {
		logger.Error("cmd.migrationCommand", zap.Error(err), zap.String("command", "migration down"))
		return err
	}

	if err := goose.Down(db, "dbMigration/postgres"); err != nil {
		logger.Error("cmd.migrationCommand", zap.Error(err), zap.String("command", "migration down"))
		return err
	}

	return nil
}

func elasticsearch(ctx context.Context, logger logger.IZapLogger, tx *gorm.DB) error {
	productPgRepo := productPgRepo.ProvideProductPostgresSQLRepository()
	productEsRepo := productEsRepo.ProvideProductElasticsearchRepository()
	tx = tx.Begin()

	products, err := productPgRepo.FindAllProductForElasticIndex(ctx, tx)
	if err != nil {
		tx.Rollback()
		logger.Error("cmd.migrationCommand", zap.Error(err), zap.String("command", "migration elasticsearch"))
		return err
	}

	if err = productEsRepo.BulkCreateProduct(ctx, products); err != nil {
		tx.Rollback()
		logger.Error("cmd.migrationCommand", zap.Error(err), zap.String("command", "migration elasticsearch"))
		return err
	}

	return nil
}
