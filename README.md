# quakejs-proxy -- play on QuakeJS servers with a native ioquake3 client

## Description

QuakeJS-Proxy is a Golang Quake 3 Arena proxy server which relays UDP packets from a [ioquake3](https://ioquake3.org) client to a [QuakeJS](https://github.com/inolen/quakejs) WebSocket server. It allows you to play on QuakeJS servers using a native ioquake3 client.

This means using your custom keybindings, custom config, etc.

## Building

Clone the repository and run:

```shell
$ go build cmd/main.go
```

This project was successfully tested with Go 1.15 and should work with all later versions.

## Running

The only required parameter is:
- `-ws`: the URI to the QuakeJS server, i.e. `quakejs.com:27960`

Optional parameters:
- `-listen`: listen on a specific IP address. By default, it will listen on all available interfaces
- `-log-exchanges`: useful when debugging, prints a log line every time a packet is exchanged through the proxy
- `-hexdump`: useful when debugging, prints a hexdump of every packet going through the proxy
- `-log-new-conn=false`: disable logging every new connection

## Troubleshooting

**Connecting to the proxy server gives me "Game mismatch: this is a Quake 3 Arena server"**

Make sure you are using an up-to-date build of ioquake3. The Windows build available on their website is old and doesn't support the new protocol used by QuakeJS. Instructions for building ioquake3 on a modern Windows are available in [BUILDING-IOQUAKE3.md](BUILDING-IOQUAKE3.md).

**Connecting to the proxy server crashes the game on "Downloading gamestate"**

Make sure your PK3 files match up with the server's. Most QuakeJS servers use Quake 3's demo files, and attempting to join with the full version will crash your game.
See [GET-DEMO-PK.md](GET-DEMO-PK3.md) for instructions on how to download the demo PK3 files from a QuakeJS server.