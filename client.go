package shopify

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"time"
)

const DefaultTimeout = time.Duration(10 * time.Second)

type Requester func(req *http.Request) (map[string]json.RawMessage, error)

// Client is the facade for all API connections to Shopify. Obtain a new instance with the NewClient function.
type Client struct {
	RemoteJSONResource
	Settings ClientSettings
	client   *http.Client
	Verbose  bool
}

// NewClient returns a new client with default settings and the given host, username and password.
func NewClient(host, username, password string) *Client {
	settings := ClientSettings{host: host, username: username, password: password, timeout: DefaultTimeout}
	return NewClientWithSettings(settings)
}

// NewClientWithSettings creates a new client with the given settings.
func NewClientWithSettings(settings ClientSettings) *Client {
	client := &http.Client{
		Timeout: settings.timeout,
	}
	return &Client{
		client:   client,
		Settings: settings,
		RemoteJSONResource: &ShopifyRemoteJSONResource{
			URLBuilder:     &ShopifyAdminURLBuilder{},
			RemoteResource: NewRemoteResource(settings),
		},
	}
}

func NewClientWithSettingsAndRemoteResource(settings ClientSettings, remote RemoteJSONResource) *Client {
	return &Client{
		Settings:           settings,
		RemoteJSONResource: remote,
	}
}

// Connect attempts a connection to the configured server. If the server responds with a 4xx or 5xx
// status code, this method will return an error. It's recommended to use this function at least
// once after creating a new client to ensure valid settings and credentials.
func (c *Client) Connect() error {

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/admin/", c.Settings.ShopURL()), nil)
	if err != nil {
		return err
	}
	resp, err := c.client.Do(req)
	if err != nil {
		return fmt.Errorf("Error connecting to server: \"%s\"", err.Error())
	} else if resp.StatusCode >= 400 && resp.StatusCode < 500 {
		return fmt.Errorf("Server responded with \"%s\", check credentials", resp.Status)
	} else if resp.StatusCode >= 500 {
		return fmt.Errorf("Error connecting to server: %s", resp.Status)
	}

	return nil
}

func (c *Client) RequestAndDecode(r *http.Request, name string, v interface{}) error {
	return nil
}

func (c *Client) Webhooks() *Webhooks {
	return &Webhooks{RemoteJSONResource: c}
}

func (c *Client) Apps() *APIPermissions {
	return &APIPermissions{RemoteJSONResource: c}
}

func (c *Client) Metafields() *Metafields {
	return &Metafields{RemoteJSONResource: c}
}

func (c *Client) FullfillmentServices() *FulfillmentServices {
	return &FulfillmentServices{RemoteJSONResource: c}
}

func (c *Client) Orders() *Orders {
	return &Orders{RemoteJSONResource: c}
}

func (c *Client) Transactions() *Transactions {
	return &Transactions{RemoteJSONResource: c}
}

func (c *Client) Themes() *Themes {
	return &Themes{RemoteJSONResource: c}
}

func (c *Client) Assets(theme *Theme) *Assets {
	return &Assets{
		RemoteJSONResource: c,
		Theme:              theme,
	}
}

func (c *Client) debug(msg string) {
	if c.Verbose {
		log.Println(msg)
	}
}

func (c *Client) Request(req *http.Request) (io.ReadCloser, error) {
	c.debug(fmt.Sprintf("%s: %s \n", req.Method, req.URL))
	c.Settings.AuthenticateRequest(req)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	b, _ := httputil.DumpResponse(resp, true)
	c.debug(fmt.Sprintf("Response: \n%s", b))

	if err := FromResponse(resp); err != nil {
		return nil, err
	}

	return resp.Body, nil
}
