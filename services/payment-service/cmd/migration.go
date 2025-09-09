package cmd

import (
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/infrastructure/postgresql"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/pkg/logger"
	"github.com/spf13/cobra"
)

var migrationCommand = &cobra.Command{
	Use: "migration",
	Run: func(cmd *cobra.Command, args []string) {

		logger := logger.ProvideLogger()
		pgsql := postgresql.ProvidePostgreSQLInfrastructure()

		if len(args) == 0 {
			logger.Error("please insert argument up or down")
			return
		} else if args[0] == "up" {
			if err := Up(pgsql.SqlDB); err != nil {
				logger.Error(fmt.Sprintf("err : %v", err))
				return
			}
		} else if args[0] == "down" {
			if err := Down(pgsql.SqlDB); err != nil {
				logger.Error(fmt.Sprintf("err : %v", err))
				return
			}
		} else {
			logger.Error("migration argument not found")
			return
		}
	},
}
