package main

import "encoding/json"

type PlayerData struct {
	Id   int16
	PosX float32
	PosY float32
	Rot  float32
	Vx   float32
	Vy   float32
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
