package shopify

import (
	"net/http"
)

const DefaultPageSize = 50

type Pagination struct {
	Limit  int
	LastID ShopifyID
	NoMore bool
}

func (p Pagination) Paginate(req *http.Request) {
	// TODO: add query params
}

func (p Pagination) Update(paginateable Paginateable) Pagination {
	return Pagination{
		Limit:  p.Limit,
		LastID: paginateable.LastID(),
		NoMore: paginateable.Size() == 0 || paginateable.Size() < p.Limit,
	}
}

type Paginateable interface {
	Size() int
	LastID() ShopifyID
}
