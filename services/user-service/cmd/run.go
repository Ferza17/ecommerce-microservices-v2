package cmd

import (
	"context"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/config"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/transport/grpc"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/transport/http"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/transport/kafka"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/transport/rabbitmq"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/spf13/cobra"
)

var runCommand = &cobra.Command{
	Use: "run",
	Run: func(cmd *cobra.Command, args []string) {
		// Root context with cancel
		ctx, cancel := context.WithCancel(context.Background())

		// Initialize servers
		grpcServer := grpc.ProvideGrpcServer()
		httpServer := http.ProvideHttpServer()
		rabbitMQServer := rabbitmq.ProvideServer()
		kafkaServer := kafka.ProvideServer()

		wg := new(sync.WaitGroup)
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

		log.Println("Starting services in env:", config.Get().Env)

		// Start servers
		start := func(name string, fn func(context.Context) error, closeFn func()) {
			wg.Add(1)
			go func() {
				defer wg.Done()
				if err := fn(ctx); err != nil {
					log.Printf("error serving %s: %v", name, err)
				}
				// closeFn should only be called *after* Serve exits
				if closeFn != nil {
					closeFn()
				}
			}()
		}

		start("gRPC", grpcServer.Serve, grpcServer.Close)
		start("RabbitMQ", rabbitMQServer.Serve, rabbitMQServer.Close)
		start("Kafka", kafkaServer.Serve, kafkaServer.Close)
		start("HTTP", httpServer.Serve, httpServer.Close)

		// Metrics server (fire and forget, doesn’t need to close)
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := http.ServeHttpPrometheusMetricCollector(); err != nil {
				log.Printf("error serving metrics: %v", err)
			}
		}()

		// Wait for the quit signal
		<-quit
		log.Println("Received quit signal, shutting down...")
		cancel() // tell all servers to stop

		// Wait for all goroutines to finish
		wg.Wait()
		log.Println("All services stopped. Exiting...")
	},
}
