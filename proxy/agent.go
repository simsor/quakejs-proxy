package proxy

import (
	"encoding/hex"
	"fmt"
	"net"

	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

type agent struct {
	remote net.Addr
	sock   net.PacketConn
	ws     *websocket.Conn
}

func (a *agent) ws2sock() {
	log := logrus.WithFields(logrus.Fields{
		"remote": a.remote.String(),
		"func":   "ws2sock",
	})
	defer a.ws.Close()

	for {
		t, p, err := a.ws.ReadMessage()
		if err != nil {
			log.WithField("err", err).Error("Could not read WebSocket message")
			return
		}
		log.Info("Got WS message")

		if t != websocket.BinaryMessage {
			log.Error("Wrong message type")
			return
		}

		n, err := a.sock.WriteTo(p, a.remote)
		if err != nil {
			log.WithError(err).Error("Could not write to socket")
			return
		}

		if n != len(p) {
			log.WithFields(logrus.Fields{
				"written": n,
				"length":  len(p),
			}).Error("Did not write full message")
			return
		}
		log.Info("Successfully sent WS message to socket")

		fmt.Println(hex.Dump(p))
	}
}
