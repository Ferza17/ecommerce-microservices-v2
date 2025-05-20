package cmd

import (
	"context"
	"fmt"
	bootstrap2 "github.com/ferza17/ecommerce-microservices-v2/api-gateway/bootstrap"
	"github.com/ferza17/ecommerce-microservices-v2/api-gateway/config"
	"github.com/ferza17/ecommerce-microservices-v2/api-gateway/transport/graphql"
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
	cmd.AddCommand(rootCommand, runCommand, graphqlCommand)
	if err := cmd.Execute(); err != nil {
		log.Panic(err)
	}
}

var (
	bootstrap     *bootstrap2.Bootstrap
	graphQLServer *graphql.GraphQLTransport
)

func init() {
	config.SetConfig(".")
	bootstrap = bootstrap2.NewBootstrap()
	graphQLServer = graphql.NewGraphQLTransport(
		config.Get().HttpHost,
		config.Get().HttpPort,
		bootstrap,
	)
}

func Shutdown(ctx context.Context) (err error) {
	if err = bootstrap.RabbitMQInfrastructure.Close(); err != nil {
		bootstrap.Logger.Error(fmt.Sprintf("Failed to close a connection: %v", err))
	}

	if err = bootstrap.RpcClientInfrastructure.Close(); err != nil {
		bootstrap.Logger.Error(fmt.Sprintf("Failed to close a connection: %v", err))
	}

	if err = bootstrap.TelemetryInfrastructure.Close(ctx); err != nil {
		bootstrap.Logger.Error(fmt.Sprintf("Failed to close a connection: %v", err))
	}

	bootstrap.Logger.Info("Exit...")
	return
}
