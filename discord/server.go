package discord

import (
	"log"
	"net/http"
	"net/url"
)

type Server struct {
	Client  http.Client
	Webhook url.URL
}

func (d Server) SendMessage(msg string) {
	resp, err := http.DefaultClient.PostForm(
		d.Webhook.
			ResolveReference(&url.URL{RawQuery: "wait=1"}).
			String(),
		url.Values{"content": {msg}},
	)
	log.Printf("Posted '%s' to discord, result: %v", msg, resp)
	err = resp.Body.Close()
	if err != nil {
		panic(err)
	}
}
