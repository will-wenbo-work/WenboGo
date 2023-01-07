package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEventValidator(t *testing.T) {

	var goodevent = getGoodEvent()
	var badevent = getBadEvent()
	testCases := []struct {
		testevent      payload
		expectedResult bool
		expectedErr    error
	}{
		{
			testevent:      goodevent,
			expectedResult: true,
			expectedErr:    nil,
		},
		{
			testevent:      badevent,
			expectedResult: false,
			expectedErr:    errors.New("invalid title"),
		},
	}

	for _, tc := range testCases {
		result, err := validateReq(tc.testevent)

		assert.Equal(t, result, tc.expectedResult)
		assert.Equal(t, err, tc.expectedErr)

	}
}

func TestEmailValidator(t *testing.T) {
	expectedResult := [2]bool{
		true,
		false,
	}

	for i, email := range []string{
		"good@exmaple.com",
		"bad-example",
	} {
		if expectedResult[i] != isValidEmailAddress(email) {
			t.Errorf("error when verifing %v, Received %v, expected %v", email, isValidEmailAddress(email), expectedResult[i])
		}
	}

}

func getGoodEvent() payload {
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

func getBadEvent() payload {
	var testMaintainer maintainer
	testMaintainer.Email = "aaa@gmail.com"
	testMaintainer.Name = "zhaowenbo"

	var testEvent payload
	testEvent.Company = "Company"
	testEvent.Description = "Description"
	testEvent.License = "license"
	testEvent.Source = "Ssource"
	testEvent.Version = "vVersion"
	testEvent.Website = "www.holiday_destroyer.com"
	testEvent.Maintainers = append(testEvent.Maintainers, testMaintainer)

	return testEvent
}
