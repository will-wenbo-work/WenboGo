package main

import "net/mail"

var fields = [...]string{"Title", "Version", "Maintainers", "Company", "Website", "Source", "License", "Description"}

func validateReq(newEvent event) string {
	if newEvent.Title == "" {
		return "title is empty"
	}

	if newEvent.Version == "" {
		return "version is empty"
	}

	if newEvent.Company == "" {
		return "company is empty"
	}

	if newEvent.Website == "" {
		return "website is empty"
	}

	if newEvent.Source == "" {
		return "source is empty"
	}

	if newEvent.License == "" {
		return "license is empty"
	}

	if newEvent.Description == "" {
		return "description is empty"
	}

	if len(newEvent.Maintainers) == 0 {
		return "invalid maintainer"
	} else {
		for _, maintainer := range newEvent.Maintainers {
			if !isValidEmailAddress(maintainer.Email) {
				return "maintainer email address is not valid"
			}
			if maintainer.Name == "" {
				return "maintainer name is not valid"
			}
		}
	}

	return ""
}

func isValidEmailAddress(emailAddress string) bool {
	_, err := mail.ParseAddress(emailAddress)
	return err == nil
}
