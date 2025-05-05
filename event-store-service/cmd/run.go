package cmd

import (
	"context"
	"fmt"
	rabbitMQServer "github.com/ferza17/ecommerce-microservices-v2/event-store-service/server/rabbitmq"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"syscall"
)

var runCommand = &cobra.Command{
	Use: "run",
	Run: func(cmd *cobra.Command, args []string) {
		dependency.Logger.Info("Run...")
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

		go func() {
			dependency.Logger.Info("========== Starting GraphQL Server ==========")
			rabbitMQServer.
				NewServer(
					rabbitMQServer.NewLogger(dependency.Logger),
					rabbitMQServer.NewMongoDBInfrastructure(dependency.MongoDBInfrastructure),
				).
				Serve()
		}()

		<-quit
		if err := Shutdown(context.Background()); err != nil {
			dependency.Logger.Error(fmt.Sprintf("Failed to close a connection: %v", err))
			return
		}
	},
}
