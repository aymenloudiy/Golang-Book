package greeting

import (
	"errors"
	"fmt"
)

func Hello(name string) (string,error) {
	// If there is no name given return an empty string and a new error
	if name == "" {
		return "",errors.New("empty name")
	}
	// If a name was received, return a value that embeds the name in a greeting message.
	message := fmt.Sprintf("Hi, %v. Welcome!" ,name)
	return message,nil
}