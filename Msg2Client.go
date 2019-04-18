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
)

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

func newMessageCreatePlayer(pData *PlayerData) []byte {
	cm := new(MessageCreatePlayer)
	cm.Player = *pData
	cm.IsOwner = false
	cm.Mtype = UPDATE_GAME

	// func JSONMsg2Client(msg *Msg2Client) []byte {
	// 	//plr, err := json.Marshal(player)
	// 	plr, err := json.Marshal(msg)
	// 	if err == nil {
	// 		return plr
	// 	}
	// 	return nil
	// }
	// func JSONUpdate2Client(msg *UpdateMessage) []byte {
	// 	//plr, err := json.Marshal(player)
	// 	plr, err := json.Marshal(msg)
	// 	if err == nil {
	// 		return plr
	// 	}
	// 	return nil
	// }

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

func createUpdateMessage(g *Game) []byte {
	// /*
	// 	b := make([]PlayerData, len(g.clients))
	// 	for cli := range g.clients {
	// 		b = append(b, *cli.player)
	// 	}
	// */
	// b := make([]PlayerData, len(g.clients))
	// for cli := range g.clients {
	// 	b = append(b, *cli.player)
	// }
	// updateMessage := &MessageUpdateGame{
	// 	players2Update: b[0:len(g.clients)],
	// 	flag:           1,
	// }

	// plr, err := json.Marshal(updateMessage)
	// fmt.Println(plr)

	// if err == nil {
	// 	return plr
	// } else {
	// 	fmt.Println(err.Error())

	// }
	return nil

}
