package greeting

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string,error) {
	// If there is no name given return an empty string and a new error
	if name == "" {
		return "",errors.New("empty name")
	}
	// If a name was received, return a value that embeds the name in a greeting message.
	message := fmt.Sprintf(randomFormat(),name)
	return message,nil
}
// Hellos returns a map that associates each of the named people with a greeting message.
func Hellos(names []string)(map[string]string,error)  {
	messages := make(map[string]string)
	// Loop through the received slice of names, calling the Hello function to get a message for each name.
	// There is a blank operator here because the keys of names (index) are not needed, but the name value is.
	for _,name := range names {
		message,err := Hello(name)
		if err!=nil {
			return nil,err
		}
	// In the map, associate the retrieved message with the name.
		messages[name]=message
	}
	return messages,nil
}
// randomFormat returns one of a set of greeting messages. The returned message is selected at random.
func randomFormat() string {
	// A slice of message formats.
formats := []string{
	"Hi, %v. Welcome!",
	"Great to see you, %v!",
	"Hail, %v! Well met!",
}
    // Return a randomly selected message format by specifying a random index for the slice of formats.
return formats[rand.Intn(len(formats))]
}