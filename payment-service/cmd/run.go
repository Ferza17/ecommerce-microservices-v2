package cmd

import (
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/transport/grpc"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/transport/rabbitmq"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"syscall"
)

var runCommand = &cobra.Command{
	Use: "run",
	Run: func(cmd *cobra.Command, args []string) {
		grpcServer := grpc.ProvideGrpcServer()
		rebbitMQServer := rabbitmq.ProvideGrpcServer()
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

		go func() {
			grpcServer.Serve()
		}()

		go func() {
			rebbitMQServer.Serve()
		}()

		<-quit
		grpcServer.GracefulStop()
	},
}
