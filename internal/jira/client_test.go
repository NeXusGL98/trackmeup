package jira

import (
	"bytes"
	"io"
	"net/http"
	"testing"
)

const (
	testUrl = "https://jira.atlassian.com"
)

type fakeHttpTransport func(req *http.Request) (*http.Response, error)

func (t fakeHttpTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	return t(req)
}

func TestClient(t *testing.T) {

	t.Run("Should not return an error", func(t *testing.T) {

		customClient := &http.Client{
			Transport: fakeHttpTransport(func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: 200,
					Body:       io.NopCloser(bytes.NewBufferString(`{"key":"value"}`)),
				}, nil
			}),
		}

		_, err := NewClient(testUrl, customClient)

		if err != nil {
			t.Error("Expected no error, got:", err)
		}
	})

}
