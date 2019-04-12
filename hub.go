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

		fmt.Println(fmt.Sprintf("Player id is %d", pData.id))
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
	//broadcast chan []byte
	broadcast chan MCMessage

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client
}

func newHub() *Hub {
	return &Hub{
		//broadcast:  make(chan []byte),
		broadcast:  make(chan MCMessage),
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
			//fmt.Println("Yeah registered")
			select {
			case client.send <- []byte("Welcome brother"):
			default:
				close(client.send)
				delete(h.clients, client)
			}
			for cli := range h.clients {
				if cli != client {
					cli.send <- []byte("We have a brother")
				}
			}

		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}

		case message := <-h.broadcast:
			//checkMessage(message)

			for client := range h.clients {
				select {
				case client.send <- []byte(message.message):
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}
