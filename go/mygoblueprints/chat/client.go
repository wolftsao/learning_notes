package main

import (
	"time"

	"github.com/gorilla/websocket"
)

type client struct {
	socket   *websocket.Conn
	send     chan *message
	room     *room
	userData map[string]any
}

func (c *client) read() {
	defer c.socket.Close()
	for {
		var msg *message
		if err := c.socket.ReadJSON(&msg); err != nil {
			return
		}
		msg.When = time.Now()
		msg.Name = c.userData["name"].(string)
		// msg.AvatarURL, _ = c.room.avatar.GetAvatarURL(c)
		if avatarUrl, ok := c.userData["avatar_url"]; ok {
			msg.AvatarURL = avatarUrl.(string)
		}

		// if avatarURL, ok := c.userData["avatar_url"]; ok {
		// 	msg.AvatarURL = avatarURL.(string)
		// }

		c.room.forward <- msg
	}
}

func (c *client) write() {
	defer c.socket.Close()
	for msg := range c.send {
		err := c.socket.WriteJSON(msg)
		if err != nil {
			break
		}
	}
}
