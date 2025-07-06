package handlers

import (
	"log"

	"github.com/gorilla/websocket"
	"github.com/lijuuu/WssEventDispatcher/dispatcher"
)

func Fizz(ctx *dispatcher.WssContext) error {
	log.Println("buzz received")
	return ctx.Conn.WriteMessage(websocket.TextMessage, []byte(`{"type":"buzz"}`))
}
