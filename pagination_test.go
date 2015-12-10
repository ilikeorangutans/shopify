package shopify

import ()

func ExamplePagination() {
	var pagination Paginator

	var products = &Products{}

	for {
		prods, err := products.List(pagination)
		if err != nil {
			return
		}

		pagination = pagination.Update(prods)
		if pagination.NoMore {
			break
		}
	}
}
