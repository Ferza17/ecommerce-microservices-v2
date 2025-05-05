package cmd

import (
	"context"
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var runCommand = &cobra.Command{
	Use: "run",
	Run: func(cmd *cobra.Command, args []string) {
		quit := make(chan os.Signal, 2)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

		go func() {
			log.Println("========== Starting RPC Server ==========")
			grpcServer.Serve()
		}()
		go func() {
			log.Println("========== Starting RabbitMQ Consumer ==========")
			rabbitMQServer.Serve()
		}()

		<-quit
		if err := Shutdown(context.Background()); err != nil {
			log.Fatalln(err)
			return
		}
		log.Println("Exit...")
	},
}
