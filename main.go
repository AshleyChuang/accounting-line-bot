package main

import (
	"log"
	"net/http"

	"github.com/line/line-bot-sdk-go/linebot"
)

func main() {
	bot, err := linebot.New("7e718ef87940965b417a74d60c6d1a55", "gi+JkIOmr+gXanE4+t/YuEV2FDjt9JxcQRPoPjOMXy/RntS4j6HoPD7tBYbXJvqAvQiFzx1yW78+6TEmgpw7JuQ1Lp03jSx2XTebg8CMBcmPobShwB0XecmoLAB8nRzPsD5pM9n/x3EN6zvqsxDg4QdB04t89/1O/w1cDnyilFU=")
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/callback", func(w http.ResponseWriter, req *http.Request) {
		events, err := bot.ParseRequest(req)
		if err != nil {
			if err == linebot.ErrInvalidSignature {
				w.WriteHeader(400)
			} else {
				w.WriteHeader(500)
			}
			return
		}
		for _, event := range events {
			if event.Type == linebot.EventTypeMessage {
				switch message := event.Message.(type) {
				case *linebot.TextMessage:
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text)).Do(); err != nil {
						log.Print(err)
					}
				}
			}
		}
	})
	// This is just sample code.
	// For actual use, you must support HTTPS by using `ListenAndServeTLS`, a reverse proxy or something else.
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
