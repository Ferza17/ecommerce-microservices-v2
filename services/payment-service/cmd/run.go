package cmd

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/config"
	transportGrpc "github.com/ferza17/ecommerce-microservices-v2/payment-service/transport/grpc"
	transportHttp "github.com/ferza17/ecommerce-microservices-v2/payment-service/transport/http"
	transportKafka "github.com/ferza17/ecommerce-microservices-v2/payment-service/transport/kafka"

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

		grpcServer := transportGrpc.Provide()
		httpServer := transportHttp.Provide()
		kafkaConsumer := transportKafka.Provide()

		// Quit the signal channel
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

		// WaitGroup to ensure all services cleanups properly
		var wg sync.WaitGroup

		log.Println("running run command...")

		// Start gRPC server
		wg.Add(1)
		go func() {
			defer wg.Done()
			log.Printf("Starting gRPC server on %s:%s...", config.Get().ConfigServicePayment.RpcHost, config.Get().ConfigServicePayment.RpcPort)

			// Start the server (these blocks)
			if err := grpcServer.Serve(ctx); err != nil {
				log.Fatalf("failed to start gRPC server: %s", err)
				return
			}

		}()

		// Start HTTP Server
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := httpServer.Serve(ctx); err != nil {
				log.Fatal(err)
				return
			}
		}()

		// Start Kafka Consumer
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := kafkaConsumer.Serve(ctx); err != nil {
				log.Fatal(err)
				return
			}
		}()

		// Start HTTP Metric Collector
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := transportHttp.ServeHttpPrometheusMetricCollector(); err != nil {
				log.Fatal(err)
				return
			}
		}()

		log.Println(fmt.Sprintf("starting %s", config.Get().ConfigServicePayment.ServiceName))
		<-quit
		log.Println("Received quit signal, cleaning up...")

		// Graceful shutdown for gRPC
		//grpcServer.GracefulStop()

		// Wait for other servers to clean up
		wg.Wait()
		log.Println("All services stopped. Exiting.")

	},
}
