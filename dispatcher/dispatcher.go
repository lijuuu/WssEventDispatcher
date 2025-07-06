// Package dispatcher provides a WebSocket event dispatcher pattern
// allowing modular, extensible, and testable WebSocket event handling.
package dispatcher

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/websocket"
)

// WssDispatcher manages routing of WebSocket messages
// to their corresponding handler functions based on the `type` field.
type WssDispatcher struct {
	handlers map[string]WssHandler
}

// WssContext wraps the request-specific data passed to each handler,
// including the WebSocket connection and the raw message payload.
// It also supports future use of context.Context for deadlines, tracing, etc.
type WssContext struct {
	Ctx     context.Context
	Conn    *websocket.Conn
	Payload json.RawMessage
}

// WssMessage defines the expected message structure sent over WebSocket.
//
//	Example:
//	{
//	  "type": "join",
//	  "payload": { "challengeId": "abc123" }
//	}
type WssMessage struct {
	Type    string          `json:"type"`    // event name, used for dispatching
	Payload json.RawMessage `json:"payload"` // raw payload forwarded to the handler
}

// WssHandler is the function signature for all WebSocket event handlers.
// The dispatcher passes in a WssContext that contains the connection and payload.
type WssHandler func(ctx *WssContext) error

// Upgrader upgrades an HTTP connection to a WebSocket connection.
// Note: In production, tighten the CheckOrigin logic.
var Upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all origins for development/testing
	},
}

// NewWssDispatcher initializes and returns a new event dispatcher.
func NewWssDispatcher() *WssDispatcher {
	return &WssDispatcher{
		handlers: make(map[string]WssHandler),
	}
}

// Register binds a given event type to a WssHandler.
// Only one handler per event type is supported; registering again will overwrite.
func (d *WssDispatcher) Register(event string, handler WssHandler) {
	d.handlers[event] = handler
}

// Dispatch decodes the incoming message, determines the event type,
// and routes it to the appropriate handler.
// Returns an error if the message is invalid or the handler is not registered.
func (d *WssDispatcher) Dispatch(conn *websocket.Conn, message []byte) error {
	var msg WssMessage
	if err := json.Unmarshal(message, &msg); err != nil {
		return err
	}

	handler, ok := d.handlers[msg.Type]
	if !ok {
		return errors.New("unrecognized event type: " + msg.Type)
	}

	ctx := &WssContext{
		Ctx:     context.Background(), // can be extended with timeout/deadline info
		Conn:    conn,
		Payload: msg.Payload,
	}

	return handler(ctx)
}
