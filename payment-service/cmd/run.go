package cmd

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/config"
	transportGrpc "github.com/ferza17/ecommerce-microservices-v2/payment-service/transport/grpc"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/transport/http"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/transport/rabbitmq"
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
		ctx, cancel := context.WithCancel(cmd.Context())
		defer cancel()

		grpcServer := transportGrpc.ProvideGrpcServer()
		httpServer := http.ProvideHttpServer()
		rabbitMQServer := rabbitmq.ProvideGrpcServer()

		// Quit the signal channel
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

		p := make(chan bool)

		// WaitGroup to ensure all services cleanups properly
		var wg sync.WaitGroup

		log.Println("running run command...")

		// Start gRPC server
		wg.Add(1)
		go func() {
			defer wg.Done()
			log.Printf("Starting gRPC server on %s:%s...", config.Get().PaymentServiceRpcHost, config.Get().PaymentServiceRpcPort)

			// Signal that we're about to start the server
			close(p)

			// Start the server (these blocks)
			grpcServer.Serve()

		}()

		// Start RabbitMQ server
		wg.Add(1)
		go func() {
			defer wg.Done()
			log.Println("Starting RabbitMQ server...")
			rabbitMQServer.Serve()
		}()

		// Start HTTP Server
		wg.Add(1)
		go func() {
			defer wg.Done()
			log.Printf("Starting HTTP server on %s:%s...", config.Get().PaymentServiceHttpHost, config.Get().PaymentServiceHttpPort)
			if err := httpServer.Serve(ctx); err != nil {
				log.Fatal(err)
				return
			}
		}()

		// Start HTTP Metric Collector
		wg.Add(1)
		go func() {
			defer wg.Done()
			log.Printf("Starting HTTP Metric Collector on %s:%s...", config.Get().PaymentServiceHttpHost, config.Get().PaymentServiceMetricHttpPort)
			if err := http.ServeHttpPrometheusMetricCollector(); err != nil {
				log.Fatal(err)
				return
			}
		}()

		log.Println(fmt.Sprintf("starting %s", config.Get().PaymentServiceServiceName))
		<-quit
		log.Println("Received quit signal, cleaning up...")

		// Graceful shutdown for gRPC
		grpcServer.GracefulStop()

		// Wait for other servers to clean up
		wg.Wait()
		log.Println("All services stopped. Exiting.")

	},
}
