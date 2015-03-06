package main

import (
	"log"
	"net/http"
	"net/url"
)

type Requester func(req *http.Request) (resp *http.Response, err error)

type URLBuilder func(string) string

type settings struct {
	host, username, password string
}

type ShopifyClient struct {
	settings settings
	client   *http.Client
}

func Connect(host, username, password string) *ShopifyClient {
	client := &http.Client{}

	return &ShopifyClient{client: client, settings: settings{host: host, username: username, password: password}}
}

func (sc *ShopifyClient) Webhooks() *Webhooks {
	return &Webhooks{requester: sc.doRequest, urlBuilder: sc.buildURL}
}

func (sc *ShopifyClient) doRequest(req *http.Request) (resp *http.Response, err error) {
	req.SetBasicAuth(sc.settings.username, sc.settings.password)

	return sc.client.Do(req)
}

func (sc *ShopifyClient) buildURL(input string) string {
	url, err := url.Parse("https://" + sc.settings.host + input)
	if err != nil {
		log.Fatal(err)
	}

	return url.String()
}
