package cmd

import (
	"log"

	"github.com/piovani/kafka_aux/config"
	"github.com/spf13/cobra"
)

func Execute() {
	root := &cobra.Command{
		Short:   "Kafka Aux",
		Version: "1.0.0",
	}

	if err := config.StarConfig(); err != nil {
		log.Fatal(err)
	}

	root.AddCommand(
		Rest,
	)

	root.Execute()
}
