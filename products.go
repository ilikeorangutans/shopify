package shopify

import (
	"net/http"
)

type Products struct {
	RemoteJSONResource
}

type Product struct {
	CommonFields
	Title string `json:"title"`
}
type ProductList []*Product

func (pl ProductList) Size() int {
	return len(pl)
}

func (pl ProductList) LastID() ShopifyID {

	return 0
}

func trypagination() chan *Product {

	prodChan := make(chan *Product)

	products := &Products{}
	var prods ProductList

	go func() {
		defer close(prodChan)
		pagination := Pagination{}
		for {
			prods, _ = products.List(pagination)
			pagination = pagination.Update(prods)

			for i := range prods {
				product := prods[i]
				prodChan <- product
			}

			if pagination.NoMore {
				break
			}
		}
	}()

	return prodChan
}

func (p *Products) List(pagination Pagination) (ProductList, error) {
	url := p.BuildURL("products")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	pagination.Paginate(req)
	var products []*Product
	err = p.RequestAndDecode(req, "products", &products)
	if err != nil {
		return nil, err
	}

	return products, nil
}
