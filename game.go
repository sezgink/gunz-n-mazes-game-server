// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Game struct {
	// Registered clients.
	clients map[*Client]bool

	// Inbound messages from the clients.
	//broadcast chan []byte
	brodcast chan []byte

	//pull raw player data
	receive chan MCMessage

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client

	hub *Hub

	tickCounter int16

	playerCounter int16
}

func newGame(h *Hub) *Game {
	return &Game{
		brodcast:      make(chan []byte),
		receive:       make(chan MCMessage),
		register:      make(chan *Client),
		unregister:    make(chan *Client),
		clients:       make(map[*Client]bool),
		playerCounter: 0,
		hub:           h,
	}
}

func ticker(g *Game) {
	ticker := time.NewTicker(100 * time.Millisecond)
	for t := range ticker.C {
		fmt.Println("Tick at", t)
		g.brodcast <- newMessageUpdateGame(g)
	}

}

func (g *Game) runGame() {
	go ticker(g)
	for {

		select {
		case client := <-g.register:
			//fmt.Println("One registered")
			//g.hub.broadcast <- []byte("One registered")

			g.clients[client] = true
			client.player = &PlayerData{Id: g.playerCounter, PosX: 0, PosY: 0, Rot: 0, Vx: 0, Vy: 0}
			g.playerCounter += 1
			//Send client create playerMan message
			//client.send <-
			/*
				select {
				case client.send <- CreatePlrCreatorMessage(client.player):
				}
			*/
			//client.send <- CreatePlrCreatorMessage(client.player)
			//fmt.Println(CreatePlrCreatorMessage(client.player))

			//otherPlayers := make([]PlayerData, len(g.clients)-1)
			var otherPlayers []PlayerData

			for cli := range g.clients {
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
				//case client.send <- CreatePlrCreatorMessage(client.player):
			}
			//fmt.Println("Yeah registered")
			/*
				select {
				case client.send <- []byte("Welcome brother"):
				default:
					close(client.send)
					delete(g.clients, client)
				}
				for cli := range g.clients {
					if cli != client {
						cli.send <- []byte("We have a brother")
					}
				}
			*/

		case client := <-g.unregister:
			select {
			case g.hub.broadcast <- newMessageDestroyPlayer(client.player):
			}
			if _, ok := g.clients[client]; ok {
				delete(g.clients, client)
				//close(client.send)
			}

		case message := <-g.brodcast:
			//fmt.Println("Distrtubute")
			select {
			case g.hub.broadcast <- message:
			}

			//checkMessage(message)

			/*
				for client := range g.clients {
					select {
					case client.send <- []byte(message):
					default:
						close(client.send)
						delete(g.clients, client)
					}
				}
			*/
		case message := <-g.receive:
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
				g.hub.broadcast <- newMessageCreateFire(playerFire.PosX, playerFire.PosY, playerFire.Rot)
				fmt.Printf("%f %f %f", playerFire.PosX, playerFire.PosX, playerFire.Rot)
			} else {
				fmt.Println("Unknown message")
			}
		}
		//fmt.Println("End select")
		//g.tickCounter += 1
	}
}
