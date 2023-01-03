package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIfEventExist(t *testing.T) {
	var savedEvent = getSavedEvent()
	SaveEvent(savedEvent)

	var unsavedEvent = getUnsavedEvent()

	assert.True(t, isEventExist(savedEvent))
	assert.False(t, isEventExist(unsavedEvent))

}

func getSavedEvent() event {
	var testMaintainer maintainer
	testMaintainer.Email = "aaa@gmail.com"
	testMaintainer.Name = "zhaowenbo"

	var testEvent event
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

func getUnsavedEvent() event {
	var testMaintainer maintainer
	testMaintainer.Email = "aaa@gmail.com"
	testMaintainer.Name = "zhaowenbo"

	var testEvent event
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
