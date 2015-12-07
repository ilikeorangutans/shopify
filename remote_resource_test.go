package shopify

import (
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"strings"
	"testing"
)

type TestingRemoteResource struct {
	t   *testing.T
	url string
}

func NewAssertingRemoteJSONResource(t *testing.T, url string, statusCode int, response []byte) RemoteJSONResource {
	return &TestingRemoteResource{
		t:   t,
		url: url,
	}
}

func (tr *TestingRemoteResource) Request(req *http.Request) (io.ReadCloser, error) {
	assert.Equal(tr.t, t.url, req.RequestURI)
	return nil, nil
}

func (tr *TestingRemoteResource) BuildURL(segments ...string) string {
	url := strings.Join(segments, "/")
	return url
}
