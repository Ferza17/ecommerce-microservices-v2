package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

var rootCommand = &cobra.Command{
	Use: "root",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("running root command...")
	},
}

func init() {

}

func Run() {
	cmd := &cobra.Command{}
	cmd.AddCommand(rootCommand, migrationCommand, rootCommand)
	if err := cmd.Execute(); err != nil {
		log.Panic(err)
	}
}
