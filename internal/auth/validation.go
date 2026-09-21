package auth

import (
	"regexp"
	"strings"
	"unicode/utf8"
)

func validateRegistration(email, password string) map[string]string {
	errors := make(map[string]string)

	if isBlank(email) {
		errors["email"] = "email cannot be empty"
	} else if !validEmail(email) {
		errors["email"] = "must be a valid email"
	}

	if isBlank(password) {
		errors["password"] = "password cannot be empty"
	} else if lessThanMinChars(password, 8) {
		errors["password"] = "password must be at least 8 characters long"
	} else if exceedMaxBytes(password, 72) {
		errors["password"] = "password must not be more than 72 bytes long"
	}

	return errors
}

func validEmail(email string) bool {
	emailRx := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return emailRx.MatchString(email)
}

func lessThanMinChars(s string, length int) bool {
	return utf8.RuneCountInString(s) < length
}

func exceedMaxBytes(s string, length int) bool {
	return len(s) > length
}

func isBlank(s string) bool {
	return strings.TrimSpace(s) == ""
}
