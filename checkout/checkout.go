// Package checkout for test task
package checkout

import (
	"errors"
	"fmt"
)

// Checkout is API for Checkout process
type Checkout interface {
	Scan(item string) error
	GetTotalPrices() (int32, error)
}

func (ctx *CheckoutContext) Scan(item string) error {
	for _, skuInput := range item {
		if err := ctx.addSkuToCheckout(skuInput); err != nil {
			return err
		}
	}

	return nil
}

func (ctx *CheckoutContext) GetTotalPrices() (int32, error) {
	result := int32(0)
	skuTotal := *ctx.getSkuTotal()

	if len(skuTotal) == 0 {
		return result, nil
	}

	for skuName, checkoutItemsCount := range skuTotal {
		sku := ctx.findSku(skuName)

		if sku == nil {
			return 0, fmt.Errorf("SKU=%s not found", string(skuName))
		}

		if len(sku.PricesPerUnit) > 0 && checkoutItemsCount > 0 {
			reminder := checkoutItemsCount

			for reminder > 0 {
				minPrice := sku.findMinPriceForReminder(reminder)

				if minPrice.Price == 0 {
					return 0, fmt.Errorf("SKU=%s prices for %d amount can not be processed", string(skuName), reminder)
				}

				result += reminder / minPrice.Units * minPrice.Price
				reminder = reminder % minPrice.Units
			}
		}
	}

	return result, nil
}

func (ctx *CheckoutContext) addSkuToCheckout(skuInput rune) error {
	for _, sku := range ctx.SKUs {
		if rune(sku.SkuName[0]) == skuInput {
			ctx.scanHistory = append(ctx.scanHistory, skuInput)

			return nil
		}
	}

	return errors.New("SKU item not found")
}

func (ctx *CheckoutContext) getSkuTotal() *map[rune]int32 {
	skuTotal := make(map[rune]int32)

	for _, item := range ctx.scanHistory {
		if _, ok := skuTotal[item]; !ok {
			skuTotal[item] = 1
		} else {
			skuTotal[item] += 1
		}
	}

	return &skuTotal
}

func (ctx *CheckoutContext) findSku(skuName rune) *Sku {
	var sku *Sku

	for i := 0; sku == nil || i < len(ctx.SKUs); i++ {
		if rune(ctx.SKUs[i].SkuName[0]) == skuName {
			sku = &ctx.SKUs[i]
		}
	}

	return sku
}

func (sku *Sku) findMinPriceForReminder(reminder int32) PricePerUnit {
	var minPrice PricePerUnit

	for _, price := range sku.PricesPerUnit {
		isPriceLower := minPrice.Price != 0 && float32(minPrice.Price)/float32(minPrice.Units) > float32(price.Price)/float32(price.Units)
		isEnoughUnits := reminder-price.Units >= 0

		if isEnoughUnits && (minPrice.Price == 0 || isPriceLower) {
			minPrice = price
		}
	}

	return minPrice
}
