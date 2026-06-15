package dog

import (
	"errors"
	"strings"
)

type Dog struct {
	Name string `json:"name"`
}

// New Create a constructor
func New(name string) (*Dog, error) {
	name = strings.TrimSpace(name)

	if name == "" {
		return nil, errors.New("name cannot is empty")
	}

	if len(name) > 50 {
		return nil, errors.New("name too long (max 50 characters)")
	}

	return &Dog{
		Name: name,
	}, nil
}

func (d *Dog) GetName() string {
	return d.Name
}

func (d *Dog) Speak() string {
	return "Gâu gâu"
}
