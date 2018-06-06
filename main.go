package main

import (
	"log"
	"net/http"
	"os"

	"github.com/atla/lotd/game"
	"github.com/atla/lotd/tcp"
	"github.com/atla/lotd/ws"
)

const motd = `
██▓     ▄▄▄       ██▓ ██▀███      ▒█████    █████▒   ▄▄▄█████▓ ██░ ██ ▓█████ 
▓██▒    ▒████▄    ▓██▒▓██ ▒ ██▒   ▒██▒  ██▒▓██   ▒    ▓  ██▒ ▓▒▓██░ ██▒▓█   ▀ 
▒██░    ▒██  ▀█▄  ▒██▒▓██ ░▄█ ▒   ▒██░  ██▒▒████ ░    ▒ ▓██░ ▒░▒██▀▀██░▒███   
▒██░    ░██▄▄▄▄██ ░██░▒██▀▀█▄     ▒██   ██░░▓█▒  ░    ░ ▓██▓ ░ ░▓█ ░██ ▒▓█  ▄ 
░██████▒ ▓█   ▓██▒░██░░██▓ ▒██▒   ░ ████▓▒░░▒█░         ▒██▒ ░ ░▓█▒░██▓░▒████▒
░ ▒░▓  ░ ▒▒   ▓▒█░░▓  ░ ▒▓ ░▒▓░   ░ ▒░▒░▒░  ▒ ░         ▒ ░░    ▒ ░░▒░▒░░ ▒░ ░
░ ░ ▒  ░  ▒   ▒▒ ░ ▒ ░  ░▒ ░ ▒░     ░ ▒ ▒░  ░             ░     ▒ ░▒░ ░ ░ ░  ░
 ░ ░     ░   ▒    ▒ ░  ░░   ░    ░ ░ ░ ▒   ░ ░         ░       ░  ░░ ░   ░   
   ░  ░      ░  ░ ░     ░            ░ ░                       ░  ░  ░   ░  ░
																				
	▓█████▄  ██▀███   ▄▄▄        ▄████  ▒█████   ███▄    █                        
	▒██▀ ██▌▓██ ▒ ██▒▒████▄     ██▒ ▀█▒▒██▒  ██▒ ██ ▀█   █                        
	░██   █▌▓██ ░▄█ ▒▒██  ▀█▄  ▒██░▄▄▄░▒██░  ██▒▓██  ▀█ ██▒                       
	░▓█▄   ▌▒██▀▀█▄  ░██▄▄▄▄██ ░▓█  ██▓▒██   ██░▓██▒  ▐▌██▒                       
	░▒████▓ ░██▓ ▒██▒ ▓█   ▓██▒░▒▓███▀▒░ ████▓▒░▒██░   ▓██░                       
	▒▒▓  ▒ ░ ▒▓ ░▒▓░ ▒▒   ▓▒█░ ░▒   ▒ ░ ▒░▒░▒░ ░ ▒░   ▒ ▒                        
	░ ▒  ▒   ░▒ ░ ▒░  ▒   ▒▒ ░  ░   ░   ░ ▒ ▒░ ░ ░░   ░ ▒░                       
	░ ░  ░   ░░   ░   ░   ▒   ░ ░   ░ ░ ░ ░ ▒     ░   ░ ░                        
	░       ░           ░  ░      ░     ░ ░           ░                        
	░                                                                            
`

func main() {

	logFile, err := os.OpenFile("log.txt", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	log.SetOutput(logFile)

	game := game.GetInstance()
	game.MOTD = motd

	// Create a simple file server
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)

	tcpServer := tcp.NewServer("8023")
	go tcpServer.Start()

	webSocketServer := ws.NewWebSocketServer("8080")
	webSocketServer.Start()
}
