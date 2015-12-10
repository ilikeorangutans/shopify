package shopify

import (
	"net/http"
)

type Products struct {
	RemoteJSONResource
}

type Product struct {
	CommonFields
	Title    string     `json:"title"`
	Handle   string     `json:"handle"`
	Variants []*Variant `json:"variants"`
}

type Variant struct {
	CommonFields
	Title string `json:"title"`
}

func (p *Products) List(paginator Paginator) (ProductList, error) {
	url := p.BuildURL("products")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	paginator.Paginate(req)
	var products []*Product
	err = p.RequestAndDecode(req, "products", &products)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (p *Products) Get(id ShopifyID) (*Product, error) {
	req, err := http.NewRequest("GET", p.BuildURL("products", id.String()), nil)
	if err != nil {
		return nil, err
	}
	var prod *Product
	if err := p.RequestAndDecode(req, "product", &prod); err != nil {
		return nil, err
	}
	return prod, nil
}

type ProductList []*Product

func (pl ProductList) Size() int {
	return len(pl)
}

func (pl ProductList) LastID() ShopifyID {
	if pl.Size() == 0 {
		return 0
	}

	return pl[len(pl)-1].ID
}
