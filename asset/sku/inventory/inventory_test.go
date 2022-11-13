package inventory

import (
	"errors"
	"sampo/asset/item"
	"testing"
)

func TestInventory_TrackSeveralItems(t *testing.T) {
	t.Parallel()
	var testInventory Inventory = New()
	var testItems []item.Item = []item.Item{
		makeTestItem("test-item1"),
		makeTestItem("test-item2"),
		makeTestItem("test-item3"),
	}
	var sums [3]Amount

	addAndRemoveFromSeveralItemsStocks(t, testInventory, testItems, sums)
}

func TestInventory_RemoveStock_ThrowsErrorOnRemovingTooMuchStock(t *testing.T) {
	t.Parallel()
	var testInventory Inventory = New()
	testItem := item.Item{}
	removeStockFromNonExistingItem(t, testInventory, testItem)
	removeTooMuchStockFromExistingItem(t, testInventory, testItem)
}

func removeTooMuchStockFromExistingItem(t *testing.T, testInventory Inventory, testItem item.Item) {
	testInventory.AddStock(testItem, 100)
	err := testInventory.RemoveStock(testItem, 101)
	if err == nil {
		t.Fatalf("error not thrown on removing too much stock")
	}
	if !errors.Is(err, insufficientAmountError) {
		t.Fatalf("invalid error thrown: %s", err.Error())
	}
}

func removeStockFromNonExistingItem(t *testing.T, testInventory Inventory, testItem item.Item) {
	err := testInventory.RemoveStock(testItem, 100)
	if err == nil {
		t.Fatalf("error not thrown on removing too much stock")
	}
	if !errors.Is(err, insufficientAmountError) {
		t.Fatalf("invalid error thrown: %s", err.Error())
	}
}

func addAndRemoveFromSeveralItemsStocks(t *testing.T, testInventory Inventory, testItems []item.Item, sums [3]Amount) {
	testAddingAmount(t, testInventory, testItems[0], &sums[0], 1)
	testAddingAmount(t, testInventory, testItems[0], &sums[0], 1)
	testAddingAmount(t, testInventory, testItems[0], &sums[0], 1)

	testAddingAmount(t, testInventory, testItems[1], &sums[1], 10)
	testAddingAmount(t, testInventory, testItems[1], &sums[1], 10)
	testAddingAmount(t, testInventory, testItems[1], &sums[1], 10)

	testAddingAmount(t, testInventory, testItems[2], &sums[2], 1)
	testAddingAmount(t, testInventory, testItems[2], &sums[2], 10)
	testAddingAmount(t, testInventory, testItems[2], &sums[2], 100)

	testRemovingAmount(t, testInventory, testItems[1], &sums[1], 10)
	testRemovingAmount(t, testInventory, testItems[1], &sums[1], 10)
	testRemovingAmount(t, testInventory, testItems[1], &sums[1], 10)

	testRemovingAmount(t, testInventory, testItems[0], &sums[0], 1)
	testRemovingAmount(t, testInventory, testItems[0], &sums[0], 1)
	testRemovingAmount(t, testInventory, testItems[0], &sums[0], 1)
}

func testAddingAmount(t *testing.T, testInventory Inventory, testItem item.Item, sum *Amount, amountToAdd Amount) {
	testInventory.AddStock(testItem, amountToAdd)
	*sum = *sum + amountToAdd
	retrievedAmount := testInventory.GetStock(testItem)
	if *sum != retrievedAmount {
		t.Fatalf("invalid amount, expected %d, retrieved %d\n", *sum, retrievedAmount)
	}
	return
}

func testRemovingAmount(t *testing.T, testInventory Inventory, testItem item.Item, sum *Amount, amountToRemove Amount) {
	err := testInventory.RemoveStock(testItem, amountToRemove)
	if err != nil {
		t.Fatalf("could not remove stock, error thrown: %s", err.Error())
	}
	*sum = *sum - amountToRemove
	retrievedAmount := testInventory.GetStock(testItem)
	if *sum != retrievedAmount {
		t.Fatalf("invalid amount, expected %d, retrieved %d\n", *sum, retrievedAmount)
	}
	return
}

func makeTestItem(id item.Id) item.Item {
	var testItem item.Item = item.Item{
		Id: id,
	}
	return testItem
}
