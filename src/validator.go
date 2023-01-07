package main

import (
	"errors"
	"net/mail"
)

func validateReq(newPayload payload) (bool, error) {
	if newPayload.Title == "" {
		return false, errors.New("invalid title")
	}

	if newPayload.Version == "" {
		return false, errors.New("invalid version")
	}

	if newPayload.Company == "" {
		return false, errors.New("invalid company")
	}

	if newPayload.Website == "" {
		return false, errors.New("invalid website")
	}

	if newPayload.Source == "" {
		return false, errors.New("invalid source")
	}

	if newPayload.License == "" {
		return false, errors.New("invalid license")
	}

	if newPayload.Description == "" {
		return false, errors.New("invalid description")
	}

	if len(newPayload.Maintainers) == 0 {
		return false, errors.New("invalid Maintainer")
	} else {
		for _, maintainer := range newPayload.Maintainers {
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
