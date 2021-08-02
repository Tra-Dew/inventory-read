package main

import (
	"github.com/d-leme/tradew-inventory-read/cmd"
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

	itemsUpdatedConsumer := &cobra.Command{
		Use:   "items-updated-consumer",
		Short: "Starts items-updated-consumer",
		Run:   cmd.ItemsUpdated,
	}

	root.PersistentFlags().String("settings", "./settings.yml", "path to settings.yaml config file")
	root.AddCommand(api, itemsUpdatedConsumer)

	root.Execute()
}
