package cmd

import (
	"log"

	"github.com/ferza17/ecommerce-microservices-v2/event-store-service/config"
	"github.com/spf13/cobra"
)

var rootCommand = &cobra.Command{
	Use: "root",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("running root command...")
	},
}

func Run() {
	cmd := &cobra.Command{}
	cmd.AddCommand(rootCommand, runCommand)
	if err := cmd.Execute(); err != nil {
		log.Panic(err)
	}
}

func init() {
	config.SetConfig(".")
}
