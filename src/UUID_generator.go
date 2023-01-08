package main

var counter int = 1
var chars string = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func getUUID() string {

	var currIdStr string
	var curr = counter
	counter++
	for curr > 0 {
		curr--
		currIdStr = currIdStr + string(chars[curr%62])
		curr /= 62
	}
	return currIdStr
}
