package handlers

import (
	"regexp"
	"strings"
)

func isValidPhone(phone string) bool {

	e164Regex := `^((8|\+7)[\- ]?)?(\(?\d{3}\)?[\- ]?)?[\d\- ]{7,10}$`
	re := regexp.MustCompile(e164Regex)
	phone = strings.ReplaceAll(phone, " ", "")

	return re.Find([]byte(phone)) != nil
}

func isEmpty(str string) bool {
	return str == ""
}
