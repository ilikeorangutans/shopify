package shopify

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"
)

func ErrorReturningTestServer(statusCode int) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "error", statusCode)
	}))
}

func TestConnectOK(t *testing.T) {
	ts := ErrorReturningTestServer(http.StatusOK)
	defer ts.Close()

	client := NewClient(ts.URL, "valid", "valid")
	err := client.Connect()

	assert.Nil(t, err)
}

func TestConnectWithInvalidCredentials(t *testing.T) {
	ts := ErrorReturningTestServer(http.StatusUnauthorized)
	defer ts.Close()

	client := NewClient(ts.URL, "invalid", "invalid")
	err := client.Connect()
	assert.NotNil(t, err)
	assert.True(t, strings.Contains(err.Error(), strconv.Itoa(http.StatusUnauthorized)))
}

func TestConnectWithInternalServerError(t *testing.T) {
	ts := ErrorReturningTestServer(http.StatusInternalServerError)
	defer ts.Close()

	client := NewClient(ts.URL, "", "")
	err := client.Connect()
	assert.NotNil(t, err)
	assert.True(t, strings.Contains(err.Error(), strconv.Itoa(http.StatusInternalServerError)))
}

func TestConnectWithTimeout(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(time.Duration(100 * time.Millisecond))
	}))
	defer ts.Close()
	settings := ClientSettings{Host: ts.URL, Username: "", Password: "", Timeout: time.Duration(100 * time.Millisecond)}
	client := NewClientWithSettings(settings)
	client.Settings.Timeout = time.Duration(1 * time.Millisecond)

	err := client.Connect()

	assert.NotNil(t, err)
}
