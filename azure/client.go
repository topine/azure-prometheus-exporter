package azure

import "net/http"

type Client interface {
	authenticate() (string, error)
	listSubscriptions(token string) ([]Subscription, error)
	getResources(subscription Subscription, token string) ([]Resource, error)
}

const (
	authentication = "https://login.microsoftonline.com/%s/oauth2/token"
)

type client struct {
	httpClient                       *http.Client
	clientId, clientSecret, tenantId string
}

func (c *client) getResources(subscription Subscription, token string) ([]Resource, error) {
	panic("implement me")
}

func (c *client) authenticate() (string, error) {
	panic("implement me")
}

func (c *client) listSubscriptions(token string) ([]Subscription, error) {
	panic("implement me")
}
