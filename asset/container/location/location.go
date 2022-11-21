// Package location defines a location which is a special form of container.Container, expected to be the root
// of a container.Container taxonomy. All container.Container should have a unique root Location
package location

import (
	"errors"
	"fmt"
	"sampo/asset/container"
	"strings"
)

// New creates a new Location with a given Name. The name cannot be empty.
func New(name Name) (Location, error) {
	l := Location{}
	err := l.SetName(name)
	if err != nil {
		return Location{}, fmt.Errorf("cous not set location name: %w", err)
	}
	return Location{name: name}, nil
}

// A Location works as a container.Container that has a unique name and is indexed by the Locator.
type Location struct {
	container.Container
	name Name
}

type Name string

func (l *Location) Name() Name {
	return l.name
}

func (l *Location) SetName(name Name) error {
	if strings.TrimSpace(string(name)) == "" {
		return emptyNameError
	}
	l.name = name
	return nil
}

var emptyNameError = errors.New("name cannot be empty")
