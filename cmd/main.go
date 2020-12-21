package main

import (
	"flag"

	"github.com/simsor/quakejs-proxy/proxy"
	"github.com/sirupsen/logrus"
)

var (
	wsServer = flag.String("ws", "127.0.0.1:27961", "Hostname of the WebSocket server")
	listen   = flag.String("listen", "", "Host to listen on")

	hexdump    = flag.Bool("hexdump", false, "Print a hex dump of every packet")
	logNewConn = flag.Bool("log-new-conn", true, "Logs every new connection")
	logExch    = flag.Bool("log-exchanges", false, "Logs all exchanges going through the proxy")
)

func main() {
	flag.Parse()

	proxy.SetHexdumpPackets(*hexdump)
	proxy.SetLogExchanges(*logExch)
	proxy.SetLogNewConnections(*logNewConn)

	server := proxy.New(*listen, *wsServer)
	defer server.Close()

	logrus.WithFields(logrus.Fields{
		"listen": *listen,
		"dest":   *wsServer,
	}).Info("Proxy server listening")

	err := server.Start()
	if err != nil {
		panic(err)
	}
}
