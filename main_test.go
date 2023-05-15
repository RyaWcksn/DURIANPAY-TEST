package main

import "testing"

func TestGetItem(t *testing.T) {

	allStock := map[string]*Data{
		"B": {
			Price:    0.75,
			Quantity: 48,
		},
		"M": {
			Price:    1.00,
			Quantity: 36,
		},
		"C": {
			Price:    1.35,
			Quantity: 24,
		},
		"W": {
			Price:    1.50,
			Quantity: 30,
		},
	}

	emptyStock := map[string]*Data{
		"B": {
			Price:    0.75,
			Quantity: 0,
		},
		"M": {
			Price:    1.00,
			Quantity: 36,
		},
		"C": {
			Price:    1.35,
			Quantity: 24,
		},
		"W": {
			Price:    1.50,
			Quantity: 30,
		},
	}

	type req struct {
		items    []string
		amount   float64
		itemList map[string]*Data
	}

	tests := []struct {
		name    string
		request req
		want    string
	}{
		{
			name: "Success",
			request: req{
				items:    []string{"B", "M"},
				amount:   2.000,
				itemList: allStock,
			},
			want: "0.25",
		},
		{
			name: "No stock",
			request: req{
				items:    []string{"B"},
				amount:   2.000,
				itemList: emptyStock,
			},
			want: "Item B is out of stock",
		},
		{
			name: "Not enough money",
			request: req{
				items:    []string{"B", "M"},
				amount:   1.00,
				itemList: allStock,
			},
			want: "Not enough money",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ans := GetItem(test.request.items, float64(test.request.amount), test.request.itemList)
			if ans != test.want {
				t.Errorf("got %s, want %s", ans, test.want)
			}
		})
	}
}
