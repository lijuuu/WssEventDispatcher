package main

import (
	"log"
	"net/http"

	"github.com/lijuuu/WssEventDispatcher/dispatcher"
	"github.com/lijuuu/WssEventDispatcher/handlers"
)

func main() {
	// create a new WebSocket event dispatcher
	d := dispatcher.NewWssDispatcher()

	// register a handler for the "ping" event
	// responds with {"type": "pong"}
	d.Register("ping", handlers.Ping)
	// register a handler for the "ping" event
	// responds with {"type": "pong"}
	d.Register("fizz", handlers.Fizz)

	// define the WebSocket HTTP endpoint
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		// upgrade the incoming HTTP request to a WebSocket connection
		conn, err := dispatcher.Upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println("upgrade error:", err)
			return
		}
		defer conn.Close()

		log.Println("client connected")

		// continuously read messages from client
		for {
			_, msg, err := conn.ReadMessage()
			if err != nil {
				log.Println("read error:", err)
				break
			}

			// dispatch the message to the appropriate handler
			if err := d.Dispatch(conn, msg); err != nil {
				log.Println("dispatch error:", err)
				break
			}
		}
	})

	// start the WebSocket server
	log.Println("WebSocket server running on ws://localhost:7777/ws")
	if err := http.ListenAndServe(":7777", nil); err != nil {
		log.Fatal("server failed:", err)
	}
}
