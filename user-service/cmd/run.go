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
		ctx, cancel := context.WithCancel(cmd.Context())
		defer cancel()

		grpcServer := grpc.ProvideGrpcServer()
		httpServer := http.ProvideHttpServer()
		rabbitMQServer := rabbitmq.ProvideRabbitMQServer()

		wg := new(sync.WaitGroup)
		wg.Add(1)
		go func() {
			defer wg.Done()
			log.Println("========== Starting RPC Server ==========")
			grpcServer.Serve()
		}()

		wg.Add(1)
		go func() {
			defer wg.Done()
			log.Println("========== Starting RabbitMQ Consumer ==========")
			rabbitMQServer.Serve(ctx)
		}()

		wg.Add(1)
		go func() {
			defer wg.Done()
			log.Println("========== Starting HTTP Server ==========")
			if err := httpServer.Serve(ctx); err != nil {
				log.Panic(err)
				return
			}
		}()

		wg.Add(1)
		go func() {
			defer wg.Done()
			log.Println("========== Starting Metric Collector HTTP Server ==========")
			if err := http.ServeHttpPrometheusMetricCollector(); err != nil {
				log.Panic(err)
				return
			}
		}()

		// Wait for all goroutines to complete
		wg.Wait()

		log.Println("All services stopped. Exiting...")
	},
}
