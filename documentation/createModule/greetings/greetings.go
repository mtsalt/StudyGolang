package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// // Hello returns a greeting for the named person.
// func Hello(name string) (string, error) {
// 	// If no name was given, return an error with message.
// 	if name == "" {
// 		return "", errors.New("empty name")
// 	}
// 	// Create a message using a random format.
// 	message := fmt.Sprintf("Hi, %v. Welcome!", name)
// 	return message, nil
// }

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	// If no name was given, return an error with a message.
	if name == "" {
		return name, errors.New("empty name")
	}
	// Create a message using a random format.
	// message := fmt.Sprintf(randomFormat(), name)
	message := fmt.Sprint(randomFormat())
	return message, nil
}

// Hellos returns a map that associations each of the named people
// with a greeting message.
func Hellos(names []string) (map[string]string, error) {
	// A map to associate names with messages.
	messages := make(map[string]string)
	// Loop through the received slice of names, calling
	// the Hello Ffunction to get a message for each name.
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		// In the map, associate the retrieved message with
		// the name.
		messages[name] = message
	}
	return messages, nil
}

// init  sets initial values for variables used in the function.
func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
	// A slice of message formats.
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v",
		"Hail, %v! Well met!",
	}

	// Return a randomly selected message format by specifying
	// a random index of the slice of formats.
	return formats[rand.Intn(len(formats))]
}
