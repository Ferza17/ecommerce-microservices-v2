package cmd

import (
	"context"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/config"
	"log"
	"sync"

	"github.com/ferza17/ecommerce-microservices-v2/user-service/transport/grpc"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/transport/http"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/transport/rabbitmq"
	"github.com/spf13/cobra"
)

var runCommand = &cobra.Command{
	Use: "run",
	Run: func(cmd *cobra.Command, args []string) {
		ctx, cancel := context.WithCancel(cmd.Context())
		defer cancel()

		// Initialize servers
		grpcServer := grpc.ProvideGrpcServer()
		httpServer := http.ProvideHttpServer()
		rabbitMQServer := rabbitmq.ProvideRabbitMQServer()

		wg := new(sync.WaitGroup)

		log.Println("Starting services in env : ", config.Get().Env)

		// Start GRPC Server
		wg.Add(1)
		go func() {
			defer wg.Done()
			err := grpcServer.Serve(ctx)
			if err != nil {
				log.Fatalf("error serving grpc server: %v", err)
				return
			}
		}()

		// Start RabbitMQ Consumer
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := rabbitMQServer.Serve(ctx); err != nil {
				log.Fatalf("error serving rabbitmq server: %v", err)
				return
			}
		}()

		// Start HTTP Server
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := httpServer.Serve(ctx); err != nil {
				log.Fatalf("error serving http server: %v", err)
				return
			}
		}()

		// Start Metric Collector (simplified)
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := http.ServeHttpPrometheusMetricCollector(); err != nil {
				log.Fatalf("error serving http metric collector: %v", err)
				return
			}
		}()

		// Wait for all goroutines to complete
		wg.Wait()

		log.Println("All services stopped. Exiting...")
	},
}
