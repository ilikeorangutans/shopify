package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type Requester func(req *http.Request) (map[string]json.RawMessage, error)

type URLBuilder func(string) string

type settings struct {
	host, username, password string
}

type ShopifyClient struct {
	settings settings
	client   *http.Client
	Verbose  bool
}

func Connect(host, username, password string) *ShopifyClient {
	client := &http.Client{}

	resp, err := client.Head(fmt.Sprintf("https://%s/admin/", host))
	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode != 200 {
		log.Fatalf("Server doesn't seem to be up: \"%s\"\n", resp.Status)
	}

	return &ShopifyClient{client: client, settings: settings{host: host, username: username, password: password}}
}

func (sc *ShopifyClient) Channels() *Channels {
	return &Channels{requester: sc.doRequest, urlBuilder: sc.buildURL}
}

func (sc *ShopifyClient) Webhooks() *Webhooks {
	return &Webhooks{requester: sc.doRequest, urlBuilder: sc.buildURL}
}

func (sc *ShopifyClient) debug(msg string) {
	if sc.Verbose {
		log.Println(msg)
	}
}

func (sc *ShopifyClient) doRequest(req *http.Request) (map[string]json.RawMessage, error) {

	sc.debug(fmt.Sprintf("%s: %s \n", req.Method, req.URL))
	req.SetBasicAuth(sc.settings.username, sc.settings.password)

	resp, err := sc.client.Do(req)
	if err != nil {
		return nil, err
	}

	b, _ := httputil.DumpResponse(resp, true)
	sc.debug(fmt.Sprintf("Response: \n%s", b))

	if resp.StatusCode < 200 || resp.StatusCode > 399 {
		log.Fatal(resp.Status)
	}

	decoder := json.NewDecoder(resp.Body)

	var d map[string]json.RawMessage

	decoder.Decode(&d)

	return d, nil
}

func (sc *ShopifyClient) buildURL(input string) string {
	url, err := url.Parse("https://" + sc.settings.host + input)
	if err != nil {
		log.Fatal(err)
	}

	return url.String()
}
