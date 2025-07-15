package cmd

import (
	"context"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/transport/grpc"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/transport/http"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/transport/rabbitmq"
	"github.com/spf13/cobra"
	"log"
	"sync"
)

var runCommand = &cobra.Command{
	Use: "run",
	Run: func(cmd *cobra.Command, args []string) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		// Initialize servers
		grpcServer := grpc.ProvideGrpcServer()
		httpServer := http.ProvideHttpServer()
		rabbitMQServer := rabbitmq.ProvideRabbitMQServer()

		wg := new(sync.WaitGroup)

		// Start GRPC Server
		wg.Add(1)
		go func() {
			defer wg.Done()
			err := grpcServer.Serve(ctx)
			if err != nil {
				return
			}
		}()

		// Start RabbitMQ Consumer
		wg.Add(1)
		go func() {
			defer wg.Done()
			rabbitMQServer.Serve(ctx)
		}()

		// Start HTTP Server
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := httpServer.Serve(ctx); err != nil {
				log.Panic(err)
				return
			}
		}()

		// Start Metric Collector (simplified)
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := http.ServeHttpPrometheusMetricCollector(); err != nil {
				return
			}
		}()

		// Wait for all goroutines to complete
		wg.Wait()

		log.Println("All services stopped. Exiting...")
	},
}
