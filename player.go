package main

import "encoding/json"

type PlayerData struct {
	id        int16
	posX      float32
	posY      float32
	rot       float32
	vx        float32
	vy        float32
	isCreated bool
}

func ParsePlayerJSON(jsonData []byte) *PlayerData {
	var pData *PlayerData
	json.Unmarshal((jsonData), pData)
	return pData
}
func JSONPlayerData(player *PlayerData) []byte {
	//plr, err := json.Marshal(player)
	plr, err := json.Marshal(player)
	if err == nil {
		return plr
	}
	return nil
}
