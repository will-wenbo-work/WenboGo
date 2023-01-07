package main

import (
	"errors"
	"sync"
)

var payloadDB = allpayload{}
var title2IdMap = make(map[string][]string)
var version2IdMap = make(map[string][]string)
var maintainersName2IdMap = make(map[string][]string)
var maintainersEmail2IdMap = make(map[string][]string)
var company2IdMap = make(map[string][]string)
var website2IdMap = make(map[string][]string)
var source2IdMap = make(map[string][]string)
var license2IdMap = make(map[string][]string)
var description2IdMap = make(map[string][]string)
var id2PayloadMap = make(map[string]payload)
var mu sync.Mutex

// save the event, lock the "table" when writing.
func SavePayload(newPayload payload) {
	mu.Lock()
	payloadDB = append(payloadDB, newPayload)
	id2PayloadMap[newPayload.id] = newPayload
	IndexingEachField(newPayload)
	mu.Unlock()
}

func fetchPayload(id string) (payload, error) {
	val, ok := id2PayloadMap[id]
	// If the key exists
	if ok {
		return val, nil
	} else {
		return val, errors.New("invalid payload ID")
	}

}

// we have hashmap for each field, key-value: <field name - eventID>, eventID is len(events)
func IndexingEachField(newPayload payload) {
	titleList, ok := title2IdMap[newPayload.Title]
	if ok {
		title2IdMap[newPayload.Title] = append(titleList, newPayload.id)
	} else {
		var newTitleList []string
		newTitleList = append(newTitleList, newPayload.id)
		title2IdMap[newPayload.Title] = newTitleList
	}

	versionList, ok := version2IdMap[newPayload.Version]
	if ok {
		version2IdMap[newPayload.Version] = append(versionList, newPayload.id)
	} else {
		var newVersionList []string
		newVersionList = append(newVersionList, newPayload.id)
		version2IdMap[newPayload.Version] = newVersionList
	}

	companyList, ok := company2IdMap[newPayload.Company]
	if ok {
		company2IdMap[newPayload.Company] = append(companyList, newPayload.id)
	} else {
		var newCompanyList []string
		newCompanyList = append(newCompanyList, newPayload.id)
		company2IdMap[newPayload.Company] = newCompanyList
	}

	websitList, ok := website2IdMap[newPayload.Website]
	if ok {
		website2IdMap[newPayload.Website] = append(websitList, newPayload.id)
	} else {
		var websitList []string
		websitList = append(websitList, newPayload.id)
		website2IdMap[newPayload.Website] = websitList
	}

	sourceList, ok := source2IdMap[newPayload.Source]
	if ok {
		source2IdMap[newPayload.Source] = append(sourceList, newPayload.id)
	} else {
		var sourceList []string
		sourceList = append(sourceList, newPayload.id)
		source2IdMap[newPayload.Source] = sourceList
	}

	licenseList, ok := license2IdMap[newPayload.License]
	if ok {
		license2IdMap[newPayload.License] = append(licenseList, newPayload.id)
	} else {
		var licenseList []string
		licenseList = append(licenseList, newPayload.id)
		license2IdMap[newPayload.License] = licenseList
	}

	descriptionList, ok := description2IdMap[newPayload.Description]
	if ok {
		description2IdMap[newPayload.Description] = append(descriptionList, newPayload.id)
	} else {
		var descriptionList []string
		descriptionList = append(descriptionList, newPayload.id)
		description2IdMap[newPayload.Description] = descriptionList
	}

	for _, v := range newPayload.Maintainers {
		maintainerNameList, ok := maintainersName2IdMap[v.Name]
		if ok {
			maintainersName2IdMap[v.Name] = append(maintainerNameList, newPayload.id)
		} else {
			var maintainerNameList []string
			maintainerNameList = append(maintainerNameList, newPayload.id)
			maintainersName2IdMap[v.Name] = maintainerNameList
		}

		maintainerEmailList, ok := maintainersEmail2IdMap[v.Name]
		if ok {
			maintainersEmail2IdMap[v.Email] = append(maintainerEmailList, newPayload.id)
		} else {
			var maintainerEmailList []string
			maintainerEmailList = append(maintainerEmailList, newPayload.id)
			maintainersEmail2IdMap[v.Email] = maintainerEmailList
		}
	}
}

func deleteRecord(payloadId string) bool {
	delete(id2PayloadMap, payloadId)

	for i, singlePayload := range payloadDB {
		if singlePayload.id == payloadId {
			payloadDB = append(payloadDB[:i], payloadDB[i+1:]...)
			return true
		}
	}
	return false
}

// search by parameter, we get the list of eventIDs for each field, then get the intersection of all the lists. the intersection is the result of search
func searchPayloadByField(eventParams payloadSearchParam) []string {
	resultSet := make(map[string]bool)
	for _, v := range payloadDB {
		resultSet[v.id] = true
	}

	var resultList []string

	//title
	if eventParams.Title != "" {
		tempSet := make(map[string]bool)
		for _, id := range title2IdMap[eventParams.Title] {
			tempSet[id] = true
		}

		for k, _ := range resultSet {
			if !tempSet[k] {
				resultSet[k] = false
			}
		}
	}

	//version
	if eventParams.Version != "" {
		tempSet := make(map[string]bool)
		for _, id := range version2IdMap[eventParams.Version] {
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
		tempSet := make(map[string]bool)
		for _, id := range company2IdMap[eventParams.Company] {
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
		tempSet := make(map[string]bool)
		for _, id := range website2IdMap[eventParams.Website] {
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
		tempSet := make(map[string]bool)
		for _, id := range source2IdMap[eventParams.Source] {
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
		tempSet := make(map[string]bool)
		for _, id := range license2IdMap[eventParams.License] {
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
		tempSet := make(map[string]bool)
		for _, id := range description2IdMap[eventParams.Description] {
			tempSet[id] = true
		}

		for k, _ := range resultSet {
			if !tempSet[k] {
				resultSet[k] = false
			}
		}
	}

	for _, maintainerName := range eventParams.MaintainersNames {
		tempSet := make(map[string]bool)
		for _, id := range maintainersName2IdMap[maintainerName] {
			tempSet[id] = true
		}

		for k, _ := range resultSet {
			if !tempSet[k] {
				resultSet[k] = false
			}
		}
	}

	for _, maintainerEmail := range eventParams.MaintainersEmails {
		tempSet := make(map[string]bool)
		for _, id := range maintainersEmail2IdMap[maintainerEmail] {
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

func deleteIdForEachIndex(eventId int) {
}
