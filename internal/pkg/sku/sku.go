package sku

import (
	"errors"
)

// SKUCollection represents SKU catalog
type SKUCollection struct {
	SKUs []SKU `json:"skus"`
}

// SKU represents collection of PricePerUnit for an item
type SKU struct {
	Name   string         `json:"name"`
	Prices []PricePerUnit `json:"prices"`
}

// PricePerUnit represents price per specific amount
type PricePerUnit struct {
	Units int32 `json:"units"`
	Price int32 `json:"price"`
}

// FindMinAvailablePrice search the best price for available SKU units
func (sku *SKU) FindMinAvailablePrice(units int32) (*PricePerUnit, error) {
	if len(sku.Prices) <= 0 {
		return nil, errors.New("prices collection is empty")
	}

	minPrice := sku.Prices[0]

	for _, price := range sku.Prices {
		isPriceLower := minPrice.Price/minPrice.Units > price.Price/price.Units
		isEnoughUnits := units-price.Units >= 0

		if isEnoughUnits && isPriceLower {
			minPrice = price
		}
	}

	return &minPrice, nil
}
