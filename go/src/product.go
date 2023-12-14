package shoppingcart

import "errors"

type Product struct {
	code string
}

func (p *Product) Name() (string, error) {
	switch p.code {
	case "apple":
		return "りんご", nil
	case "grape":
		return "ぶどう", nil
	case "orange":
		return "みかん", nil
	}
	return "", errors.New("code値が不正です")
}

func (p *Product) Price() (int, error) {
	switch p.code {
	case "apple":
		return 50, nil
	case "grape":
		return 100, nil
	case "orange":
		return 70, nil
	}
	return 0, errors.New("code値が不正です")
}
