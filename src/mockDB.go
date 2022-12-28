package main

import "sync"

// "github.com/elastic/go-elasticsearch/v8"
var events = allEvents{}
var titleMap = make(map[string][]int)
var versionMap = make(map[string][]int)
var maintainersMap = make(map[string][]int)
var companyMap = make(map[string][]int)
var websiteMap = make(map[string][]int)
var sourceMap = make(map[string][]int)
var licenseMap = make(map[string][]int)
var descriptionMap = make(map[string][]int)
var mu sync.Mutex

func SaveEvent(newEvent event) {
	mu.Lock()
	NewEventDataDedup()
	events = append(events, newEvent)
	IndexingEachField(len(events), newEvent)
	mu.Unlock()
}

func NewEventDataDedup(newEvent event) {
	//
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

	versionList, ok := versionMap[newEvent.version]
	if ok {
		versionMap[newEvent.version] = append(versionList, id)
	} else {
		var newVersionList []int
		newVersionList = append(newVersionList, id)
		titleMap[newEvent.version] = newVersionList
	}

	companyList, ok := companyMap[newEvent.company]
	if ok {
		companyMap[newEvent.company] = append(companyList, id)
	} else {
		var newCompanyList []int
		newCompanyList = append(newCompanyList, id)
		titleMap[newEvent.version] = newCompanyList
	}

	websitList, ok := websiteMap[newEvent.website]
	if ok {
		websiteMap[newEvent.website] = append(websitList, id)
	} else {
		var websitList []int
		websitList = append(websitList, id)
		titleMap[newEvent.version] = websitList
	}

	sourceList, ok := sourceMap[newEvent.source]
	if ok {
		sourceMap[newEvent.source] = append(sourceList, id)
	} else {
		var sourceList []int
		sourceList = append(sourceList, id)
		titleMap[newEvent.version] = sourceList
	}

	licenseList, ok := licenseMap[newEvent.license]
	if ok {
		licenseMap[newEvent.license] = append(licenseList, id)
	} else {
		var licenseList []int
		licenseList = append(licenseList, id)
		licenseMap[newEvent.version] = licenseList
	}

	descriptionList, ok := descriptionMap[newEvent.Description]
	if ok {
		descriptionMap[newEvent.license] = append(descriptionList, id)
	} else {
		var descriptionList []int
		descriptionList = append(descriptionList, id)
		descriptionMap[newEvent.version] = descriptionList
	}
}

func SearchEventByField(eventParams event) {
	resultSet := make(map[int]bool)

	//title
	if eventParams.Title != "" {
		// titleList := companyMap[eventParams.Title]
		for _, id := range companyMap[eventParams.Title] {
			resultSet[id] = true
		}
	}

	//version
	if eventParams.version != "" {
		// versionList := versionMap[eventParams.version]
		tempSet := make(map[int]bool)
		for _, id := range versionMap[eventParams.version] {
			tempSet[id] = true
		}

		for k, _ := range resultSet {
			if !tempSet[k] {
				resultSet[k] = false
			}
		}
	}

	//company
	if eventParams.company != "" {
		tempSet := make(map[int]bool)
		for _, id := range companyMap[eventParams.version] {
			tempSet[id] = true
		}

		for k, _ := range resultSet {
			if !tempSet[k] {
				resultSet[k] = false
			}
		}
	}

	//website
	if eventParams.website != "" {
		tempSet := make(map[int]bool)
		for _, id := range websiteMap[eventParams.website] {
			tempSet[id] = true
		}

		for k, _ := range resultSet {
			if !tempSet[k] {
				resultSet[k] = false
			}
		}
	}

	//source
	if eventParams.source != "" {
		tempSet := make(map[int]bool)
		for _, id := range websiteMap[eventParams.source] {
			tempSet[id] = true
		}

		for k, _ := range resultSet {
			if !tempSet[k] {
				resultSet[k] = false
			}
		}
	}

	//license
	if eventParams.license != "" {
		tempSet := make(map[int]bool)
		for _, id := range websiteMap[eventParams.license] {
			tempSet[id] = true
		}

		for k, _ := range resultSet {
			if !tempSet[k] {
				resultSet[k] = false
			}
		}
	}

	//description
	if eventParams.Description != "" {
		tempSet := make(map[int]bool)
		for _, id := range websiteMap[eventParams.Description] {
			tempSet[id] = true
		}

		for k, _ := range resultSet {
			if !tempSet[k] {
				resultSet[k] = false
			}
		}
	}
}
