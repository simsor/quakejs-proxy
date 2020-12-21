package proxy

import (
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

	listener *net.UDPConn
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
	addr := net.UDPAddr{
		Port: s.ListenPort,
		IP:   net.ParseIP(s.ListenAddress),
	}
	listener, err := net.ListenUDP("udp4", &addr)
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
		n, addr, err := s.listener.ReadFromUDP(p)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"err": err,
			}).Error("Could not ReadFromUDP() a connection")
			continue
		}

		data := make([]byte, n)
		copy(data, p[:n])

		a := s.getAgent(addr)
		a.udpData <- data
	}

}

// Close closes both sockets
func (s *SocketServer) Close() error {
	return s.listener.Close()
}

func (s *SocketServer) getAgent(addr *net.UDPAddr) *agent {
	a, ok := s.agents[addr.String()]
	if ok && a.running {
		return a
	}

	// WebSocket client
	u := url.URL{Scheme: "ws", Host: s.Destination, Path: "/"}
	ws, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		panic(err)
	}

	a = &agent{
		ws:      ws,
		sock:    s.listener,
		addr:    addr,
		running: true,
	}
	a.udpData = make(chan []byte, 10000)
	go a.ws2sock()
	go a.sock2ws()
	s.agents[addr.String()] = a

	if logNewConnections {
		logrus.WithField("remote", addr.String()).Info("New client")
	}
	return a
}
