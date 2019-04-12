package main

import "encoding/json"

type PlayerData struct {
	id   int16
	posX float32
	posY float32
	vx   float32
	vy   float32
}

func ParsePlayerJSON(jsonData []byte) *PlayerData {
	var pData *PlayerData
	json.Unmarshal((jsonData), pData)
	return pData
}
