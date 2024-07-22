package cmd

import (
	"log"

	"github.com/piovani/kafka_aux/ui/api"
	"github.com/spf13/cobra"
)

var Rest = &cobra.Command{
	Use:   "api",
	Short: "Start API Rest running",
	Run: func(cmd *cobra.Command, args []string) {
		api := api.NewApi()
		if err := api.StartRest(); err != nil {
			log.Fatal(err)
		}
	},
}
