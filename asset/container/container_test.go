package container

import (
	"sampo/asset/item"
	"testing"
)

func TestContainer(t *testing.T) {
	var c Container = New()

	expectedItem := item.Item{Id: "some id"}
	unexpectedItem := item.Item{Id: "some other id"}
	c.Add(expectedItem)

	if expectedItemNotFound(c, expectedItem) {
		t.Errorf("expected item not found!")
	}
	if unexpectedItemFound(c, unexpectedItem) {
		t.Errorf("unexpected item found!")
	}
}

//TODO
func TestContainer_ShouldCheckForAnItemsPresenceInsideItsChildren(t *testing.T) {

}

func unexpectedItemFound(c Container, unexpectedItem item.Item) bool {
	return c.CheckIf(unexpectedItem).IsPresent()
}

func expectedItemNotFound(c Container, expectedItem item.Item) bool {
	return !c.CheckIf(expectedItem).IsPresent()
}
