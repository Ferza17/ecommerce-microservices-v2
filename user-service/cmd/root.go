package cmd

import (
	"github.com/ferza17/ecommerce-microservices-v2/user-service/config"
	"github.com/spf13/cobra"
	"log"
)

var rootCommand = &cobra.Command{
	Use: "root",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("running root command...")
	},
}

func Run() {
	cmd := &cobra.Command{}
	cmd.AddCommand(rootCommand, runCommand, migrationCommand, aclCommand)
	if err := cmd.Execute(); err != nil {
		log.Panic(err)
	}
}

func init() {
	config.SetConfig(".")
}
