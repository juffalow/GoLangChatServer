//
// Chat server for chat.juffalow.com
//
//
// Build with : http://iris-go.com/
// See : https://github.com/iris-contrib/examples/blob/master/websocket/main.go


package main

import (
	"fmt"
	"time"
	"github.com/kataras/iris"
)

type Client struct {
    username string
    wsc iris.WebsocketConnection
}

func (client *Client) run() {
    client.wsc.Join("global")

    client.wsc.On("chat", func(message string) {
		t := time.Now()
        client.wsc.To("global").Emit("chat", "[ " + t.Format("3:04:05") + " ] " + client.username + ": " + message)
    })

	client.wsc.On("login", func(message string) {
		client.username = message
		fmt.Printf("\nNew user: ", message)
	})

    client.wsc.OnDisconnect(func() {
        fmt.Printf("\nConnection with ID: %v has been disconnected!", client.username)
    })
}

func NewClient(c iris.WebsocketConnection) *Client {
    client := &Client{wsc: c}

    client.run()

    return client
}

func main() {
	iris.Config.Websocket.Endpoint = "/chatapp"

	iris.Websocket.OnConnection(func(c iris.WebsocketConnection) {
        NewClient(c)
	})

	iris.Listen(":8080")
}
