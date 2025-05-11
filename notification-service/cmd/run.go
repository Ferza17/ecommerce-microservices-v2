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
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

		go func() {
			rabbitMQTransport.Serve()
		}()

		<-quit
		if err := Shutdown(context.Background()); err != nil {
			panic(err)
		}
		log.Println("Exiting...")
	},
}
