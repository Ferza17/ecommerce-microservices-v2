// Initialize GRPC Server

package cmd

import (
	"github.com/ferza17/ecommerce-microservices-v2/product-service/config"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/grpc"
	"github.com/spf13/cobra"
	"log"
)

var rpcCommand = &cobra.Command{
	Use: "grpc",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("========== Starting RPC Server ==========")

		grpc.NewServer(
			config.Get().RpcHost,
			config.Get().RpcPort,
			grpc.NewLogger(logger),
			grpc.NewPostgresConnector(pgsqlConn),
		).Serve()
	},
}
