package inventory

import (
	"time"
)

// ItemUpdatedEvent ...
type ItemUpdatedEvent struct {
	ID             string     `json:"id"`
	OwnerID        string     `json:"owner_id"`
	Name           string     `json:"name"`
	Description    *string    `json:"description"`
	TotalQuantity  int64      `json:"total_quantity"`
	LockedQuantity int64      `json:"locked_quantity"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      *time.Time `json:"updated_at"`
}

// ItemsUpdatedEvent ...
type ItemsUpdatedEvent struct {
	Items []*ItemUpdatedEvent `json:"items"`
}

// ToDomain ...
func (event *ItemUpdatedEvent) ToDomain() *Item {
	return &Item{
		ID:             event.ID,
		OwnerID:        event.OwnerID,
		Name:           event.Name,
		Description:    event.Description,
		TotalQuantity:  event.TotalQuantity,
		LockedQuantity: event.LockedQuantity,
		CreatedAt:      event.CreatedAt,
		UpdatedAt:      event.UpdatedAt,
	}
}

// ToDomain ...
func (event *ItemsUpdatedEvent) ToDomain() []*Item {
	updatedItems := make([]*Item, len(event.Items))

	for i, eventItem := range event.Items {
		updatedItems[i] = eventItem.ToDomain()
	}

	return updatedItems
}
