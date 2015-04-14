package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Channel struct {
	Id         int `json:"id"`
	ShopId     int `json:"shop_id"`
	ProviderId int `json:"provider_id"`
}

type Channels struct {
	requester  Requester
	urlBuilder URLBuilder
}

func (c *Channels) List() []*Channel {
	req, err := http.NewRequest("GET", c.urlBuilder("/admin/channels.json"), nil)
	if err != nil {
		log.Fatal(err)
	}

	d, err := c.requester(req)
	if err != nil {
		log.Fatal(err)
	}

	var ch []*Channel
	json.Unmarshal(d["channels"], &ch)

	return ch
}
