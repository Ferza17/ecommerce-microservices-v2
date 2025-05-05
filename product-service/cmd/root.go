package cmd

import (
	"context"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/bootstrap"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/config"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/server/grpc"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/server/rabbitmq"
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
	cmd.AddCommand(rootCommand, runCommand, rpcCommand, migrationCommand)
	if err := cmd.Execute(); err != nil {
		log.Panic(err)
	}
}

var (
	dependency     *bootstrap.Bootstrap
	grpcServer     *grpc.Server
	rabbitMQServer *rabbitmq.Server
)

func init() {
	config.SetConfig(".")
	dependency = bootstrap.NewBootstrap()
	grpcServer = grpc.NewServer(dependency)
	rabbitMQServer = rabbitmq.NewServer(dependency)
}

func Shutdown(ctx context.Context) (err error) {
	grpcServer.GracefulStop()

	if err = dependency.PostgreSQLInfrastructure.Close(); err != nil {
		return err
	}

	log.Println("Shutdown...")
	return
}
