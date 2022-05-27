package tools

import "regexp"

var email_regex = regexp.MustCompile(`/^[A-Z0-9._%+-]+@[A-Z0-9-]+.+.[A-Z]{2,4}$/i`)
var phone_regex = regexp.MustCompile(`^((\+7|7|8)+([0-9]){10})$`)

func EmailValid(email string) bool {
	return email_regex.MatchString(email)
}

func PhoneValid(phone string) bool {
	return phone_regex.MatchString(phone)
}
