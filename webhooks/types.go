// Package webhooks provides types/client for making requests to /webhooks.
package webhooks

// Webhook represents a postman webhook.
type Webhook struct {
	ID         string `json:"id"` // Output only.
	Name       string `json:"name"`
	Collection string `json:"collection"`
	WebhookURL string `json:"webhookUrl"` // Output only.
	UID        string `json:"uid"`        // Output only.
}

type webhookWrapper struct {
	Webhook Webhook `json:"webhook"`
}
