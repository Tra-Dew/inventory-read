package cmd

import (
	"context"
	"reflect"

	"github.com/d-leme/tradew-inventory-read/pkg/core"
	"github.com/d-leme/tradew-inventory-read/pkg/inventory"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// ItemsUpdated ...
func ItemsUpdated(command *cobra.Command, args []string) {
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
		core.WithSubscriberID(settings.Events.ItemsUpdated),
		core.WithMaxRetries(3),
		core.WithType(reflect.TypeOf(inventory.ItemsUpdatedEvent{})),
		core.WithTopicID(settings.Events.ItemsUpdated),
		core.WithHandler(func(payload interface{}) error {
			message := payload.(*inventory.ItemsUpdatedEvent)

			logrus.Info("processing received event")

			ctx := context.Background()

			ids := make([]string, len(message.Items))
			for i, item := range message.Items {
				ids[i] = item.ID
			}

			if err := container.InventoryRepository.UpdateBulk(ctx, message.ToDomain()); err != nil {
				logrus.
					WithError(err).
					Error("error while updating items")
				return err
			}

			logrus.Info("items updated successfully")

			return nil
		}))

	if err := consumer.Run(); err != nil {
		logrus.
			WithError(err).
			Error("shutting down with error")
	}
}
