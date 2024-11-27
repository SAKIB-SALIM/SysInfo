package main

import (
	"main.go/modules/system"
	"main.go/webhooks" // Replace "your_project" with the actual module name if different
)

func main() {
	for _, webhook := range webhooks.Webhooks {
		system.Run(webhook.URL)
	}
}
