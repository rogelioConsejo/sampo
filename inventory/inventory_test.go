package inventory

import (
	"math/rand"
	"testing"
)

func TestInventory(t *testing.T) {
	t.Parallel()
	var testInventory Inventory = New()
	var testItems []Item = []Item{
		makeTestItem("test-item1"),
		makeTestItem("test-item2"),
		makeTestItem("test-item3"),
	}

	var sums [3]Amount
	testAddingAmount(t, testInventory, testItems[0], &sums[0], 10)
	testAddingAmount(t, testInventory, testItems[0], &sums[0], 20)
	testAddingAmount(t, testInventory, testItems[0], &sums[0], 30)
	testAddingAmount(t, testInventory, testItems[0], &sums[0], 10.1)
	testAddingAmount(t, testInventory, testItems[0], &sums[0], 1)
	testAddingAmount(t, testInventory, testItems[0], &sums[0], 11.2)
	testAddingAmount(t, testInventory, testItems[0], &sums[0], 0.01)

	testAddingAmount(t, testInventory, testItems[1], &sums[1], 10)
	testAddingAmount(t, testInventory, testItems[1], &sums[1], 10)
	testAddingAmount(t, testInventory, testItems[1], &sums[1], 10)

	testAddingAmount(t, testInventory, testItems[2], &sums[2], .01)
	testAddingAmount(t, testInventory, testItems[2], &sums[2], .1)
	testAddingAmount(t, testInventory, testItems[2], &sums[2], 1)
}

func TestInventory_AddStock_ThrowsErrorOnNegativeAmount(t *testing.T) {
	t.Parallel()
	var testInventory Inventory = New()
	testItem := makeTestItem("test-item")

	err := testInventory.AddStock(testItem, Amount(-1000*rand.Float64()-1))
	_ = testInventory.GetStock(testItem)
	if isEmpty(err) {
		t.Fatal("error not thrown on negative amount")
	}
	if err != negativeAmountError {
		t.Fatalf("invalid error, got %s, expected %s\n", err.Error(), negativeAmountError)
	}
}

func testAddingAmount(t *testing.T, testInventory Inventory, testItem Item, sum *Amount, amountToAdd Amount) {
	err := testInventory.AddStock(testItem, amountToAdd)
	if isNotEmpty(err) {
		t.Fatalf("error adding stock: %s", err.Error())
	}
	expectedAmount := *sum + amountToAdd
	retrievedAmount := testInventory.GetStock(testItem)
	if retrievedAmount != expectedAmount {
		t.Fatalf("invalid amount, expected %f, retrieved %f\n", *sum, retrievedAmount)
	}
	*sum = expectedAmount
	return
}

func makeTestItem(name ItemName) Item {
	var testItem Item = Item{
		name,
	}
	return testItem
}

func isEmpty(err error) bool {
	return err == nil
}
func isNotEmpty(err error) bool {
	return err != nil
}
