package shoppingcart

import "errors"

type Stock struct {
	Count   int
	Product *Product
}

type Order struct {
	Count int
	Total int
}

type ShoppingCart struct {
	Products map[string]Stock
}

func (shoppingCart *ShoppingCart) Add(code string, count int) int {
	v, exists := shoppingCart.Products[code]
	if exists {
		v.Count += count
		shoppingCart.Products[code] = v
		return v.Count
	} else {
		shoppingCart.Products[code] = Stock{Count: count, Product: &Product{code: code}}
		return count
	}
}

func (shoppingCart *ShoppingCart) Remove(code string, count int) int {
	v, exists := shoppingCart.Products[code]
	if exists {
		v.Count -= count
		shoppingCart.Products[code] = v
		// 0個になったらmapから削除
		if v.Count == 0 {
			delete(shoppingCart.Products, code)
			return 0
		}
		return v.Count
	} else {
		return 0
	}
}

func (shoppingCart *ShoppingCart) Order() (*Order, error) {
	if len(shoppingCart.Products) == 0 {
		return nil, errors.New("カートが空なので注文できません")
	}

	count := 0
	total := 0
	for _, stock := range shoppingCart.Products {
		price, err := stock.Product.Price()
		if err != nil {
			return nil, err
		}
		count += stock.Count
		total += stock.Count * price
	}

	return &Order{
		Count: count,
		Total: total,
	}, nil
}
