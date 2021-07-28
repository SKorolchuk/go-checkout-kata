package sku

import (
	"errors"
	"fmt"
)

// Catalog represents collection of SKU
type Catalog struct {
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

// GetOptimalCheckoutPrice calculate total checkout price
func (SKU *SKU) GetOptimalCheckoutPrice(unitsToCheckout int32) (int32, error) {
	result := int32(0)

	if len(SKU.Prices) > 0 && unitsToCheckout > 0 {
		reminder := unitsToCheckout

		for reminder > 0 {
			minPrice, err := SKU.findMinAvailablePrice(reminder)
			if err != nil {
				return 0, err
			}

			if minPrice == nil {
				return 0, fmt.Errorf("SKU=%s prices for %d amount can not be processed", SKU.Name, reminder)
			}

			result += reminder / minPrice.Units * minPrice.Price
			reminder = reminder % minPrice.Units
		}
	}

	return result, nil
}

// GetSKUbyName finds SKU by name from Catalog
func (catalog *Catalog) GetSKUbyName(skuName rune) *SKU {
	SKUs := catalog.SKUs

	for i := 0; i < len(SKUs); i++ {
		if rune(SKUs[i].Name[0]) == skuName {
			return &SKUs[i]
		}
	}

	return nil
}

func (SKU *SKU) findMinAvailablePrice(units int32) (*PricePerUnit, error) {
	if len(SKU.Prices) <= 0 {
		return nil, errors.New("prices collection is empty")
	}

	minPrice := SKU.Prices[0]

	for _, price := range SKU.Prices {
		isPriceLower := minPrice.Price/minPrice.Units > price.Price/price.Units
		isEnoughUnits := units-price.Units >= 0

		if isEnoughUnits && isPriceLower {
			minPrice = price
		}
	}

	return &minPrice, nil
}
