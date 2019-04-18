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

/*
type CreatorMessage struct {
	isOwner bool
	//player2Create *PlayerData
	//This one probably should go to data than poiniter
	player2Create PlayerData

	//fires2Create  []FireData
	flag int
}
*/
type CreatorMessage struct {
	IsOwner bool
	//player2Create *PlayerData
	//This one probably should go to data than poiniter
	Player2Create PlayerData

	//fires2Create  []FireData
	Flag int
}

type MultiCreatorMessage struct {
	//IsOwner bool
	//player2Create *PlayerData
	//This one probably should go to data than poiniter
	OwningPlayer   PlayerData
	Players2Create []PlayerData

	//fires2Create  []FireData
	Flag int
}

func CreateMultiCreatorMessage(pData *PlayerData, otherPlayers []PlayerData) []byte {
	//plr, err := json.Marshal(player)
	//pData.isOwner = true

	cm := new(MultiCreatorMessage)
	cm.OwningPlayer = *pData
	cm.Players2Create = otherPlayers
	//cm.IsOwner = true

	cm.Flag = 4

	//fmt.Println("Cm is", *cm)

	plr, err := json.Marshal(*cm)
	if err == nil {
		//fmt.Println("Plr is", plr)
		return plr
	} else {
		fmt.Println("err is", err)
	}
	return nil
}

func CreatePlrCreatorMessage(pData *PlayerData) []byte {
	//plr, err := json.Marshal(player)
	//pData.isOwner = true

	cm := new(CreatorMessage)
	cm.Player2Create = *pData
	cm.IsOwner = true

	cm.Flag = 0

	//fmt.Println("Cm is", *cm)

	plr, err := json.Marshal(*cm)
	if err == nil {
		//fmt.Println("Plr is", plr)
		return plr
	} else {
		fmt.Println("err is", err)
	}
	return nil
}
func CreateCreatorMessage(pData *PlayerData) []byte {
	//plr, err := json.Marshal(player)

	cm := new(CreatorMessage)
	cm.Player2Create = *pData
	cm.IsOwner = false

	cm.Flag = 0

	plr, err := json.Marshal(*cm)
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
