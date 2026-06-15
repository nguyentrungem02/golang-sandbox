package cat

import (
	"errors"
	"strings"
)

type Cat struct {
	Name string `json:"name"`
}

// New Create a constructor
func New(name string) (*Cat, error) {
	name = strings.TrimSpace(name)

	if name == "" {
		return nil, errors.New("name cannot is empty")
	}

	if len(name) > 50 {
		return nil, errors.New("name too long (max 50 characters)")
	}

	return &Cat{
		Name: name,
	}, nil
}

func (c *Cat) GetName() string {
	return c.Name
}

func (c *Cat) Speak() string {
	return "Meo meo"
}

func (c *Cat) Eat() string {
	return "Meo an ca"
}
