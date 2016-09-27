package main

//
//
//

import (
	"fmt"
	"time"
	"github.com/kataras/iris"
	"github.com/kataras/go-websocket"
	"encoding/json"
    "strings"
)

type Client struct {
    username string
    wsc iris.WebsocketConnection
	chatRoom *ChatRoom
	index int
}

func (client *Client) run() {
    client.wsc.Join(client.chatRoom.name)

    client.wsc.On("chat", func(message string) {
		t := time.Now()

		m := Message{t.Format(time.RFC3339), client.username, strings.TrimSpace(message)}

		client.chatRoom.AddMessage(m)

		b, _ := json.Marshal(m)

        client.wsc.To(websocket.All).Emit("chat", string(b))
    })

	client.wsc.On("login", func(username string) {
		if (len(username) > 0) && (len(username) <= 16) {
			t := time.Now()
			fmt.Printf("[ %s ] New user: %s\n", t.Format(time.RFC3339),  username)

	        client.username = username

			initialInformations := InitialInformation{client.chatRoom.GetUsernames(), client.chatRoom.GetLastMessages()}

			b, _ := json.Marshal(initialInformations)
			client.wsc.Emit("init", string(b))

			client.wsc.To(websocket.NotMe).Emit("user:connect", username)

			return
		}
		client.wsc.To(client.chatRoom.name).Emit("login_error", "Username is too short or too long!")
	})

    client.wsc.OnDisconnect(func() {
		client.chatRoom.Disconnected(client.username, client.wsc.ID())

		client.wsc.To(client.chatRoom.name).Emit("user:disconnect", client.username)
    })
}
