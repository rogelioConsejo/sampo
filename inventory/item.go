package inventory

type Item struct {
	ItemName `json:"name"`
}

type ItemName string
