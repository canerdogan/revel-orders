package controllers

import (
	"github.com/canerdogan/revel-orders/Godeps/_workspace/src/github.com/revel/revel"
	"github.com/canerdogan/revel-orders/Godeps/_workspace/src/golang.org/x/net/websocket"
	"github.com/canerdogan/revel-orders/app/chatroom"
	"github.com/canerdogan/revel-orders/app/models"
)

type WebSocket struct {
	GorpController
	*revel.Controller
}

func (c WebSocket) RoomSocket(ws *websocket.Conn) revel.Result {
	// Join the room.
	subscription := chatroom.Subscribe()
	defer subscription.Cancel()

	// In order to select between websocket messages and subscription events, we
	// need to stuff websocket events into a channel.
	newMessages := make(chan models.Requests)
	go func() {
		var msg models.Requests
		for {
			err := websocket.Message.Receive(ws, &msg)
			if err != nil {
				close(newMessages)
				return
			}
			newMessages <- msg
		}
	}()

	// Now listen for new events from either the websocket or the chatroom.
	for {
		select {
		case event := <-subscription.New:
			if websocket.JSON.Send(ws, &event) != nil {
				// They disconnected.
				return nil
			}
		case msg, ok := <-newMessages:
			// If the channel is closed, they disconnected.
			if !ok {
				return nil
			}
			// Otherwise, say something.
			chatroom.SendRequest(&msg)
		}
	}
	return nil
}
