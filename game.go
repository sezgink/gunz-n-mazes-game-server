// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

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
	//g.distribute<-

}

func (g *Game) runGame() {
	for {
		if g.tickCounter > 1000 {
			g.tickCounter = 0
			OnTick(g)
		}
		select {
		case client := <-g.register:
			g.clients[client] = true
			client.player = &PlayerData{id: g.playerCounter, posX: 0, posY: 0, rot: 0, vx: 0, vy: 0, isCreated: false}
			client.send <- 
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
				close(client.send)
			}

		case message := <-g.distribute:
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
			g.hub.broadcast <- message
		}
		g.tickCounter += 1
	}
}
