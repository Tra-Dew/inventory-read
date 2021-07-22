package inventory

import (
	"time"
)

// ItemLockCompletedEvent ...
type ItemLockCompletedEvent struct {
	ID       string `json:"id"`
	Quantity int64  `json:"quantity"`
}

// ItemsLockCompletedEvent ...
type ItemsLockCompletedEvent struct {
	Items []*ItemLockCompletedEvent `json:"items"`
}

// ItemCreatedEvent ...
type ItemCreatedEvent struct {
	ID             string    `json:"id"`
	OwnerID        string    `json:"owner_id"`
	Name           string    `json:"name"`
	Description    *string   `json:"description"`
	TotalQuantity  int64     `json:"total_quantity"`
	LockedQuantity int64     `json:"locked_quantity"`
	CreatedAt      time.Time `json:"created_at"`
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
			ID:             item.ID,
			OwnerID:        item.OwnerID,
			Name:           item.Name,
			Description:    item.Description,
			TotalQuantity:  item.TotalQuantity,
			LockedQuantity: item.LockedQuantity,
			CreatedAt:      item.CreatedAt,
		}
	}

	return items
}
