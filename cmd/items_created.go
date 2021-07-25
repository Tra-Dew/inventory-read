package cmd

import (
	"context"
	"reflect"

	"github.com/d-leme/tradew-inventory-read/pkg/core"
	"github.com/d-leme/tradew-inventory-read/pkg/inventory"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// ItemsCreated ...
func ItemsCreated(command *cobra.Command, args []string) {
	settings := new(core.Settings)

	err := core.FromYAML(command.Flag("settings").Value.String(), settings)
	if err != nil {
		logrus.
			WithError(err).
			Fatal("unable to parse settings, shutting down...")
	}

	container := NewContainer(settings)

	consumer := core.NewMessageBrokerSubscriber(
		core.WithSessionSNS(container.SNS),
		core.WithSessionSQS(container.SQS),
		core.WithSubscriberID(settings.Events.ItemsCreated),
		core.WithMaxRetries(3),
		core.WithType(reflect.TypeOf(inventory.ItemsCreatedEvent{})),
		core.WithTopicID(settings.Events.ItemsCreated),
		core.WithHandler(func(payload interface{}) error {
			message := payload.(*inventory.ItemsCreatedEvent)

			logrus.Info("processing received event")

			ctx := context.Background()

			if err := container.InventoryRepository.InsertBulk(ctx, message.ToDomain()); err != nil {
				logrus.
					WithError(err).
					Error("error while inserting items")
				return err
			}

			logrus.Info("items inserted successfully")

			return nil
		}))

	if err := consumer.Run(); err != nil {
		logrus.
			WithError(err).
			Error("shutting down with error")
	}
}
