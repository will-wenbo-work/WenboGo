package main

import "sync"

// "github.com/elastic/go-elasticsearch/v8"
var events = allEvents{}
var titleMap = make(map[string][]int)

// var titleMap map[string][]int
var versionMap map[string][]int
var maintainersMap map[string][]int
var company map[string][]int
var website map[string][]int
var source map[string][]int
var license map[string][]int
var description map[string][]int
var mu sync.Mutex

func SaveEvent(newEvent event) {
	mu.Lock()
	events = append(events, newEvent)
	IndexingEachField(len(events), newEvent)
	mu.Unlock()
}

func IndexingEachField(id int, newEvent event) {
	titleList, ok := titleMap[newEvent.Title]
	if ok {
		titleMap[newEvent.Title] = append(titleList, id)
	} else {
		var newTitleList []int
		newTitleList = append(newTitleList, id)
		titleMap[newEvent.Title] = newTitleList
	}
}

func SearchEventByField() {
	//number from title
	//number from version
	//..
	//find set interaction
}
