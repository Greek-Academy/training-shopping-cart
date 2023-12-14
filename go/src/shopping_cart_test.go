package shoppingcart

import (
	"fmt"
	"testing"
)

func TestAddRemove(t *testing.T) {
	testCases := []struct {
		code     string
		add      int
		remove   int
		expected int
	}{
		{"apple", 2, 1, 1},
		{"orange", 3, 3, 0},
		{"grape", 10, 5, 5},
	}

	for _, testCase := range testCases {
		shoppingCart := ShoppingCart{Products: make(map[string]Stock)}

		t.Run(
			fmt.Sprintf("Add Test, code: %s, add: %d", testCase.code, testCase.add),
			func(t *testing.T) {
				actual := shoppingCart.Add(testCase.code, testCase.add)
				if testCase.add != actual {
					t.Errorf("Expected %d, but got %d", testCase.expected, actual)
				}
			})

		t.Run(
			fmt.Sprintf("Remove Test, code: %s, remove: %d", testCase.code, testCase.add),
			func(t *testing.T) {
				actual := shoppingCart.Remove(testCase.code, testCase.remove)
				if testCase.expected != actual {
					t.Errorf("Expected %d, but got %d", testCase.expected, actual)
				}
			})
	}
}

func TestOrder(t *testing.T) {

	type P struct {
		fruit string
		count int
	}
	type E struct {
		Count int
		Total int
	}
	testCases := []struct {
		products []P
		expected E
	}{
		{
			products: []P{{"apple", 3}, {"apple", 3}, {"apple", 3}},
			expected: E{9, 450},
		},
		{
			products: []P{{"apple", 20}, {"grape", 4}, {"orange", 8}},
			expected: E{32, 1960},
		},
		{
			products: []P{{"grape", 3}, {"orange", 9}},
			expected: E{12, 930},
		},
	}

	for _, testCase := range testCases {
		shoppingCart := ShoppingCart{Products: make(map[string]Stock)}

		t.Run(
			fmt.Sprintf("Order Test, products: %v", testCase.products),
			func(t *testing.T) {
				for _, product := range testCase.products {
					shoppingCart.Add(product.fruit, product.count)
				}

				actual, err := shoppingCart.Order()
				if err != nil {
					t.Errorf("failed: %s", err.Error())
					return
				}

				expected := Order(testCase.expected)
				if expected != *actual {
					t.Errorf("Expected %v, but got %v", testCase.expected, *actual)
				}
			})
	}
}
