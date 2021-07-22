package inventory

import (
	"time"
)

// ItemLockCompletedEvent ...
type ItemLockCompletedEvent struct {
	ID        string    `json:"id"`
	Quantity  int64     `json:"quantity"`
	UpdatedAt time.Time `json:"updated_at"`
}

// ItemsLockCompletedEvent ...
type ItemsLockCompletedEvent struct {
	Items []*ItemLockCompletedEvent `json:"items"`
}

// ItemCreatedEvent ...
type ItemCreatedEvent struct {
	ID            string    `json:"id"`
	OwnerID       string    `json:"owner_id"`
	Name          string    `json:"name"`
	Description   *string   `json:"description"`
	TotalQuantity int64     `json:"total_quantity"`
	CreatedAt     time.Time `json:"created_at"`
}

// ItemsCreatedEvent ...
type ItemsCreatedEvent struct {
	Items []*ItemCreatedEvent `json:"items"`
}

// ItemUpdatedEvent ...
type ItemUpdatedEvent struct {
	ID            string    `json:"id"`
	OwnerID       string    `json:"owner_id"`
	Name          string    `json:"name"`
	Description   *string   `json:"description"`
	TotalQuantity int64     `json:"total_quantity"`
	UpdatedAt     time.Time `json:"updated_at"`
}

// ItemsUpdatedEvent ...
type ItemsUpdatedEvent struct {
	Items []*ItemUpdatedEvent `json:"items"`
}

// ToDomain ...
func (event *ItemsCreatedEvent) ToDomain() []*Item {
	items := make([]*Item, len(event.Items))

	for i, item := range event.Items {
		items[i] = &Item{
			ID:            item.ID,
			OwnerID:       item.OwnerID,
			Name:          item.Name,
			Description:   item.Description,
			TotalQuantity: item.TotalQuantity,
			CreatedAt:     item.CreatedAt,
		}
	}

	return items
}

// ToDomain ...
func (event *ItemUpdatedEvent) ToDomain(item *Item) *Item {
	return &Item{
		ID:             event.ID,
		OwnerID:        event.OwnerID,
		Name:           event.Name,
		Description:    event.Description,
		TotalQuantity:  event.TotalQuantity,
		LockedQuantity: item.LockedQuantity,
		CreatedAt:      item.CreatedAt,
		UpdatedAt:      &event.UpdatedAt,
	}
}

// ToDomain ...
func (event *ItemsUpdatedEvent) ToDomain(items []*Item) []*Item {
	updatedItems := make([]*Item, len(event.Items))

	for i, eventItem := range event.Items {

		var foundItem *Item
		for _, item := range items {
			if item.ID == eventItem.ID {
				foundItem = item
			}
		}

		updatedItems[i] = eventItem.ToDomain(foundItem)
	}

	return updatedItems
}

// ToDomain ...
func (event *ItemLockCompletedEvent) ToDomain(item *Item) *Item {
	return &Item{
		ID:             event.ID,
		LockedQuantity: event.Quantity,
		UpdatedAt:      &event.UpdatedAt,

		OwnerID:       item.OwnerID,
		Name:          item.Name,
		Description:   item.Description,
		TotalQuantity: item.TotalQuantity,
		CreatedAt:     item.CreatedAt,
	}
}

// ToDomain ...
func (event *ItemsLockCompletedEvent) ToDomain(items []*Item) []*Item {
	updatedItems := make([]*Item, len(event.Items))

	for i, eventItem := range event.Items {

		var foundItem *Item
		for _, item := range items {
			if item.ID == eventItem.ID {
				foundItem = item
			}
		}

		updatedItems[i] = eventItem.ToDomain(foundItem)
	}

	return updatedItems
}
