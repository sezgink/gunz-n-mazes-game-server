package main

type CreatorMessage struct {
	players2Create []PlayerData
}
type UpdateMessage struct {
	players2Update []PlayerData
}
type Msg2Client struct {
	creatorMessage CreatorMessage
	updateMessage  UpdateMessage
}
