package azure

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type Client interface {
	authenticate() (string, error)
	listSubscriptions(token string) ([]Subscription, error)
	getResources(subscription Subscription, token string) ([]Resource, error)
}

const (
	authentication = "/%s/oauth2/token"

)

type client struct {
	httpClient   *http.Client
	clientId     string
	clientSecret string
	tenantId     string
	loginBaseURL string
	azureBaseURL string
}

func (c *client) getResources(subscription Subscription, token string) ([]Resource, error) {
	panic("implement me")
}

func (c *client) authenticate() (string, error) {
	payload := url.Values{}

	payload.Set("grant_type", "client_credentials")
	payload.Set("resource", c.azureBaseURL)
	payload.Set("client_id", c.clientId)
	payload.Set("client_secret", c.clientSecret)

	req, err := http.NewRequest("POST", c.loginBaseURL+fmt.Sprintf(authentication,c.tenantId), strings.NewReader(payload.Encode()))
	if err != nil {
		log.Error("Error creating request.", err)
		return "", err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err := c.httpClient.Do(req)

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	if err != nil {
		log.Error("Error during authentication.", err)
		return "", err
	}
	if http.StatusOK != resp.StatusCode {
		log.Error("Authentication failed.")
		return "", fmt.Errorf("Authentication failed. Http Status %d ", resp.StatusCode)
	}

	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error("Error reading response", err)
		return "", err
	}

	jsonMap := make(map[string]interface{})
	err = json.Unmarshal(response, &jsonMap)
	if err != nil {
		log.Error("Error parsing response", err)
		return "", err
	}
	return jsonMap["access_token"].(string), nil
}

func (c *client) listSubscriptions(token string) ([]Subscription, error) {
	panic("implement me")
}