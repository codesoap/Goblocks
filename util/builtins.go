package util

import (
	"time"
)

//blockId is automatically allocated
//send channel is used to update blocks data
//Gothreads are waked up by messages on rec channels
//action is a map of whatever was in action json object for corressponding action
var FunctionMap = map[string]func(blockId int, send chan Change, rec chan bool, action map[string]interface{}) {
	"#Date": Date,
}

//Update time according to "format" property
func Date(blockId int, send chan Change, rec chan bool, action map[string]interface{}) {
	run := true
	for run {
		send <- Change{blockId, time.Now().Format(action["format"].(string)), true}
		//Block until other thread will ping you
		run = <-rec
	}
}
