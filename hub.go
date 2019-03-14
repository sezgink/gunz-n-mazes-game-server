// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"time"
)

func checkMessage(message []byte) bool {
	str := fmt.Sprintf("%s", message)
	fmt.Println("Cheking the message...")
	msec := time.Now().UnixNano() / 1000000
	fmt.Println(fmt.Sprintf("Time is %d", msec))
	fmt.Println(str)
	if str == "deneme" {
		return true

	} else {
		return false
	}
}

type PlayerData struct {
	posX float32
	posY float32
	vx   float32
	vy   float32
}

func parsePlayerData(message []byte) bool {
	str := fmt.Sprintf("%s", message)
	fmt.Println("Cheking the message...")
	fmt.Println(str)
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

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client
}

func newHub() *Hub {
	return &Hub{
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:
			checkMessage(message)
			for client := range h.clients {
				select {
				case client.send <- []byte(message):
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}
