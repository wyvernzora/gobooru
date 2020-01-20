package client

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/wyvernzora/gobooru/types"
)

const UserAgent string = "gobooru/0.1"

// BaseClient provides basic authentication and host resolution for all type-specific
// client types
type BaseClient struct {
	HTTPClient *http.Client
	Host       string
	User       string
	APIKey     string
}

func NewBaseClient(host string, user string, apiKey string) (*BaseClient, error) {
	if host == "" {
		return nil, errors.New("host cannot be empty")
	}

	return &BaseClient{
		HTTPClient: &http.Client{},
		Host:       host,
		User:       user,
		APIKey:     apiKey,
	}, nil
}

func (client *BaseClient) NewRequest(method string, path string) (*http.Request, error) {
	url := "https://" + client.Host + path
	request, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	request.SetBasicAuth(client.User, client.APIKey)
	request.Header.Add("User-Agent", UserAgent)
	return request, nil
}

func (client *BaseClient) ExecuteRequest(request *http.Request, result interface{}) error {
	log.Printf("Executing request: %s\n", request.URL.String())

	resp, err := client.HTTPClient.Do(request)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode/100 == 2 {
		return json.Unmarshal(body, result)
	}

	var danbooruError types.Error
	err = json.Unmarshal(body, &danbooruError)
	if err != nil {
		return err
	}
	return &danbooruError
}
