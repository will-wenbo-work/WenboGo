package main

import "strconv"

var curr int = 0

func getUUID() string {

	currIdStr := strconv.Itoa(curr)
	curr++
	return currIdStr
}
