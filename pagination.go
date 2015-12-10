package shopify

import (
	"fmt"
	"net/http"
	"strconv"
)

// DefaultPageSize is how many entities are retrieved with a request per default.
const DefaultPageSize = 50

// Paginator allows paginated retrieval of entities
type Paginator struct {
	// How many entities should be retrieved with each request?
	Limit int
	// LastID is the last id in the most recent request.
	LastID ShopifyID
	// NoMore is set if there are no more entities to retrieve.
	NoMore bool
	// Total is how many entities have been encountered so far.
	Total int
}

// Paginate updates the given request with the necessary query parameters.
func (p Paginator) Paginate(req *http.Request) {
	q := req.URL.Query()
	effectiveLimit := DefaultPageSize
	if p.Limit > 0 {
		effectiveLimit = p.Limit
	}
	q.Set("limit", strconv.Itoa(effectiveLimit))
	if p.LastID != 0 {
		q.Set("since_id", fmt.Sprintf("%d", p.LastID))
	}
	req.URL.RawQuery = q.Encode()
}

// Update returns a new and updated paginator to request the next set of entities.
func (p Paginator) Update(paginateable Paginateable) Paginator {
	return Paginator{
		Limit:  p.Limit,
		LastID: paginateable.LastID(),
		NoMore: paginateable.Size() == 0 || paginateable.Size() < p.Limit,
		Total:  p.Total + paginateable.Size(),
	}
}

// Paginateable describes anything that can be paginated over. Collections that support pagination need to
// provide this interface.
type Paginateable interface {
	Size() int
	LastID() ShopifyID
}
