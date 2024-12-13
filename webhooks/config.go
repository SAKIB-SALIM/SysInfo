package webhooks

// WebhookConfig represents the configuration of a Discord webhook.
type WebhookConfig struct {
	Name string // A descriptive name for the webhook.
	URL  string // The URL of the Discord webhook.
}

// Webhooks is a list of configured Discord webhooks.
var Webhooks = []WebhookConfig{
	{
		Name: "Something0", // Name identifying the first webhook.
		URL:  "", // Add your webhook here
	},
	{
		Name: "Something1", // Name identifying the second webhook.
		URL:  "", // Add your webhook here
	},
}
