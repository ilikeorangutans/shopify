package shopify

import (
	"time"
)

type Timestamps struct {
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

type ShopifyID int64

type Identifiable interface {
	ID() ShopifyID
}

type CommonFields struct {
	ID     int64 `json:"id"`
	ShopID int64 `json:"shop_id"`
	Timestamps
}
