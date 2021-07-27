package processor

import (
	"errors"
	"fmt"
	"github.com/SKorolchuk/go-checkout-kata/internal/pkg/sku"
	"sync"
)

type IProcessor interface {
	SetSKUCatalog(catalog *sku.SKUCollection) error
	AddSKUToCheckout(skuName rune) error
	GetTotalPrices() (int32, error)
	CleanCheckoutHistory() error
}

type processorContext struct {
	catalog     *sku.SKUCollection
	scanHistory []rune
}

var Instance IProcessor = &processorContext{
	catalog:     nil,
	scanHistory: make([]rune, 0),
}

var mutex = &sync.Mutex{}

func (ctx *processorContext) SetSKUCatalog(catalog *sku.SKUCollection) error {
	if catalog == nil || catalog.SKUs == nil {
		return errors.New("catalog is not specified")
	}

	mutex.Lock()
	ctx.catalog = catalog
	mutex.Unlock()

	return nil
}

func (ctx *processorContext) AddSKUToCheckout(SKUName rune) error {
	for _, currentSKU := range ctx.catalog.SKUs {
		if rune(currentSKU.Name[0]) == SKUName {
			mutex.Lock()
			ctx.scanHistory = append(ctx.scanHistory, SKUName)
			mutex.Unlock()

			return nil
		}
	}

	return errors.New("SKU item not found")
}

func (ctx *processorContext) GetTotalPrices() (int32, error) {
	result := int32(0)
	total := ctx.getTotalUnitsPerSKU()

	if len(total) == 0 {
		return result, nil
	}

	for name, count := range total {
		SKU := ctx.getSKUbyName(name)

		if SKU == nil {
			return 0, fmt.Errorf("SKU=%s not found", string(name))
		}

		if len(SKU.Prices) > 0 && count > 0 {
			reminder := count

			for reminder > 0 {
				minPrice, err := SKU.FindMinAvailablePrice(reminder)
				if err != nil {
					return 0, err
				}

				if minPrice == nil {
					return 0, fmt.Errorf("SKU=%s prices for %d amount can not be processed", string(name), reminder)
				}

				result += reminder / minPrice.Units * minPrice.Price
				reminder = reminder % minPrice.Units
			}
		}
	}

	return result, nil
}

func (ctx *processorContext) CleanCheckoutHistory() error {
	mutex.Lock()
	ctx.scanHistory = make([]rune, 0)
	mutex.Unlock()

	return nil
}

func (ctx *processorContext) getTotalUnitsPerSKU() map[rune]int32 {
	total := make(map[rune]int32)

	for _, item := range ctx.scanHistory {
		if _, ok := total[item]; !ok {
			total[item] = 1
		} else {
			total[item] += 1
		}
	}

	return total
}

func (ctx *processorContext) getSKUbyName(skuName rune) *sku.SKU {
	SKUs := ctx.catalog.SKUs

	for i := 0; i < len(SKUs); i++ {
		if rune(SKUs[i].Name[0]) == skuName {
			return &SKUs[i]
		}
	}

	return nil
}
