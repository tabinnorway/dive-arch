package formaterror

import (
	"errors"
	"strings"
)

func FormatError(err string) error {
	if strings.Contains(err, "email") {
		return errors.New("email already registered")
	}

	if strings.Contains(err, "title") {
		return errors.New("title already taken")
	}
	if strings.Contains(err, "hashedPassword") {
		return errors.New("incorrect username/password combination")
	}
	return errors.New("incorrect details")
}
