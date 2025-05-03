package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

// Initialize GRPC Server
var graphqlCommand = &cobra.Command{
	Use: "graphql",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("========== Starting GraphQL Server ==========")

		//quit := make(chan os.Signal, 1)
		//signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		//
		//go func() {
		//	grpcServer.Serve()
		//}()
		//
		//<-quit
		//
		//if err := Shutdown(context.Background()); err != nil {
		//	log.Fatalln(err)
		//	return
		//}
	},
}
