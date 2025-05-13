package cmd

import (
	"context"
	"database/sql"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/bootstrap"
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
		} else if args[0] == "elasticsearch" {
			if err := elasticsearch(context.Background(), dependency); err != nil {
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

func elasticsearch(ctx context.Context, dependency *bootstrap.Bootstrap) error {
	tx := dependency.ProductPostgresSQLRepository.OpenTransactionWithContext(ctx)

	products, err := dependency.ProductPostgresSQLRepository.FindAllProductForElasticIndex(ctx, tx)
	if err != nil {
		tx.Rollback()
		return err
	}

	if err = dependency.ProductElasticsearchRepository.BulkCreateProduct(ctx, products); err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
