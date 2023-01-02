package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFlattenNormalEventData(t *testing.T) {
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

	var result = FlatenEvent2EventSearchParam(testEvent)

	assert.Equal(t, "title", result.Title)
	assert.Equal(t, "Company", result.Company)
	assert.Equal(t, "Description", result.Description)
	assert.Equal(t, "license", result.License)
	assert.Equal(t, "Ssource", result.Source)
	assert.Equal(t, "vVersion", result.Version)
	assert.Equal(t, "www.holiday_destroyer.com", result.Website)
	assert.Equal(t, "zhaowenbo", result.MaintainersNames[0])
	assert.Equal(t, 1, len(result.MaintainersNames))
	assert.Equal(t, "aaa@gmail.com", result.MaintainersEmails[0])
	assert.Equal(t, 1, len(result.MaintainersNames))
}
