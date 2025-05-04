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
		bootstrap.Logger.Info("Run...")
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

		go func() {
			bootstrap.Logger.Info("========== Starting GraphQL Server ==========")
			graphQLServer.Serve()
		}()

		<-quit
		if err := Shutdown(context.Background()); err != nil {
			bootstrap.Logger.Error(fmt.Sprintf("Failed to close a connection: %v", err))
			return
		}
	},
}
