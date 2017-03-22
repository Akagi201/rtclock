package main

import (
	"net/http"

	"github.com/Akagi201/light"
	log "github.com/Sirupsen/logrus"
	"golang.org/x/net/websocket"
)

func main() {
	root := light.New()

	root.Get("/clock", websocket.Handler(handleClock).ServeHTTP)
	root.Get("/", handleHome)

	log.Infof("HTTP listening at: %v", opts.ListenAddr)
	http.ListenAndServe(opts.ListenAddr, root)
}
