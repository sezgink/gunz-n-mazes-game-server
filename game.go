// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import "fmt"
import "time"

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

func ticker(g *Game) {
	ticker := time.NewTicker(10000 * time.Millisecond)
	for t := range ticker.C {
		fmt.Println("Tick at", t)
		g.distribute <- newMessageUpdateGame(g)
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

		case message := <-g.distribute:
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
		case message := <-g.rawDataPuller:
			fmt.Println(message)
		}
		//fmt.Println("End select")
		//g.tickCounter += 1
	}
}
