package cmd

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/bootstrap"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/config"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/server/grpc"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/server/rabbitmq"

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

	if err = dependency.PostgresqlInfrastructure.Close(); err != nil {
		dependency.Logger.Error(fmt.Sprintf("failed to close postgresql connection : %v", err))
		return err
	}

	if err = dependency.RabbitMQInfrastructure.Close(); err != nil {
		dependency.Logger.Error(fmt.Sprintf("failed to close rabbitmq connection : %v", err))
		return err
	}

	log.Println("Shutdown...")
	return
}
