
# WebSocket Event Dispatcher (Go)

This is a minimal example demonstrating how to implement the **Dispatcher Pattern** for handling WebSocket events in Go.

Instead of writing multiple conditionals (`if`, `switch`) for handling each WebSocket event, this design cleanly routes incoming messages to registered handlers using a centralized dispatcher — similar to how HTTP routers work.

## 🧠 What’s Inside

- 📦 `dispatcher/dispatcher.go`: Core dispatcher implementation
- 🧪 `main.go`: Basic WebSocket server using the dispatcher
- 🧾 JS snippet: Sample client code to send events

## 💡 How It Works

Each incoming message is expected to have this structure:

```json
{
  "type": "join",
  "payload": {
    "userId": "abc123",
    "challengeId": "xyz789"
  }
}
````

You register handlers like this:

```go
dispatcher.Register("join", handlers.Join})
```

The dispatcher then automatically routes any `"type": "join"` message to the above handler.

