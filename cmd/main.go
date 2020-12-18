package main

import (
	"flag"

	"github.com/simsor/quakejs-proxy/proxy"
	"github.com/sirupsen/logrus"
)

var (
	wsServer = flag.String("ws", "127.0.0.1:27961", "Hostname of the WebSocket server")
	listen   = flag.String("listen", "", "Host to listen on")
)

func main() {
	flag.Parse()

	server := proxy.New(*listen, *wsServer)

	logrus.WithFields(logrus.Fields{
		"listen": *listen,
		"dest":   *wsServer,
	}).Info("Proxy server listening")

	err := server.Start()
	if err != nil {
		panic(err)
	}
	server.Close()
}
