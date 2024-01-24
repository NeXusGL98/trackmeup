package jira

import "net/http"

type BasicAuthTransport struct {
	Username string
	APIToken string

	// Transport is the underlying HTTP transport to use when making requests.
	// It will default to http.DefaultTransport if nil.
	Transport http.RoundTripper
}

func (t *BasicAuthTransport) RoundTrip(req *http.Request) (*http.Response, error) {

	// Clone the request to avoid mutating the passed in request.
	clonedReq := req.Clone(req.Context())

	// Set the Authorization header.

	clonedReq.SetBasicAuth(t.Username, t.APIToken)

	// Make the HTTP request.
	return t.transport().RoundTrip(clonedReq)

}

func (t *BasicAuthTransport) Client() *http.Client {

	return &http.Client{Transport: t}

}

func (t *BasicAuthTransport) transport() http.RoundTripper {

	if t.Transport == nil {
		return http.DefaultTransport
	}

	return t.Transport
}
