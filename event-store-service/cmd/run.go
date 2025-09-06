package cmd

import (
	"context"
	"github.com/ferza17/ecommerce-microservices-v2/event-store-service/transport/rabbitmq"
	"log"
	"sync"

	grpcTransport "github.com/ferza17/ecommerce-microservices-v2/event-store-service/transport/grpc"
	"github.com/ferza17/ecommerce-microservices-v2/event-store-service/transport/http"
	"github.com/spf13/cobra"
)

var runCommand = &cobra.Command{
	Use: "run",
	Run: func(cmd *cobra.Command, args []string) {
		ctx, cancel := context.WithCancel(cmd.Context())
		defer cancel()

		grpcServer := grpcTransport.Provide()
		httpServer := http.Provide()
		rabbitmqServer := rabbitmq.Provide()

		wg := new(sync.WaitGroup)

		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := grpcServer.Serve(ctx); err != nil {
				log.Fatalf("error serving grpc server: %v", err)
				return
			}
		}()

		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := httpServer.Serve(ctx); err != nil {
				log.Fatalf("error serving http server: %v", err)
				return
			}
		}()

		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := rabbitmqServer.Serve(ctx); err != nil {
				log.Fatalf("error serving rabbitmq server: %v", err)
				return
			}
		}()

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
