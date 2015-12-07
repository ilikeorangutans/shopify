package shopify

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type RemoteResource interface {
	Request(*http.Request) (io.ReadCloser, error)
}

type httpRemoteResource struct {
	client   *http.Client
	settings ClientSettings
}

func (rr *httpRemoteResource) Request(req *http.Request) (io.ReadCloser, error) {
	rr.settings.AuthenticateRequest(req)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	resp, err := rr.client.Do(req)
	if err != nil {
		return nil, err
	}

	if err := FromResponse(resp); err != nil {
		return nil, err
	}

	return resp.Body, nil
}

func NewRemoteResource(settings ClientSettings) RemoteResource {
	client := &http.Client{
		Timeout: settings.timeout,
	}
	return &httpRemoteResource{
		client:   client,
		settings: settings,
	}
}

type RemoteJSONResource interface {
	URLBuilder
	RemoteResource
	RequestAndDecode(*http.Request, string, interface{}) error
}

type ShopifyRemoteJSONResource struct {
	URLBuilder
	RemoteResource
}

func (sr *ShopifyRemoteJSONResource) RequestAndDecode(req *http.Request, element string, v interface{}) error {
	reader, err := sr.Request(req)
	if err != nil {
		return err
	}
	defer reader.Close()

	decoder := json.NewDecoder(reader)
	var raw map[string]json.RawMessage
	if err = decoder.Decode(&raw); err != nil {
		return err
	}

	data, found := raw[element]
	if !found {
		return fmt.Errorf("Element \"%s\" could not be found in response from server.", element)
	}

	return json.Unmarshal(data, v)
}
