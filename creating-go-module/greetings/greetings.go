package greetings

import (
	"errors"
	"fmt"
	"math/rand"

)

func Hello(name string) (string, error){
	if name == "" {
		return "", errors.New("empty name")
	}
	message := fmt.Sprintf(random_format(), name)
	return message, nil
}

func Hellos(names []string) (map[string] string, error) {
	messages := make(map[string]string)

	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, errors.New("empty name")
		}
		messages[name] = message
	}
	return messages, nil
}

func random_format() string {
	formats := []string {
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v!, well met!",
	}
	return formats[rand.Intn(len(formats))]
}