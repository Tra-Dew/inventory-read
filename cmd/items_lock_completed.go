package cmd

import (
	"context"
	"reflect"

	"github.com/Tra-Dew/inventory-read/pkg/core"
	"github.com/Tra-Dew/inventory-read/pkg/inventory"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// ItemsLockCompleted ...
func ItemsLockCompleted(command *cobra.Command, args []string) {
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
		core.WithSubscriberID(settings.Events.ItemsLockCompleted),
		core.WithMaxRetries(3),
		core.WithType(reflect.TypeOf(inventory.ItemsLockCompletedEvent{})),
		core.WithTopicID(settings.Events.ItemsLockCompleted),
		core.WithHandler(func(payload interface{}) error {
			message := payload.(*inventory.ItemsLockCompletedEvent)

			logrus.Info("processing received event")

			ctx := context.Background()

			ids := make([]string, len(message.Items))
			for i, item := range message.Items {
				ids[i] = item.ID
			}

			items, err := container.InventoryRepository.GetByIDs(ctx, ids)

			if err != nil {
				logrus.
					WithError(err).
					Error("error while getting items")
				return err
			}

			logrus.Infof("found %d to update", len(items))

			if err := container.InventoryRepository.UpdateBulk(ctx, message.ToDomain(items)); err != nil {
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
