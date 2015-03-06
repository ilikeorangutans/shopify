package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

type WebhookResponse struct {
	Webhooks []Webhook `json:"webhooks"`
}

type Webhook struct {
	Format  string `json:"format"`
	Id      int    `json:"id"`
	Topic   string `json:"topic"`
	Address string `json:"address"`
}

type Webhooks struct {
	requester  Requester
	urlBuilder URLBuilder
}

func (webhooks *Webhooks) create(topic string, address *url.URL, format string) {

	payload := fmt.Sprintf("{\"webhook\":{\"topic\":\"%s\", \"address\":\"%s\", \"format\": \"%s\"}}", topic, address.String(), format)
	req, err := http.NewRequest("POST", webhooks.urlBuilder("/admin/webhooks.json"), strings.NewReader(payload))
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := webhooks.requester(req)
	if err != nil {
		log.Fatal(err)
	}

	body, _ := httputil.DumpResponse(resp, true)
	log.Printf("RESPONSE: %s \n", body)

	webhooks.requester(req)
}

func (webhooks *Webhooks) list() []Webhook {
	req, err := http.NewRequest("GET", webhooks.urlBuilder("/admin/webhooks.json"), nil)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := webhooks.requester(req)
	if err != nil {
		log.Fatal(err)
	}

	decoder := json.NewDecoder(resp.Body)

	var m *WebhookResponse
	decoder.Decode(&m)

	return m.Webhooks
}
