package main

import (
	"github.com/Tra-Dew/inventory-read/cmd"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			logrus.WithField("error", err).Error("error main")
		}
	}()

	root := &cobra.Command{}

	api := &cobra.Command{
		Use:   "api",
		Short: "Starts api handlers",
		Run:   cmd.Server,
	}

	itemsCreatedConsumer := &cobra.Command{
		Use:   "items-created-consumer",
		Short: "Starts items-created-consumer",
		Run:   cmd.ItemsCreatedRequested,
	}

	root.PersistentFlags().String("settings", "./settings.yml", "path to settings.yaml config file")
	root.AddCommand(api, itemsCreatedConsumer)

	root.Execute()
}
