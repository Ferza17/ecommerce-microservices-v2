// Initialize GRPC Server

package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

var rpcCommand = &cobra.Command{
	Use: "grpc",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("========== Starting RPC Server ==========")
		//grpc.NewServer(
		//	os.Getenv("RPC_HOST"),
		//	os.Getenv("RPC_PORT"),
		//	grpc.NewLogger(logger),
		//).Serve()
	},
}
