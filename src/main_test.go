package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIfEventExist(t *testing.T) {
	var savedEvent = getSavedEvent()
	SavePayload(savedEvent)

	var unsavedEvent = getUnsavedEvent()

	assert.True(t, isPayloadExist(savedEvent))
	assert.False(t, isPayloadExist(unsavedEvent))

}

func getSavedEvent() payload {
	var testMaintainer maintainer
	testMaintainer.Email = "aaa@gmail.com"
	testMaintainer.Name = "zhaowenbo"

	var testEvent payload
	testEvent.Title = "title"
	testEvent.Company = "Company"
	testEvent.Description = "Description"
	testEvent.License = "license"
	testEvent.Source = "Ssource"
	testEvent.Version = "vVersion"
	testEvent.Website = "www.holiday_destroyer.com"
	testEvent.Maintainers = append(testEvent.Maintainers, testMaintainer)

	return testEvent
}

func getUnsavedEvent() payload {
	var testMaintainer maintainer
	testMaintainer.Email = "aaa@gmail.com"
	testMaintainer.Name = "zhaowenbo"

	var testEvent payload
	testEvent.Title = "unsaved event title"
	testEvent.Company = "Company"
	testEvent.Description = "Description"
	testEvent.License = "license"
	testEvent.Source = "Ssource"
	testEvent.Version = "vVersion"
	testEvent.Website = "www.holiday_destroyer.com"
	testEvent.Maintainers = append(testEvent.Maintainers, testMaintainer)

	return testEvent
}
