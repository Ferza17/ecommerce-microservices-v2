package cmd

import (
	"context"
	"fmt"
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
			dependency.Logger.Info("========== Starting GraphQL RabbitMQTransport ==========")
			rabbitMQServer.Serve()
		}()

		<-quit
		if err := Shutdown(context.Background()); err != nil {
			dependency.Logger.Error(fmt.Sprintf("Failed to close a connection: %v", err))
			return
		}
	},
}
