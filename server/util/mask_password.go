package util

import (
	"regexp"
)

// MaskPasswword takes string and replaces each character with *
// Yes - this gives away some knowledge about the password (length),
// but this function is intended to be used in locations where that is OK.
func MaskPassword(password string) (maskedPassword string) {
	regex := regexp.MustCompile(".")
	maskedPassword = regex.ReplaceAllString(password, "*")

	return maskedPassword
}
