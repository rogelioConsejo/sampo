package locator

import (
	"sampo/asset/container"
	"sampo/asset/container/location"
)

func Get() Locator {
	if locatorInstance == nil {
		l := locator{}
		locatorInstance = &l
	}
	return locatorInstance
}

type Locations map[location.Name]location.Location

//Locator is a unique (singleton) element that helps non-fungible elements (like and item.Item or a container.Container)
//have one unique location without having a reference to its location.Location in the item itself. This is to prevent
//two locations listing a same element or an element having more than one location.Location and at the same time keeping
//one single source of truth for any element location.Location. By having a Locator, we keep the responsibility
//of tracking an element's unique location.Location outside the element and outside the location.Location themselves.
type Locator interface {
	Locate(container.Container) location.Name
	Add(location.Location)
	Locations() Locations
}

type locator struct {
	locations Locations
}

func (l *locator) Locations() Locations {
	return l.locations
}

func (l *locator) Add(location location.Location) {
	if l.locations == nil {
		l.locations = make(Locations)
	}
	l.locations[location.Name()] = location
}

func (l *locator) Locate(_ container.Container) location.Name {
	return ""
}

var locatorInstance Locator
