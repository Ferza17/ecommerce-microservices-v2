package cmd

import (
	"github.com/ferza17/ecommerce-microservices-v2/user-service/transport/grpc"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/transport/rabbitmq"
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var runCommand = &cobra.Command{
	Use: "run",
	Run: func(cmd *cobra.Command, args []string) {
		//ctx, cancel := context.WithCancel(cmd.Context())
		//defer cancel()

		grpcServer := grpc.ProvideGrpcServer()
		rabbitMQServer := rabbitmq.ProvideRabbitMQServer()

		// Quit signal channel
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

		// Ensure Service Cleanup Properly
		wg := new(sync.WaitGroup)
		// Error channel to collect errors from goroutines
		errChan := make(chan error, 2)

		wg.Add(1)
		go func() {
			log.Println("========== Starting RPC Server ==========")
			grpcServer.Serve()
		}()

		wg.Add(1)
		go func() {
			log.Println("========== Starting RabbitMQ Consumer ==========")
			rabbitMQServer.Serve()
		}()

		// Wait for shutdown signal or server errors
		select {
		case <-quit:
			log.Println("Received quit signal, cleaning up...")
		case err := <-errChan:
			log.Printf("Server error occurred: %v", err)
		}

		// Graceful shutdown
		log.Println("Shutting down servers...")

		// Shutdown RPC server
		grpcServer.GracefulStop()

		// Wait for all goroutines to complete
		wg.Wait()
		log.Println("All services stopped. Exiting...")
	},
}
