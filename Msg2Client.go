package main

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
	creatorMessage CreatorMessage
	updateMessage  UpdateMessage
}
