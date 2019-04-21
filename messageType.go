package main

import (
	"encoding/json"
	"fmt"
)

type MessageType int

const (
	UPDATE_GAME    MessageType = 0
	CREATE_FIRE    MessageType = 1
	CREATE_PLAYER  MessageType = 2
	DESTROY_PLAYER MessageType = 3
	CREATE_GAME    MessageType = 4
	PlAYER_STATE   MessageType = 5
	PLAYER_FIRE    MessageType = 6
)

type Message struct {
	Mtype   MessageType `json:"Mtype"`
	message json.RawMessage
}

type MessageUpdateGame struct {
	Players []PlayerData
	Mtype   MessageType
}

type MessageFire struct {
	Mtype MessageType
}

type MessageCreatePlayer struct {
	IsOwner bool
	Player  PlayerData
	Mtype   MessageType
}

type MessageDestroyPlayer struct {
	Player PlayerData
	Mtype  MessageType
}

type MessageCreateGame struct {
	Player       PlayerData
	OtherPlayers []PlayerData
	Fires        []FireData
	Mtype        MessageType
}

type MessagePlayerState struct {
	PosX  float32
	PosY  float32
	Vx    float32
	Vy    float32
	Rot   float32
	Mtype MessageType
}

type MessagePlayerFire struct {
	PosX  float32
	PosY  float32
	Rot   float32
	Mtype MessageType
}

func newMessageCreatePlayer(pData *PlayerData) []byte {
	cm := new(MessageCreatePlayer)
	cm.Player = *pData
	cm.IsOwner = false
	cm.Mtype = CREATE_PLAYER

	jsObj, err := json.Marshal(*cm)
	if err == nil {
		return jsObj
	}
	return nil
}

func newMessageDestroyPlayer(pData *PlayerData) []byte {
	cm := new(MessageDestroyPlayer)
	cm.Player = *pData
	cm.Mtype = DESTROY_PLAYER

	jsObj, err := json.Marshal(*cm)
	if err == nil {
		return jsObj
	}
	return nil
}

func newMessageCreateGame(pData *PlayerData, otherPlayers []PlayerData) []byte {
	cm := new(MessageCreateGame)
	cm.Player = *pData
	cm.OtherPlayers = otherPlayers
	cm.Mtype = CREATE_GAME

	jsObj, err := json.Marshal(*cm)
	if err == nil {
		return jsObj
	} else {
		fmt.Println("err is", err)
	}
	return nil
}

func newMessageUpdateGame(g *Game) []byte {

	var b []PlayerData
	for cli := range g.clients {
		b = append(b, *cli.player)
	}
	fmt.Println(b)
	updateMessage := &MessageUpdateGame{
		Players: b[0:len(g.clients)],
		Mtype:   UPDATE_GAME,
	}

	jsObj, err := json.Marshal(updateMessage)
	if err == nil {
		return jsObj
	} else {
		fmt.Println(err.Error())

	}
	return nil

}
