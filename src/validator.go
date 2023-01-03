package main

import (
	"errors"
	"net/mail"
)

func validateReq(newEvent event) (bool, error) {
	if newEvent.Title == "" {
		return false, errors.New("invalid title")
	}

	if newEvent.Version == "" {
		return false, errors.New("invalid version")
	}

	if newEvent.Company == "" {
		return false, errors.New("invalid company")
	}

	if newEvent.Website == "" {
		return false, errors.New("invalid website")
	}

	if newEvent.Source == "" {
		return false, errors.New("invalid source")
	}

	if newEvent.License == "" {
		return false, errors.New("invalid license")
	}

	if newEvent.Description == "" {
		return false, errors.New("invalid description")
	}

	if len(newEvent.Maintainers) == 0 {
		return false, errors.New("invalid Maintainer")
	} else {
		for _, maintainer := range newEvent.Maintainers {
			if !isValidEmailAddress(maintainer.Email) {
				return false, errors.New("invalid maintainer email")
			}
			if maintainer.Name == "" {
				return false, errors.New("invalid maintainer name")
			}
		}
	}

	return true, nil
}

func isValidEmailAddress(emailAddress string) bool {
	_, err := mail.ParseAddress(emailAddress)
	return err == nil
}
