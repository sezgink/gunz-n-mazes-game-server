package main

type CreatorMessage struct {
	player2Create PlayerData
}
type UpdateMessage struct {
}

type Msg2Client struct {
	players []PlayerData
}
