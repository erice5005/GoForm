package webhook

import (
	goform "github.com/erice5005/GoForm"
	"github.com/erice5005/autowebhook"
)

type FormHandler struct {
	hf      *autowebhook.HandlerFrame
	Handler func(goform.WebhookResponse)
}

func NewFormHandler(path string, handler func(goform.WebhookResponse)) *FormHandler {
	return &FormHandler{
		hf: autowebhook.NewHandlerFram(path, func(i interface{}) {
			handler(i.(goform.WebhookResponse))
		}, goform.ParseWebhookResponse),
		Handler: handler,
	}
}

// func Run(addr string) {}
