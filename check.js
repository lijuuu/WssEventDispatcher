const ws = new WebSocket("ws://localhost:7777/ws");

ws.onopen = () => {
  console.log("connected");
  ws.send(JSON.stringify({ type: "ping", payload: {} }));
};

ws.onmessage = (e) => console.log("received:", e.data);