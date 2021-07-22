package inventory

import (
	"context"
	"time"
)

// Item ...
type Item struct {
	ID             string     `json:"id" bson:"_id"`
	OwnerID        string     `json:"owner_id" bson:"owner_id"`
	Name           string     `json:"name" bson:"name"`
	Description    *string    `json:"description" bson:"description"`
	TotalQuantity  int64      `json:"total_quantity" bson:"total_quantity"`
	LockedQuantity int64      `json:"locked_quantity" bson:"locked_quantity"`
	CreatedAt      time.Time  `json:"created_at" bson:"created_at"`
	UpdatedAt      *time.Time `json:"updated_at" bson:"updated_at"`
}

// GetItemsRequest ...
type GetItemsRequest struct {
	Token    *string `form:"token"`
	PageSize int64   `form:"page_size"`
}

// GetItemsResponse ...
type GetItemsResponse struct {
	Items []*Item `json:"items"`
	Token string  `json:"token"`
}

// Repository ...
type Repository interface {
	Get(ctx context.Context, req *GetItemsRequest) (*GetItemsResponse, error)
	GetByID(ctx context.Context, userID string, id string) ([]*Item, error)
	GetUserItems(ctx context.Context, userID string, req *GetItemsRequest) (*GetItemsResponse, error)

	InsertBulk(ctx context.Context, items []*Item) error
	UpdateBulk(ctx context.Context, items []*Item) error
}
