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

	assert.Equal(t, result)

}
