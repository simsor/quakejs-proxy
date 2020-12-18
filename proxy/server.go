package proxy

import (
	"encoding/hex"
	"fmt"
	"net"
	"net/url"

	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

// SocketServer is the proxy server
type SocketServer struct {
	ListenAddress string
	ListenPort    int
	Destination   string

	listener net.PacketConn
	agents   map[string]*agent
}

// New creates a new socket to websocket proxy server listening on the given address, on port 27960
func New(listen, dest string) SocketServer {
	return SocketServer{
		ListenAddress: listen,
		ListenPort:    27960,
		Destination:   dest,
		agents:        make(map[string]*agent),
	}
}

func (s *SocketServer) listen() error {
	// Socket server
	listener, err := net.ListenPacket("udp4", fmt.Sprintf("%s:%d", s.ListenAddress, s.ListenPort))
	if err != nil {
		return err
	}

	s.listener = listener
	return nil
}

// Start is a blocking call which starts the proxy server
func (s *SocketServer) Start() error {
	err := s.listen()
	if err != nil {
		return err
	}

	p := make([]byte, 65535)

	for {
		n, addr, err := s.listener.ReadFrom(p)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"err": err,
			}).Error("Could not ReadFrom() a connection")
			continue
		}
		//logrus.WithField("remote", addr.String()).Info("Accepted UDP connection")

		a := s.getAgent(addr)
		err = a.ws.WriteMessage(websocket.BinaryMessage, p[:n])
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"err":    err,
				"remote": addr.String(),
			}).Error("Could not WriteMessage()")
			continue
		}
		logrus.WithField("remote", addr.String()).Info("Successfully sent UDP message to WebSocket")
		fmt.Println(hex.Dump(p[:n]))
	}

}

// Close closes both sockets
func (s *SocketServer) Close() error {
	return s.listener.Close()
}

func (s *SocketServer) getAgent(ip net.Addr) *agent {
	a, ok := s.agents[ip.String()]
	if ok {
		return a
	}

	// WebSocket client
	u := url.URL{Scheme: "ws", Host: s.Destination, Path: "/"}
	ws, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		panic(err)
	}

	a = &agent{
		ws:     ws,
		sock:   s.listener,
		remote: ip,
	}
	go a.ws2sock()
	s.agents[ip.String()] = a

	return a
}
