package stripe

import (
	"encoding/json"
	"net/http"

	"github.com/stripe/stripe-cli/pkg/config"
	"github.com/stripe/stripe-cli/pkg/requests"
)

// WebhookEndpointList contains the list of webhook endpoints for the account
type WebhookEndpointList struct {
	Data []WebhookEndpoint `json:"data"`
}

// WebhookEndpoint contains the data for each webhook endpoint
type WebhookEndpoint struct {
	Application   string   `json:"application"`
	EnabledEvents []string `json:"enabled_events"`
	URL           string   `json:"url"`
}

// WebhookEndpointsList returns all the webhook endpoints on a users' account
func WebhookEndpointsList(baseURL, apiVersion, apiKey string, profile *config.Profile) WebhookEndpointList {
	params := &requests.RequestParameters{}
	params.AppendData([]string{"limit=30"})
	params.SetVersion(apiVersion)

	base := &requests.Base{
		Profile:        profile,
		Method:         http.MethodGet,
		SuppressOutput: true,
		APIBaseURL:     baseURL,
	}
	resp, _ := base.MakeRequest(apiKey, "/v1/webhook_endpoints", params, true)
	data := WebhookEndpointList{}
	json.Unmarshal(resp, &data)

	return data
}
