package handlers

import (
	"log"

	"github.com/gorilla/websocket"
	"github.com/lijuuu/WssEventDispatcher/dispatcher"
)

func Ping(ctx *dispatcher.WssContext) error {
	log.Println("ping received")
	return ctx.Conn.WriteMessage(websocket.TextMessage, []byte(`{"type":"pong"}`))
}
