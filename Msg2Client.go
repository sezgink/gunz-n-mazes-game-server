package main

import "encoding/json"

type CreatorMessage struct {
	players2Create []PlayerData
	fires2Create   []FireData
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
