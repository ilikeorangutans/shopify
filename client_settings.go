package shopify

import (
	"fmt"
	"net/http"
	"regexp"
	"time"
)

type AuthenticateRequest func(*http.Request)

// ClientSettings hold all kinds of configuration options for Clients.
type ClientSettings struct {
	Host, Username, Password string
	Timeout                  time.Duration
	// DumpRequestURLs causes clients to log every URL that's accessed.
	DumpRequestURLs bool
	// DumpRequests and DumpResponses causes clients to log HTTP requests and responses.
	DumpRequests, DumpResponses bool
}

func (cs ClientSettings) ShopURL() string {
	pattern := regexp.MustCompile("https?://")
	if pattern.MatchString(cs.Host) {
		return cs.Host
	} else {
		return fmt.Sprintf("https://%s", cs.Host)
	}
}

// AuthenticateRequest adds authentication data to the given request
func (cs ClientSettings) AuthenticateRequest(req *http.Request) {
	if len(cs.Username) > 0 && len(cs.Password) > 0 {
		req.SetBasicAuth(cs.Username, cs.Password)
	}
}

type RequestAuthenticator interface {
	AuthenticateRequest(*http.Request)
}
