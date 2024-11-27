package webhooks

type WebhookConfig struct {
	Name string // A descriptive name for the webhook
	URL  string // The URL of the Discord webhook
}

var Webhooks = []WebhookConfig{
	{
		Name: "our",
		URL:  "https://discord.com/api/webhooks/1311378570978787399/frM_TNGHzmmehHia6QSBOon-uWbu2Lez5Vd5hN4RwWFq53VsWNYcvY1R8dR8QNE0HtIF",
	},
	{
		Name: "their",
		URL:  "https://discord.com/api/webhooks/1302674995280871545/fsmwXtFfChCn7ktcF3Gy8Pu0mv8YeOv9Izht3yC7Kstm5gHsa8ovmSvepksTpKXc7ICe",
	},
}
