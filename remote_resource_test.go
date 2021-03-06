package shopify

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"io/ioutil"
	"net/http"
	"testing"
)

type TestRemoteResource struct {
	// err is the error to return on Request calls. Setting this to nil will cause Request to not return an error.
	err error
	// body is the response body.
	body       []byte
	bodyReader io.ReadCloser

	expectedURI  string
	expectedBody string
	t            *testing.T
}

func NewTestRemote(t *testing.T) *TestRemoteResource {
	return &TestRemoteResource{
		t: t,
	}
}

func (tr *TestRemoteResource) ReturnsStatus(status int) *TestRemoteResource {
	return tr
}

func (tr *TestRemoteResource) ReturnsBody(body []byte) *TestRemoteResource {
	tr.body = body
	return tr
}

func (tr *TestRemoteResource) ReturnsBodyReader(reader io.ReadCloser) *TestRemoteResource {
	tr.bodyReader = reader
	return tr
}

func (tr *TestRemoteResource) ExpectsURL(u string) *TestRemoteResource {
	tr.expectedURI = u
	return tr
}

func (tr *TestRemoteResource) Request(req *http.Request) (io.ReadCloser, error) {
	if len(tr.expectedURI) > 0 {
		assert.Equal(tr.t, tr.expectedURI, req.URL.String())
	}

	if len(tr.expectedBody) > 0 {
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			tr.t.Error(err)
		}
		assert.Equal(tr.t, tr.expectedBody, string(body))
	}

	if tr.err != nil {
		return nil, tr.err
	}
	if tr.bodyReader != nil {
		return tr.bodyReader, nil
	}
	return ioutil.NopCloser(bytes.NewReader(tr.body)), nil
}

func TestRequestAndDecodeReturnsErrors(t *testing.T) {
	remote := &ShopifyRemoteJSONResource{
		&ShopifyAdminURLBuilder{},
		&TestRemoteResource{err: fmt.Errorf("Error!")},
	}

	err := remote.RequestAndDecode(nil, "", nil)
	assert.NotNil(t, err)
}

func TestRequestAndDecodeReturnsErrorForEmptyPayload(t *testing.T) {
	json := ``
	remote := &ShopifyRemoteJSONResource{
		&ShopifyAdminURLBuilder{},
		&TestRemoteResource{body: []byte(json)},
	}

	req, _ := http.NewRequest("GET", "some/url", nil)
	var result interface{}
	err := remote.RequestAndDecode(req, "element", result)
	assert.NotNil(t, err)
}

func TestRequestAndDecodeReturnsDecodedPayload(t *testing.T) {
	remote := &ShopifyRemoteJSONResource{
		&ShopifyAdminURLBuilder{},
		&TestRemoteResource{body: []byte(ThemeJSON)},
	}

	req, _ := http.NewRequest("GET", "some/url", nil)
	var result *Theme
	err := remote.RequestAndDecode(req, "theme", &result)
	assert.Nil(t, err)
	assert.Equal(t, 828155753, result.ID)
}
