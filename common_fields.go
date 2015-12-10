package shopify

import (
	"fmt"
	"time"
)

type Timestamps struct {
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

type ShopifyID int64

func (id ShopifyID) String() string {
	return fmt.Sprintf("%d", id)
}

type Identifiable interface {
	ID() ShopifyID
}

type CommonFields struct {
	ID     ShopifyID `json:"id"`
	ShopID int64     `json:"shop_id"`
	Timestamps
}
