package container

import "sampo/asset/item"

func New() Container {
	return Container{}
}

type Container struct {
	items items
}

func (c *Container) Add(i item.Item) {
	if c.items == nil {
		c.items = make(items)
	}
	c.items[i] = nil
}

func (c *Container) CheckIf(i item.Item) ItemPresenceChecker {
	return itemPresenceChecker{container: c, item: i}
}

func (c *Container) checkIfItemIsStored(i item.Item) bool {
	_, found := c.items[i]
	return found
}

type ItemPresenceChecker interface {
	PresenceChecker
}

type PresenceChecker interface {
	IsPresent() bool
}

type itemPresenceChecker struct {
	container *Container
	item      item.Item
}

func (i itemPresenceChecker) IsPresent() bool {
	return i.container.checkIfItemIsStored(i.item)
}

type items map[item.Item]interface{}
