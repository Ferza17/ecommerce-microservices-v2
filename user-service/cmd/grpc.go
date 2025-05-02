package cmd

import (
	"context"
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/signal"
	"syscall"
)

// Initialize GRPC Server
var rpcCommand = &cobra.Command{
	Use: "grpc",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("========== Starting RPC Server ==========")

		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

		go func() {
			grpcServer.Serve()
		}()

		<-quit

		if err := Shutdown(context.Background()); err != nil {
			log.Fatalln(err)
			return
		}
	},
}
