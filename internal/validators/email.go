package validators

import (
	"fmt"
	"regexp"
)

var emailRegexp = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

//EmailAddress validator
func EmailAddress(email string) error {
	if !emailRegexp.MatchString(email) {
		return fmt.Errorf("invalid email")
	}
	return nil
}
