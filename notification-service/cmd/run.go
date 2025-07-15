package cmd

import (
	"context"
	"github.com/ferza17/ecommerce-microservices-v2/notification-service/transport/grpc"
	"github.com/ferza17/ecommerce-microservices-v2/notification-service/transport/http"
	"github.com/ferza17/ecommerce-microservices-v2/notification-service/transport/rabbitmq"
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
		rabbitMQServer := rabbitmq.ProvideRabbitMQServer()

		wg := new(sync.WaitGroup)
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := grpcServer.Serve(ctx); err != nil {
				log.Fatalf("failed to serve : %s", err)
				return
			}
		}()

		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := rabbitMQServer.Serve(ctx); err != nil {
				log.Fatalf("failed to serve : %s", err)
			}
		}()

		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := http.ServeHttpPrometheusMetricCollector(); err != nil {
				log.Fatal(err)
				return
			}
		}()

		// Wait for all goroutines to complete
		wg.Wait()

		log.Println("All services stopped. Exiting...")
	},
}
