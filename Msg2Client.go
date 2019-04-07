package main

import "encoding/json"

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
}

func CreatePlrCreatorMessage(pData *PlayerData) []byte {
	//plr, err := json.Marshal(player)
	//pData.isOwner = true
	cm := new(CreatorMessage)
	cm.player2Create = pData
	cm.isOwner = true

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

	plr, err := json.Marshal(cm)
	if err == nil {
		return plr
	}
	return nil
}

type UpdateMessage struct {
	players2Update []PlayerData
}
type DestroyMessage struct {
	players2Destroy []PlayerData
}
type Msg2Client struct {
	creatorMessage *CreatorMessage
	updateMessage  *UpdateMessage
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
