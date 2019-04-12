package main

import (
	"encoding/json"
	"fmt"
)

/*
type CreatorMessage struct {
	players2Create []PlayerData
	fires2Create   []FireData
}
*/
type CreatorMessage struct {
	isOwner       bool
	player2Create *PlayerData
	//fires2Create  []FireData
	flag int
}

func CreatePlrCreatorMessage(pData *PlayerData) []byte {
	//plr, err := json.Marshal(player)
	//pData.isOwner = true

	cm := new(CreatorMessage)
	cm.player2Create = pData
	cm.isOwner = true

	cm.flag = 0

	plr, err := json.Marshal(cm)
	if err == nil {
		return plr
	}
	return nil
}
func CreateCreatorMessage(pData *PlayerData) []byte {
	//plr, err := json.Marshal(player)

	cm := new(CreatorMessage)
	cm.player2Create = pData
	cm.isOwner = false

	cm.flag = 0

	plr, err := json.Marshal(cm)
	if err == nil {
		return plr
	}
	return nil
}

type UpdateMessage struct {
	players2Update []PlayerData
	flag           int
}
type DestroyMessage struct {
	players2Destroy []PlayerData
}
type Msg2Client struct {
	creatorMessage *CreatorMessage
	updateMessage  *UpdateMessage
}

func createUpdateMessage(g *Game) []byte {
	/*
		b := make([]PlayerData, len(g.clients))
		for cli := range g.clients {
			b = append(b, *cli.player)
		}
	*/
	b := make([]PlayerData, len(g.clients))
	for cli := range g.clients {
		b = append(b, *cli.player)
	}
	updateMessage := &UpdateMessage{
		players2Update: b[0:len(g.clients)],
		flag:           1,
	}

	plr, err := json.Marshal(updateMessage)
	fmt.Println(plr)

	if err == nil {
		return plr
	} else {
		fmt.Println(err.Error())

	}
	return nil

}

func JSONMsg2Client(msg *Msg2Client) []byte {
	//plr, err := json.Marshal(player)
	plr, err := json.Marshal(msg)
	if err == nil {
		return plr
	}
	return nil
}
func JSONUpdate2Client(msg *UpdateMessage) []byte {
	//plr, err := json.Marshal(player)
	plr, err := json.Marshal(msg)
	if err == nil {
		return plr
	}
	return nil
}
