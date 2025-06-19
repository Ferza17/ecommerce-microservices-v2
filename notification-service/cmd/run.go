package cmd

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/notification-service/config"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/health/grpc_health_v1"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

var runCommand = &cobra.Command{
	Use: "run",
	Run: func(cmd *cobra.Command, args []string) {
		// Quit the signal channel
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

		p := make(chan bool)

		// WaitGroup to ensure all services cleanups properly
		wg := new(sync.WaitGroup)

		log.Println("running run command...")

		// Start gRPC server
		wg.Add(1)
		go func() {
			defer wg.Done()
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
			rabbitMQTransport.Serve()
		}()

		// Register Consul service after gRPC is ready
		go func(c *config.Config) {
			<-p
			log.Println("Waiting for gRPC server to be ready...")

			var lastErr error
			for i := 0; i < 10; i++ { // Increase retry attempts
				conn, err := grpc.NewClient(
					fmt.Sprintf("%s:%s", c.RpcHost, c.RpcPort),
					grpc.WithTransportCredentials(insecure.NewCredentials()),
				)
				if err != nil {
					lastErr = err
					log.Printf("Failed to dial gRPC server: %v", err)
					time.Sleep(2 * time.Second)
					continue
				}

				healthClient := grpc_health_v1.NewHealthClient(conn)
				ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
				resp, err := healthClient.Check(ctx, &grpc_health_v1.HealthCheckRequest{})
				cancel()
				if err := conn.Close(); err != nil {
					log.Printf("Failed to close a connection: %v", err)
					continue
				}

				if err == nil && resp.Status == grpc_health_v1.HealthCheckResponse_SERVING {
					log.Println("gRPC server is ready and healthy.")
					break
				}

				lastErr = err
				log.Printf("Health check failed (attempt %d/10): %v", i+1, err)
				time.Sleep(3 * time.Second)
			}

			if lastErr != nil {
				log.Printf("Warning: gRPC server health check failed after 10 attempts: %v", lastErr)
			}

			log.Println("Registering service to Consul...")
			if err := c.RegisterConsulService(); err != nil {
				log.Printf("Error in Consul service registration: %v", err)
			}

		}(config.Get())

		log.Println(fmt.Sprintf("starting %s", config.Get().ServiceName))
		<-quit
		log.Println("Received quit signal, cleaning up...")

		// Graceful shutdown for gRPC
		grpcServer.GracefulStop()

		// Wait for other servers to clean up
		wg.Wait()
		log.Println("All services stopped. Exiting.")

	},
}
