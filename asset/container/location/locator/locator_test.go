package locator

import (
	locator2 "sampo/asset/container/location"
	"testing"
)

func TestLocator_IsUnique(t *testing.T) {
	var l = Get()
	var l2 = Get()

	areSameInstance := checkIfLocatorsPointToTheSameInstance(l2, l)
	if !areSameInstance {
		t.Fatalf("location is not unique")
	}
}

func TestLocator_Locations(t *testing.T) {
	l := Get()
	const testLocationName locator2.Name = "some location"
	testLocation, err := locator2.New(testLocationName)
	if err != nil {
		t.Fatalf(err.Error())
	}
	l.Add(testLocation)
	if _, found := l.Locations()[testLocationName]; !found {
		t.Fatal("expected location not found")
	}
}

func TestLocator_Locate(t *testing.T) {
}

func checkIfLocatorsPointToTheSameInstance(l2 Locator, l Locator) bool {
	return l2.(*locator) == l.(*locator)
}
