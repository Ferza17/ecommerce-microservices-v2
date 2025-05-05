package cmd

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/event-store-service/bootstrap"
	"github.com/ferza17/ecommerce-microservices-v2/event-store-service/config"
	"github.com/spf13/cobra"
	"log"
)

var rootCommand = &cobra.Command{
	Use: "root",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("running root command...")
	},
}

func Run() {
	cmd := &cobra.Command{}
	cmd.AddCommand(rootCommand, runCommand)
	if err := cmd.Execute(); err != nil {
		log.Panic(err)
	}
}

var (
	dependency *bootstrap.Bootstrap
)

func init() {
	config.SetConfig(".")
	dependency = bootstrap.NewBootstrap()
}

func Shutdown(ctx context.Context) (err error) {
	if err = dependency.RabbitMQInfrastructure.Close(); err != nil {
		dependency.Logger.Error(fmt.Sprintf("Failed to close a connection: %v", err))
		return err
	}

	if err = dependency.MongoDBInfrastructure.Close(ctx); err != nil {
		dependency.Logger.Error(fmt.Sprintf("Failed to close a connection: %v", err))
		return err
	}

	dependency.Logger.Info("Exit...")
	return
}
