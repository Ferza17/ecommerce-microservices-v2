package cmd

import (
	"context"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/transport/grpc"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/transport/http"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/transport/rabbitmq"
	"github.com/spf13/cobra"
	"log"
	"sync"
)

var runCommand = &cobra.Command{
	Use: "run",
	Run: func(cmd *cobra.Command, args []string) {
		ctx, cancel := context.WithCancel(cmd.Context())
		defer cancel()

		grpcServer := grpc.ProvideGrpcTransport()
		httpServer := http.ProvideHttpTransport()
		rabbitMQServer := rabbitmq.ProvideRabbitMQTransport()

		wg := new(sync.WaitGroup)
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := grpcServer.Serve(ctx); err != nil {
				log.Fatalln(err)
			}
		}()

		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := rabbitMQServer.Serve(ctx); err != nil {
				log.Fatalln(err)
			}
		}()

		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := httpServer.Serve(ctx); err != nil {
				log.Println("HTTP Server Error: ", err)
				return
			}
		}()

		wg.Add(1)
		go func() {
			defer wg.Done()
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
