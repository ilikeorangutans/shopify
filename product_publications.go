package shopify

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type ProductPublication struct {
	ChannelId int `json:"channel_id"`
	Id        int `json:"id"`
	ShopId    int `json:"shop_id"`
	ProductId int `json:"product_id"`
}

type ProductPublications struct {
	requester  Requester
	urlBuilder URLBuilder
}

func (pp *ProductPublications) Get(channel_id, id int) (*ProductPublication, error) {
	req, err := http.NewRequest("GET", pp.urlBuilder(fmt.Sprintf("/admin/channels/%d/product_publications/%d", channel_id, id)), nil)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := pp.requester(req)
	if err != nil {
		log.Fatal(err)
	}

	var productPublication *ProductPublication
	json.Unmarshal(resp["product_publication"], &productPublication)

	return productPublication, nil
}

func (pp *ProductPublications) List(channel_id int) []*ProductPublication {
	req, err := http.NewRequest("GET", pp.urlBuilder(fmt.Sprintf("/admin/channels/%d/product_publications.json", channel_id)), nil)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := pp.requester(req)
	if err != nil {
		log.Fatal(err)
	}

	return pp.parseList(resp["product_publications"])
}

func (pp *ProductPublications) Delete(publication ProductPublication) error {

	return nil
}

func (pp *ProductPublications) parseList(body []byte) []*ProductPublication {
	var productPublications []*ProductPublication
	json.Unmarshal(body, &productPublications)

	return productPublications
}
