// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"fmt"
	"time"
)

func checkMessage(message []byte) bool {
	str := fmt.Sprintf("%s", message)
	fmt.Println("Cheking the message...")
	msec := time.Now().UnixNano() / 1000000

	fmt.Println(str)

	fmt.Println(fmt.Sprintf("Time is %d", msec))

	var pData PlayerData
	/*
		if err := json.Unmarshal(message, &pData); err != nil {
			panic(err)
		}

		fmt.Println(fmt.Sprintf("Player id is %d", pData.id))
	*/
	if fmt.Sprintf("%c", message[0]) == "{" {
		if err := json.Unmarshal(message, &pData); err != nil {
			panic(err)
		}

		fmt.Println(fmt.Sprintf("Player id is %d", pData.Id))
	}

	if str == "deneme" {
		return true

	} else {
		return false
	}
}

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	clients map[*Client]bool

	// Inbound messages from the clients.
	broadcast chan []byte

	//Client puller for pull player data
	receive chan MCMessage

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client

	game *Game

	playerCounter int16

	tickCounter int16
}

func newHub() *Hub {
	return &Hub{
		broadcast:     make(chan []byte),
		receive:       make(chan MCMessage),
		register:      make(chan *Client),
		unregister:    make(chan *Client),
		clients:       make(map[*Client]bool),
		playerCounter: 0,
	}
}

func ticker(h *Hub) {
	ticker := time.NewTicker(100 * time.Millisecond)
	for t := range ticker.C {
		fmt.Println("Tick at", t)
		h.broadcast <- newMessageUpdateGame(h)
	}

}

func (h *Hub) run() {
	go ticker(h)
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true

			client.player = &PlayerData{Id: h.playerCounter, PosX: 0, PosY: 0, Rot: 0, Vx: 0, Vy: 0}
			h.playerCounter += 1

			var otherPlayers []PlayerData

			for cli := range h.clients {
				if cli != client {
					select {
					case cli.send <- newMessageCreatePlayer(client.player):
					}
					otherPlayers = append(otherPlayers, *cli.player)
					//cli.send <- CreateCreatorMessage(client.player)
					//fmt.Println("append")
					//fmt.Println(*client.player)
				}
			}
			fmt.Println(otherPlayers)
			select {
			case client.send <- newMessageCreateGame(client.player, otherPlayers):
			}

		case client := <-h.unregister:
			newMessageDestroyPlayer(client.player)
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}

			msg := newMessageDestroyPlayer(client.player)

			for client := range h.clients {
				//fmt.Println("Client broad")
				select {
				case client.send <- msg:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}

		case message := <-h.receive:
			var typedMessage Message
			json.Unmarshal(message.message, &typedMessage)
			if typedMessage.Mtype == PlAYER_STATE {
				//fmt.Println("\n PlAYER_STATE message")
				var playerState MessagePlayerState
				json.Unmarshal(message.message, &playerState)
				message.sender.player.PosX = playerState.PosX
				message.sender.player.PosY = playerState.PosY
				message.sender.player.Vx = playerState.Vx
				message.sender.player.Vy = playerState.Vy
				message.sender.player.Rot = playerState.Rot
				//fmt.Printf("%f %f %f %f %f", playerState.PosX, playerState.PosX, playerState.Rot, playerState.Vx, playerState.Vy)
			} else if typedMessage.Mtype == PLAYER_FIRE {
				fmt.Println("\n PLAYER_FIRE message")
				var playerFire MessagePlayerFire
				json.Unmarshal(message.message, &playerFire)
				msg := newMessageCreateFire(playerFire.PosX, playerFire.PosY, playerFire.Rot)

				for client := range h.clients {
					//fmt.Println("Client broad")
					select {
					case client.send <- msg:
					default:
						close(client.send)
						delete(h.clients, client)
					}
				}
				//fmt.Printf("%f %f %f", playerFire.PosX, playerFire.PosX, playerFire.Rot)
			} else {
				fmt.Println("Unknown message")
			}

		case message := <-h.broadcast:
			//fmt.Println("Broad message firsst")
			//fmt.Println(message)
			for client := range h.clients {
				//fmt.Println("Client broad")
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
			//checkMessage(message)

		}
	}
}
