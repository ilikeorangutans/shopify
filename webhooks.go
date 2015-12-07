package shopify

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type Webhook struct {
	CommonFields

	Format  string `json:"format"`
	Topic   string `json:"topic"`
	Address string `json:"address"`
}

type Webhooks struct {
	RemoteJSONResource
}

func (webhooks *Webhooks) Create(topic string, address *url.URL, format string) (*Webhook, error) {

	payload := fmt.Sprintf("{\"webhook\":{\"topic\":\"%s\", \"address\":\"%s\", \"format\": \"%s\"}}", topic, address.String(), format)
	req, err := http.NewRequest("POST", webhooks.BuildURL("/admin/webhooks.json"), strings.NewReader(payload))
	if err != nil {
		log.Fatal(err)
	}

	var webhook *Webhook
	err = webhooks.RequestAndDecode(req, "webhook", &webhook)
	if err != nil {
		return nil, err
	}

	return webhook, nil
}

func (webhooks *Webhooks) Delete(id int64) {
	webhook, err := webhooks.Get(id)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Deleting webhook [%d] %s %s %s", webhook.ID, webhook.Topic, webhook.Format, webhook.Address)

	req, err := http.NewRequest("POST", webhooks.BuildURL(fmt.Sprintf("/admin/webhooks/%d.json", id)), strings.NewReader("_method=delete"))
	if err != nil {
		log.Fatal(err)
	}

	_, err = webhooks.Request(req)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Webhook %d deleted\n", id)
}

func (ws *Webhooks) List() ([]*Webhook, error) {
	req, err := http.NewRequest("GET", ws.BuildURL("/admin/webhooks.json"), nil)
	if err != nil {
		return nil, err
	}

	var webhooks []*Webhook
	err = ws.RequestAndDecode(req, "webhooks", &webhooks)
	if err != nil {
		return nil, err
	}

	return webhooks, nil
}

func (ws *Webhooks) Get(id int64) (*Webhook, error) {
	req, err := http.NewRequest("GET", ws.BuildURL(fmt.Sprintf("/admin/webhooks/%d.json", id)), nil)
	if err != nil {
		return nil, err
	}
	var webhook *Webhook
	err = ws.RequestAndDecode(req, "webhook", &webhook)
	if err != nil {
		return nil, err
	}

	return webhook, nil
}
