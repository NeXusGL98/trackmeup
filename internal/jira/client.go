package jira

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type Client struct {
	baseURL *url.URL
	client  *http.Client
}

// NewClient returns a new Jira client with the given base URL and HTTP client.
// if a nil httpClient is provided, http.DefaultClient will be used
// baseURL should always be specified with a trailing slash
// baseURL is the http url of your jira instance
func NewClient(baseURL string, httpClient *http.Client) (*Client, error) {

	baseEndpoint, err := url.Parse(baseURL)

	if err != nil {
		return nil, err
	}

	// make sure the baseURL ends with a slash to avoid having to do it everywhere else
	if !strings.HasSuffix(baseURL, "/") {
		baseEndpoint.Path += "/"
	}

	// If the client is nil, use the default HTTP client.
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	return &Client{
		baseURL: baseEndpoint,
		client:  httpClient,
	}, nil
}

func (c *Client) NewRequest(ctx context.Context, method, urlString string, body interface{}) (*http.Request, error) {

	url, err := url.Parse(urlString)

	if err != nil {
		return nil, err
	}

	// Urls should not contain a leading slash since the baseURL already includes it.
	url.Path = strings.TrimLeft(url.Path, "/")

	u := c.baseURL.ResolveReference(url)

	var buf io.ReadWriter

	if body != nil {
		buf = new(bytes.Buffer)
		err = json.NewEncoder(buf).Encode(body)

		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequestWithContext(ctx, method, u.String(), buf)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	return req, nil

}

func (c *Client) Do(req *http.Request) (*http.Response, error) {

	return c.client.Do(req)

}
