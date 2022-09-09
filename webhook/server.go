package webhook

import (
	"log"

	"github.com/erice5005/autowebhook"
)

type WebhookService struct {
	ws *autowebhook.WebhookService
}

func NewWebhookService() *WebhookService {
	wx := autowebhook.NewWebhookService(nil)
	return &WebhookService{
		ws: &wx,
	}
}

func (ws WebhookService) AddHandler(fx FormHandler) {
	ws.ws.Handlers = append(ws.ws.Handlers, fx.hf)
}

func (ws WebhookService) Serve(addr string) error {
	ws.ws.Init()
	log.Printf("Handlers: %v\n", ws.ws.Handlers)
	return ws.ws.Serve(addr)
}
