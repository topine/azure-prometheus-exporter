package azure

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

var httpTest = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	if r.URL.EscapedPath() == fmt.Sprintf(authentication, "tenant-id") {
		b, err := ioutil.ReadFile("testdata/authenticationResponse.json")
		if err != nil {
			log.Panic("Error reading testdata/authentication.json file.", err)
		}

		fmt.Fprint(w, string(b))
	}
}))

func Test_client_authenticate(t *testing.T) {
	type fields struct {
		httpClient   *http.Client
		clientId     string
		clientSecret string
		tenantId     string
		loginBaseURL string
		azureBaseURL string
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{name: "authentication",
			fields: fields{
				httpClient:   httpTest.Client(),
				clientId:     "client-id",
				clientSecret: "client-secret",
				tenantId:     "tenant-id",
				loginBaseURL: httpTest.URL,
			}, want: "token-response",
			wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &client{
				httpClient:   tt.fields.httpClient,
				clientId:     tt.fields.clientId,
				clientSecret: tt.fields.clientSecret,
				tenantId:     tt.fields.tenantId,
				loginBaseURL: tt.fields.loginBaseURL,
			}
			got, err := c.authenticate()
			if (err != nil) != tt.wantErr {
				t.Errorf("authenticate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("authenticate() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_client_getResources(t *testing.T) {
	type fields struct {
		httpClient   *http.Client
		clientId     string
		clientSecret string
		tenantId     string
	}
	type args struct {
		subscription Subscription
		token        string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []Resource
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &client{
				httpClient:   tt.fields.httpClient,
				clientId:     tt.fields.clientId,
				clientSecret: tt.fields.clientSecret,
				tenantId:     tt.fields.tenantId,
			}
			got, err := c.getResources(tt.args.subscription, tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("getResources() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getResources() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_client_listSubscriptions(t *testing.T) {
	type fields struct {
		httpClient   *http.Client
		clientId     string
		clientSecret string
		tenantId     string
		loginBaseURL string
		azureBaseURL string
	}
	type args struct {
		token string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []Subscription
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &client{
				httpClient:   tt.fields.httpClient,
				clientId:     tt.fields.clientId,
				clientSecret: tt.fields.clientSecret,
				tenantId:     tt.fields.tenantId,
			}
			got, err := c.listSubscriptions(tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("listSubscriptions() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("listSubscriptions() got = %v, want %v", got, tt.want)
			}
		})
	}
}
