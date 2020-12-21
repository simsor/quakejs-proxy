package proxy

import (
	"encoding/hex"
	"fmt"
	"net"

	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

type agent struct {
	sock *net.UDPConn
	ws   *websocket.Conn
	addr *net.UDPAddr

	running bool
	udpData chan []byte
}

func (a *agent) ws2sock() {
	log := logrus.WithFields(logrus.Fields{
		"remote": a.addr.String(),
		"func":   "ws2sock",
	})
	defer func() {
		a.ws.Close()
		a.running = false
	}()

	for {
		t, p, err := a.ws.ReadMessage()
		if err != nil {
			log.WithField("err", err).Error("Could not read WebSocket message")
			return
		}

		if logExchanges {
			log.Info("Got WS message")
		}

		if t != websocket.BinaryMessage {
			log.Error("Wrong message type")
			return
		}

		n, err := a.sock.WriteTo(p, a.addr)
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

		if logExchanges {
			log.Info("Successfully sent WS message to socket")
		}

		if hexdumpPackets {
			fmt.Println(hex.Dump(p))
		}
	}
}

func (a *agent) sock2ws() {
	log := logrus.WithFields(logrus.Fields{
		"remote": a.addr.String(),
		"func":   "sock2ws",
	})

	defer func() {
		a.ws.Close()
		a.running = false
	}()

	for p := range a.udpData {
		err := a.ws.WriteMessage(websocket.BinaryMessage, p)
		if err != nil {
			log.WithField("err", err).Error("Could not WriteMessage()")
			return
		}

		if logExchanges {
			log.Info("Successfully sent UDP message to WebSocket")
		}

		if hexdumpPackets {
			fmt.Println(hex.Dump(p))
		}
	}
}
