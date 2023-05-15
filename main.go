package main

import "fmt"

func main() {

}

type Data struct {
	Price    float64
	Quantity int
}

func GetItem(items []string, amount float64, itemList map[string]*Data) string {
	var total float64

	for _, item := range items {
		if itemList[item].Quantity < 1 {
			return fmt.Sprintf("Item %v is out of stock", item)
		}
		total += itemList[item].Price
		itemList[item].Quantity = itemList[item].Quantity - 1
	}

	if amount < total {
		return "Not enough money"
	}

	remainder := amount - total
	totalPrice := fmt.Sprintf("%.2f", remainder)
	return totalPrice
}
