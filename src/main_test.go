package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIfEventExist(t *testing.T) {
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

	var testEventWithoutSave event
	testEventWithoutSave.Title = "titleWithoutSave"
	testEventWithoutSave.Company = "Company"
	testEventWithoutSave.Description = "Description"
	testEventWithoutSave.License = "license"
	testEventWithoutSave.Source = "Ssource"
	testEventWithoutSave.Version = "vVersion"
	testEventWithoutSave.Website = "www.holiday_destroyer.com"
	testEventWithoutSave.Maintainers = append(testEvent.Maintainers, testMaintainer)

	SaveEvent(testEvent)

	assert.True(t, isEventExist(testEvent))
	assert.False(t, isEventExist(testEventWithoutSave))

}
