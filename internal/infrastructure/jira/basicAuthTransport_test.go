package jira

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	// testMux is the HTTP request multiplexer used with the test server.
	testMux *http.ServeMux

	// testServer is a test HTTP server used to provide mock API responses.
	testServer *httptest.Server
)

// setup sets up a test HTTP server along with a jira.Client that is configured to talk to that test server.
// Tests should register handlers on mux which provide mock responses for the API method being tested.
func setup() {
	// Test server
	testMux = http.NewServeMux()
	testServer = httptest.NewServer(testMux)

}

// teardown closes the test HTTP server.
func teardown() {
	testServer.Close()
}

func TestBasicAuthTransport(t *testing.T) {

	t.Run("Client should have a correct BasicAuthTransport struct", func(t *testing.T) {

		tp := &BasicAuthTransport{
			Username: "username",
			APIToken: "token",
		}

		client := tp.Client()

		if client.Transport == nil {
			t.Error("Transport should not be nil")
		}

		_, ok := client.Transport.(*BasicAuthTransport)

		if !ok {
			t.Error("Transport should be a BasicAuthTransport")
		}

	})

	t.Run("BasicAuthTransport should have a correct RoundTrip method", func(t *testing.T) {

		setup()
		defer teardown()

		username, token := "username", "apiToken"

		testMux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

			u, p, ok := r.BasicAuth()

			if !ok {
				t.Error("Basic auth should be set")
			}

			if u != username {
				t.Errorf("Username should be %s, got %s", username, u)
			}

			if p != token {
				t.Errorf("Token should be %s, got %s", token, p)
			}

		})

		tp := &BasicAuthTransport{
			Username: username,
			APIToken: token,
		}

		jiraClient, _ := NewClient(testServer.URL, tp.Client())

		req, _ := jiraClient.client.NewRequest(context.Background(), "GET", ".", nil)

		jiraClient.client.Do(req)

	})

	t.Run("Should use http.DefaultTransport if transport is nil", func(t *testing.T) {

		tp := &BasicAuthTransport{
			Username: "username",
			APIToken: "apiToken",
		}

		if tp.transport() != http.DefaultTransport {
			t.Error("Transport should be http.DefaultTransport")
		}

		customtp := &BasicAuthTransport{
			Username:  "username",
			APIToken:  "apiToken",
			Transport: &http.Transport{},
		}

		if customtp.transport() == http.DefaultTransport {
			t.Error("Transport should not be http.DefaultTransport")
		}

	})
}
