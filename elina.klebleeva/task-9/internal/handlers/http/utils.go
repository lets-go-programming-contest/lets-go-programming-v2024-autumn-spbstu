package handlers

import (
	"regexp"
)

func isValidPhone(phone string) bool {

	phoneReg := `^(8|7|\+7)\s(\(\d{3}\))\s(\d{3})\-(\d{2})\-(\d{2})$`
	re := regexp.MustCompile(phoneReg)

	return re.Find([]byte(phone)) != nil
}

func isEmpty(str string) bool {
	return str == ""
}
