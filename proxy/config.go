package proxy

var hexdumpPackets = false
var logNewConnections = true
var logExchanges = false

// SetHexdumpPackets determines if the proxy should print a hexdump of every packet from both the client and server (default: false)
func SetHexdumpPackets(v bool) {
	hexdumpPackets = v
}

// SetLogNewConnections sets whether a message should be printed for every new connection to the proxy (default: true)
func SetLogNewConnections(v bool) {
	logNewConnections = v
}

// SetLogExchanges sets whether a log line should be printed every time a packet goes through the proxy (default: false)
func SetLogExchanges(v bool) {
	logExchanges = v
}
