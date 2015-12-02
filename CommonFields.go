package shopify

import (
	"time"
)

type CommonFields struct {
	ID        int       `json:"id"`
	ShopID    int       `json:"shop_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
