package main

import (
	"log"

	goform "github.com/erice5005/GoForm"
	"github.com/erice5005/GoForm/webhook"
)

func main() {
	path1Handler := webhook.NewFormHandler("/test1", func(wr goform.WebhookResponse) {
		// log.Printf("AnswerCount: %v\n", len(wr.FormResponse.Answers))
		// for _, ax := range wr.FormResponse.Answers {
		// 	log.Printf("Answer: %v\n", ax)
		// }
		ax := wr.FormResponse.MapAnswersByName()
		log.Printf("ax: %v\n", ax["Nice to meet you, {{field:01GCC86D4WQFNGC5N4SSJGG9Y2}}, how is your day going?"])
	})
	ws := webhook.NewWebhookService()
	ws.AddHandler(*path1Handler)
	ws.Serve(":8999")
}
