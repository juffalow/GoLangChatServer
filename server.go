//
// Chat server for chat.juffalow.com
//
//
// Build with : http://iris-go.com/
// See : https://github.com/iris-contrib/examples/blob/master/websocket/main.go


package main

import (
	"github.com/kataras/iris"
)

type ChatRoom struct {
	name string
	clients []*Client
	messages []Message
}

func (chatRoom *ChatRoom) Join(client *Client) {
	chatRoom.clients = append(chatRoom.clients, client)
}

// Returns slice of usernames that are currently connected
func (chatRoom *ChatRoom) GetUsernames() []string {
	var usernames []string

	for _, value := range chatRoom.clients {
		if( len(value.username) > 0 ) {
			usernames = append(usernames, value.username)
		}
	}
	return usernames
}

// Removes user from list of logged in users
func (chatRoom *ChatRoom) Disconnected(username string, id string) {
	for index, value := range chatRoom.clients {
		if (value.username == username) && (value.wsc.ID() == id) {
			chatRoom.clients = append(chatRoom.clients[:index], chatRoom.clients[index+1:]...)
		}
	}
}

// Add new message in message "queue".
// If there are more than 10 messages, it removes the oldest one.
func (chatRoom *ChatRoom) AddMessage(message Message) {
	chatRoom.messages = append(chatRoom.messages, message)
	if( len(chatRoom.messages) > 10 ) {
		chatRoom.messages = append(chatRoom.messages[:0], chatRoom.messages[1:]...)
	}
}

// Returns last N messages
func (chatRoom *ChatRoom) GetLastMessages() []Message {
	return chatRoom.messages
}

func NewClient(c iris.WebsocketConnection, chatRoom *ChatRoom) *Client {
    client := &Client{wsc: c, chatRoom: chatRoom}

    client.run()

    return client
}

func main() {
	chatRoom := &ChatRoom{name: "global"}

	iris.Config.Websocket.Endpoint = "/chatapp"

	iris.Websocket.OnConnection(func(c iris.WebsocketConnection) {
        chatRoom.Join(NewClient(c, chatRoom))
	})

	iris.Listen(":8080")
}
