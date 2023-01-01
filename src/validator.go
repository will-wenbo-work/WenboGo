package main

import "net/mail"

var fields = [...]string{"Title", "Version", "Maintainers", "Company", "Website", "Source", "License", "Description"}

func validateReq(newEvent event) string {
	//WIP

	for _, maintainer := range newEvent.Maintainers {
		if !isValidEmailAddress(maintainer.Email) {
			return "email address is not valid"
		}
		if maintainer.Name == "" {
			return "maintainer name is empty"
		}
	}
	return ""
}

func isValidEmailAddress(emailAddress string) bool {
	_, err := mail.ParseAddress(emailAddress)
	return err == nil
}
