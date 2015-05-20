package shopify

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type FulfillmentServices struct {
	requester  Requester
	urlBuilder URLBuilder
}

type FulfillmentService struct {
	Email                  string `json:"email"`
	Handle                 string `json:"handle"`
	Name                   string `json:"name"`
	CallbackURL            string `json:"callback_url"`
	Format                 string `json:"format"`
	Credential1            string `json:"credential1"`
	Credential2Exists      bool   `json:"credential2_exists"`
	InventoryManagement    bool   `json:"inventory_management"`
	ProviderId             *int   `json:"provider_id"`
	RequiresShippingMethod bool   `json:"requires_shipping_method"`
	TrackingSupport        bool   `json:"tracking_support"`
}

type FulfillmentServiceWithId struct {
	*FulfillmentService
	Id int `json:"id"`
}

func (ffs *FulfillmentServices) List() []*FulfillmentServiceWithId {
	req, err := http.NewRequest("GET", ffs.urlBuilder("/admin", "fulfillment_services.json?scope=all"), nil)
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		log.Fatal(err)
	}

	resp, err := ffs.requester(req)
	if err != nil {
		log.Fatal(err)
	}

	return decodeServicesJSON(resp["fulfillment_services"])
}

func decodeServicesJSON(data []byte) []*FulfillmentServiceWithId {
	var services []*FulfillmentServiceWithId

	err := json.Unmarshal(data, &services)
	if err != nil {
		log.Fatal(err)
	}
	return services
}

func (ffs *FulfillmentServices) Create(service *FulfillmentService) (*FulfillmentService, error) {
	tmp := make(map[string]interface{})
	tmp["fulfillment_service"] = service
	b, err := json.Marshal(tmp)
	if err != nil {
		log.Fatal(err)
	}

	payload := fmt.Sprintf("%s", b)

	req, err := http.NewRequest("POST", ffs.urlBuilder("/admin", "fulfillment_services.json"), strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		log.Fatal(err)
	}

	resp, err := ffs.requester(req)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp)

	return nil, nil
}
