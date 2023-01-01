package main

import "net/mail"

var fields = [...]string{"Title", "Version", "Maintainers", "Company", "Website", "Source", "License", "Description"}

func validateReq(newEvent event) string {
	//WIP
	return ""
}

func validateEmailAddress(emailAddress string) bool {
	_, err := mail.ParseAddress(emailAddress)
	return err == nil
}
