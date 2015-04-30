package main

import (
	"golang.org/x/net/websocket"
	"io"
	"net/http"
	"os"
)

// Echo the data received on the WebSocket.
func LogServer(ws *websocket.Conn) {
	io.Copy(os.Stdout, ws)
}

// This example demonstrates a trivial echo server.
func main() {
	http.Handle("/log", websocket.Handler(LogServer))
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
