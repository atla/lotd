package main

import (
	"context"
	"net/http"

	"github.com/atla/lotd/dba"
	"github.com/atla/lotd/game"
	"github.com/atla/lotd/motd"
	"github.com/atla/lotd/tcp"
	"github.com/atla/lotd/ws"
	"go.uber.org/fx"

	log "github.com/sirupsen/logrus"
)

func startFrontend() {
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)
}

func initServers(lifecycle fx.Lifecycle, tcpServer *tcp.Server, webSocketServer *ws.Server) {
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go tcpServer.Start()
			go webSocketServer.Start()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			tcpServer.Stop()
			webSocketServer.Stop()
			return nil
		},
	})
}

func main() {

	log.Info("Starting server")

	app := fx.New(
		dba.AccessModule,
		dba.FacadeModule,
		motd.Module,
		game.Module,
		tcp.Module,
		ws.Module,
		fx.Invoke(startFrontend,
			initServers),
	)
	app.Run()

	/*	log.WithFields(log.Fields{
			"animal": "walrus",
		}).Info("A walrus appears")
	*/
}
