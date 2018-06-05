package main

import (
	"lotd/game"
	"lotd/tcp"
	"lotd/ws"
	"net/http"
)

func main() {
	// Create a simple file server
	fs := http.FileServer(http.Dir("/public"))
	http.Handle("/", fs)

	tcpServer := tcp.NewServer(game.GetInstance())
	go tcpServer.Start()

	webSocketServer := ws.NewWebSocketServer("8000")
	webSocketServer.Start()
}
