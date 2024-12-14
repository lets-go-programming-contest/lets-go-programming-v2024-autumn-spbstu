package contact

import "testing"

func TestValidatePhone(t *testing.T) {
	validPhones := []string{
		"+7(921)123-45-67",
		"+7921123-45-67",
		"+7-921-123-45-67",
		"8(921)123-45-67",
		"8921123-45-67",
		"8-921-123-45-67",
		"+79211234567",
		"89211234567",
		"+7 921 123 45 67",
		"8 921 123 45 67",
		"+7 9211234567",
		"8(921)123-45-67",
	}

	invalidPhones := []string{
		"8.921.123.45.67",
		"+7.921.123.45.67",
		"+79 (921) 123-45-67",
		"7 (921) 123-45-67",
		"+7-921.1234-567",
		//"8921123456",
		"8921",
		"892112312",
		"+792112345678",
	}

	for _, phone := range validPhones {
		if !validatePhone(phone) {
			t.Errorf("%s ok, but got bad", phone)
		}
	}

	for _, phone := range invalidPhones {
		if validatePhone(phone) {
			t.Errorf("%s bad, but got ok", phone)
		}
	}
}
