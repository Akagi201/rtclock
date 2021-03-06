package main

import (
	"html/template"
	"io"
	"net/http"
	"time"

	log "github.com/Sirupsen/logrus"
	"golang.org/x/net/websocket"
)

func handleHome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	t := template.Must(template.New("base").Parse(string(MustAsset(opts.Template))))
	v := struct {
		Host string
	}{
		r.Host,
	}
	if err := t.Execute(w, &v); err != nil {
		log.Errorf("Template execute failed, err: %v", err)
		return
	}
}

func handleClock(ws *websocket.Conn) {
	interval, err := time.ParseDuration(opts.Interval)
	if err != nil {
		log.Errorf("Parse interval failed, err: %v", err)
		return
	}
	for {
		now := time.Now()
		clock := now.Format(time.RFC3339Nano)
		io.WriteString(ws, clock)
		time.Sleep(interval)
	}
}
