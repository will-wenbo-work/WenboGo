package main

import (
	"sync"
)

var counter int = 1
var chars string = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
var mu sync.Mutex

func getUUID() string {
	mu.Lock()
	var currIdStr string
	var curr = counter
	counter++
	for curr > 0 {
		curr--
		currIdStr = currIdStr + string(chars[curr%62])
		curr /= 62
	}
	mu.Unlock()
	return currIdStr
}
