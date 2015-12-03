package shopify

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
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
	assert.Equal(t, "Server responded with \"401 Unauthorized\", check credentials", err.Error())
}

func TestConnectWithInternalServerError(t *testing.T) {
	ts := ErrorReturningTestServer(http.StatusInternalServerError)
	defer ts.Close()

	client := NewClient(ts.URL, "", "")
	err := client.Connect()
	assert.NotNil(t, err)
	assert.Equal(t, "Error connecting to server: 500 Internal Server Error", err.Error())
}

func TestConnectWithTimeout(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(time.Duration(100 * time.Millisecond))
	}))
	defer ts.Close()
	settings := ClientSettings{host: ts.URL, username: "", password: "", timeout: time.Duration(100 * time.Millisecond)}
	client := NewClientWithSettings(settings)
	client.Settings.timeout = time.Duration(1 * time.Millisecond)

	err := client.Connect()

	assert.NotNil(t, err)
}
