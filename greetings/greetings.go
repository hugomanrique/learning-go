package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {

	if name == "" {
		return "", errors.New("name can't be empty")
	}
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func randomFormat() string {
	formats := []string{
		"Hello, %v Welcome",
		"Start aplication, Welcome %v",
		"Welcome %v",
		"Welcome, %v",
		"Welcome, %v, to the application",
		"Welcome, %v, nice to meet you",
	}
	return formats[rand.Intn(len(formats))]
}

func HellosMultiple(names []string) (map[string]string, error) {
	messages := make(map[string]string)
	for _, name := range names {
		message, error := Hello(name)
		if error != nil {
			return nil, error
		}
		messages[name] = message
	}
	return messages, nil
}
