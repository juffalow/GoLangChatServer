//
// Chat server for chat.juffalow.com
//
//
// Build with : http://iris-go.com/
// See : https://github.com/iris-contrib/examples/blob/master/websocket/main.go


package main

import (
	"fmt"
	"github.com/kataras/iris"
)

type clientPage struct {
	Title string
	Host  string
}

type Client struct {
    username string
    wsc iris.WebsocketConnection
}

func (client *Client) run() {
    client.wsc.Join("room1")

    client.wsc.On("chat", func(message string) {
        client.wsc.To("room1").Emit("chat", "From: " + client.username + ": " + message)
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
	iris.Static("/js", "./static/js", 1)

	iris.Get("/", func(ctx *iris.Context) {
		ctx.Render("client.html", clientPage{"Client Page", ctx.HostString()})
	})

	iris.Config.Websocket.Endpoint = "/my_endpoint"

	iris.Websocket.OnConnection(func(c iris.WebsocketConnection) {
        NewClient(c)
	})

	iris.Listen(":8080")
}
