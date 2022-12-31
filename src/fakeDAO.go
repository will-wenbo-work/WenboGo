package main

import "sync"

var eventsDB = allEvents{}
var titleMap = make(map[string][]int)
var versionMap = make(map[string][]int)
var maintainersNameMap = make(map[string][]int)
var maintainersEmailMap = make(map[string][]int)
var companyMap = make(map[string][]int)
var websiteMap = make(map[string][]int)
var sourceMap = make(map[string][]int)
var licenseMap = make(map[string][]int)
var descriptionMap = make(map[string][]int)
var mu sync.Mutex

// save the event, lock the "table" when writing.
func SaveEvent(newEvent event) {
	mu.Lock()
	eventsDB = append(eventsDB, newEvent)
	IndexingEachField(newEvent, len(eventsDB)-1)
	mu.Unlock()
}

// we have hashmap for each field, key-value: <field name - eventID>, eventID is len(events)
func IndexingEachField(newEvent event, id int) {
	titleList, ok := titleMap[newEvent.Title]
	if ok {
		titleMap[newEvent.Title] = append(titleList, id)
	} else {
		var newTitleList []int
		newTitleList = append(newTitleList, id)
		titleMap[newEvent.Title] = newTitleList
	}

	versionList, ok := versionMap[newEvent.Version]
	if ok {
		versionMap[newEvent.Version] = append(versionList, id)
	} else {
		var newVersionList []int
		newVersionList = append(newVersionList, id)
		versionMap[newEvent.Version] = newVersionList
	}

	companyList, ok := companyMap[newEvent.Company]
	if ok {
		companyMap[newEvent.Company] = append(companyList, id)
	} else {
		var newCompanyList []int
		newCompanyList = append(newCompanyList, id)
		companyMap[newEvent.Company] = newCompanyList
	}

	websitList, ok := websiteMap[newEvent.Website]
	if ok {
		websiteMap[newEvent.Website] = append(websitList, id)
	} else {
		var websitList []int
		websitList = append(websitList, id)
		websiteMap[newEvent.Website] = websitList
	}

	sourceList, ok := sourceMap[newEvent.Source]
	if ok {
		sourceMap[newEvent.Source] = append(sourceList, id)
	} else {
		var sourceList []int
		sourceList = append(sourceList, id)
		sourceMap[newEvent.Source] = sourceList
	}

	licenseList, ok := licenseMap[newEvent.License]
	if ok {
		licenseMap[newEvent.License] = append(licenseList, id)
	} else {
		var licenseList []int
		licenseList = append(licenseList, id)
		licenseMap[newEvent.License] = licenseList
	}

	descriptionList, ok := descriptionMap[newEvent.Description]
	if ok {
		descriptionMap[newEvent.Description] = append(descriptionList, id)
	} else {
		var descriptionList []int
		descriptionList = append(descriptionList, id)
		descriptionMap[newEvent.Description] = descriptionList
	}

	for _, v := range newEvent.Maintainers {
		maintainerNameList, ok := maintainersNameMap[v.Name]
		if ok {
			maintainersNameMap[v.Name] = append(maintainerNameList, id)
		} else {
			var maintainerNameList []int
			maintainerNameList = append(maintainerNameList, id)
			maintainersNameMap[v.Name] = maintainerNameList
		}

		maintainerEmailList, ok := maintainersEmailMap[v.Name]
		if ok {
			maintainersEmailMap[v.Email] = append(maintainerEmailList, id)
		} else {
			var maintainerEmailList []int
			maintainerEmailList = append(maintainerEmailList, id)
			maintainersEmailMap[v.Email] = maintainerEmailList
		}
	}
}

// search by parameter, we get the list of eventIDs for each field, then get the intersection of all the lists. the intersection is the result of search
func searchEventByField(eventParams eventSearchParam) []int {
	resultSet := make(map[int]bool)
	for k := range eventsDB {
		resultSet[k] = true
	}

	var resultList []int

	//title
	if eventParams.Title != "" {
		// titleList := companyMap[eventParams.Title]
		for _, id := range titleMap[eventParams.Title] {
			resultSet[id] = true
		}
	}

	//version
	if eventParams.Version != "" {
		// versionList := versionMap[eventParams.version]
		tempSet := make(map[int]bool)
		for _, id := range versionMap[eventParams.Version] {
			tempSet[id] = true
		}

		for k, _ := range resultSet {
			if !tempSet[k] {
				resultSet[k] = false
			}
		}
	}

	//company
	if eventParams.Company != "" {
		tempSet := make(map[int]bool)
		for _, id := range companyMap[eventParams.Company] {
			tempSet[id] = true
		}

		for k, _ := range resultSet {
			if !tempSet[k] {
				resultSet[k] = false
			}
		}
	}

	//website
	if eventParams.Website != "" {
		tempSet := make(map[int]bool)
		for _, id := range websiteMap[eventParams.Website] {
			tempSet[id] = true
		}

		for k, _ := range resultSet {
			if !tempSet[k] {
				resultSet[k] = false
			}
		}
	}

	//source
	if eventParams.Source != "" {
		tempSet := make(map[int]bool)
		for _, id := range sourceMap[eventParams.Source] {
			tempSet[id] = true
		}

		for k, _ := range resultSet {
			if !tempSet[k] {
				resultSet[k] = false
			}
		}
	}

	//license
	if eventParams.License != "" {
		tempSet := make(map[int]bool)
		for _, id := range licenseMap[eventParams.License] {
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
		for _, id := range descriptionMap[eventParams.Description] {
			tempSet[id] = true
		}

		for k, _ := range resultSet {
			if !tempSet[k] {
				resultSet[k] = false
			}
		}
	}

	for _, maintainerName := range eventParams.maintainersNames {
		tempSet := make(map[int]bool)
		for _, id := range maintainersNameMap[maintainerName] {
			tempSet[id] = true
		}

		for k, _ := range resultSet {
			if !tempSet[k] {
				resultSet[k] = false
			}
		}
	}

	for _, maintainerEmail := range eventParams.MaintainersEmails {
		tempSet := make(map[int]bool)
		for _, id := range maintainersEmailMap[maintainerEmail] {
			tempSet[id] = true
		}

		for k, _ := range resultSet {
			if !tempSet[k] {
				resultSet[k] = false
			}
		}
	}

	//merge intersection
	for k, _ := range resultSet {
		if resultSet[k] {
			resultList = append(resultList, k)
		}
	}

	return resultList
}
