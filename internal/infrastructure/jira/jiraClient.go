package jira

import (
	"net/http"

	httpclient "github.com/NexusGL98/trackmeup/internal/infrastructure/http"
)

type JiraClient struct {
	client *httpclient.Client
}

func NewClient(baseUrl string, httpClient *http.Client) (*JiraClient, error) {

	client, err := httpclient.NewClient(baseUrl, httpClient)

	if err != nil {
		return nil, err
	}

	return &JiraClient{
		client: client,
	}, nil
}
