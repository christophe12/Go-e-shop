package main

import "regexp"

func stringIsEmpty(s string) bool {
	stringEmptiness := false

	if s == "" {
		stringEmptiness = true
	}

	return stringEmptiness
}

func intIsEmpty(i int) bool {
	intEmptiness := false

	if i == 0 {
		intEmptiness = true
	}

	return intEmptiness
}

func isValidEmail(e string) bool {
	emailValidity := false
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	if re.MatchString(e) {
		emailValidity = true
	}

	return emailValidity
}
