// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import "fmt"

type Game struct {
	// Registered clients.
	clients map[*Client]bool

	// Inbound messages from the clients.
	//broadcast chan []byte
	distribute chan []byte

	//pull raw player data
	rawDataPuller chan MCMessage

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
		//broadcast:  make(chan []byte),
		distribute:    make(chan []byte),
		register:      make(chan *Client),
		rawDataPuller: make(chan MCMessage),
		unregister:    make(chan *Client),
		clients:       make(map[*Client]bool),
		playerCounter: 0,
		hub:           h,
	}
}
func OnTick(g *Game) {
	fmt.Println("Tick")
	g.distribute <- createUpdateMessage(g)

	/*
		Msg2Client := &Msg2Client{
			creatorMessage: new(CreatorMessage),
			updateMessage:  new(UpdateMessage),
		}
	*/
	//Msg2Client.creatorMessage.players2Create
	/*
		for cli := range g.clients {
			if !cli.player.isCreated {
				cli.send <- []byte("We have a brother")

			}
			cli.send <- []byte("We have a brother")
		}
	*/

}

func (g *Game) runGame() {
	fmt.Println("Game run")
	for {
		/*
			if g.tickCounter > 1000 {
				g.tickCounter = 0
				fmt.Println("Tick If")
				OnTick(g)
			}
		*/

		select {
		case client := <-g.register:
			//fmt.Println("One registered")
			//g.hub.broadcast <- []byte("One registered")

			g.clients[client] = true
			client.player = &PlayerData{id: g.playerCounter, posX: 0, posY: 0, rot: 0, vx: 0, vy: 0}
			//Send client create playerMan message
			//client.send <-
			client.send <- CreatePlrCreatorMessage(client.player)
			fmt.Println(CreatePlrCreatorMessage(client.player))
			for cli := range g.clients {
				if cli != client {
					cli.send <- CreateCreatorMessage(client.player)
				}
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
			if _, ok := g.clients[client]; ok {
				delete(g.clients, client)
				//close(client.send)
			}

		case message := <-g.distribute:
			fmt.Println("Distrtubute")
			g.hub.broadcast <- message
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

		}
		fmt.Println("End select")
		//g.tickCounter += 1
	}
}
