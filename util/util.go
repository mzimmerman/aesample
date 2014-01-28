package util

import "log"

func init() {
	Log("init() inside util")
}

func Log(m string) {
	log.Println(m)
}
