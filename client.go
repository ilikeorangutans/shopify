package shopify

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

type Requester func(req *http.Request) (map[string]json.RawMessage, error)

type URLBuilder func(...string) string

type settings struct {
	host, username, password string
}

type Client struct {
	settings settings
	client   *http.Client
	Verbose  bool
}

func Connect(host, username, password string) *Client {
	client := &http.Client{}

	resp, err := client.Head(fmt.Sprintf("https://%s/admin/", host))
	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode != 200 {
		log.Fatalf("Server doesn't seem to be up: \"%s\"\n", resp.Status)
	}

	return &Client{client: client, settings: settings{host: host, username: username, password: password}}
}

func (sc *Client) Webhooks() *Webhooks {
	return &Webhooks{requester: sc.doRequest, urlBuilder: sc.buildURL}
}

func (sc *Client) Apps() *APIPermissions {
	return &APIPermissions{requester: sc.doRequest, urlBuilder: sc.buildURL}
}

func (sc *Client) Metafields() *Metafields {
	return &Metafields{requester: sc.doRequest, urlBuilder: sc.buildURL}
}

func (sc *Client) FullfillmentServices() *FulfillmentServices {
	return &FulfillmentServices{requester: sc.doRequest, urlBuilder: sc.buildURL}
}

func (sc *Client) Orders() *Orders {
	return &Orders{requester: sc.doRequest, urlBuilder: sc.buildURL}
}

func (sc *Client) Transactions() *Transactions {
	return &Transactions{requester: sc.doRequest, urlBuilder: sc.buildURL}
}

func (sc *Client) debug(msg string) {
	if sc.Verbose {
		log.Println(msg)
	}
}

func (sc *Client) doRequest(req *http.Request) (map[string]json.RawMessage, error) {

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

func (sc *Client) buildURL(input ...string) string {
	url, err := url.Parse("https://" + sc.settings.host + strings.Join(input, "/"))
	if err != nil {
		log.Fatal(err)
	}

	return url.String()
}
