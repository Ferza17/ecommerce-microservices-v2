package cmd

import (
	"context"
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var runCommand = &cobra.Command{
	Use: "run",
	Run: func(cmd *cobra.Command, args []string) {
		chClose := make(chan os.Signal, 2)
		signal.Notify(chClose, syscall.SIGINT, syscall.SIGTERM)
		go func() {
			log.Println("========== Starting RPC Server ==========")
			//grpc.NewServer(
			//	os.Getenv("RPC_HOST"),
			//	os.Getenv("RPC_PORT"),
			//	//grpc.NewTracer(tracer),
			//	grpc.NewLogger(logger),
			//	//grpc.NewRedisClient(redisClient),
			//	//grpc.NewRabbitMQConnection(rabbitMQConnection),
			//	//grpc.NewElasticsearchClient(esClient),
			//	//grpc.NewPostgresClient(postgresSQlClient),
			//	//grpc.NewCassandraSession(cassandraSession),
			//).Serve()
		}()
		go func() {
			//TODO:
			// 1. Add Consumer Server
		}()

		<-chClose
		if err := Shutdown(context.Background()); err != nil {
			log.Fatalln(err)
			return
		}
		log.Println("Exit...")
	},
}
