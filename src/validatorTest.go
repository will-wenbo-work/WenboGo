package main

import (
	"testing"
)

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
