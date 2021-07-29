package checkout

import (
	"errors"
	"fmt"
	"github.com/SKorolchuk/go-checkout-kata/internal/pkg/history"
	"github.com/SKorolchuk/go-checkout-kata/internal/pkg/sku"
)

// Checkout is API for Checkout process
type Checkout interface {
	Scan(item string) error
	GetTotalPrices() (int32, error)
}

type context struct {
	catalog     *sku.Catalog
	scanHistory history.ScanHistory
}

func NewCheckout(catalog *sku.Catalog) (Checkout, error) {
	if catalog == nil || catalog.SKUs == nil {
		return nil, errors.New("catalog is not specified")
	}

	scanHistory := history.NewScanHistory()

	ctx := context{
		catalog:     catalog,
		scanHistory: scanHistory,
	}

	return &ctx, nil
}

func (ctx *context) Scan(item string) error {
	for _, skuInput := range item {
		if err := ctx.tryAddSKUToCheckout(skuInput); err != nil {
			return err
		}
	}

	return nil
}

func (ctx *context) GetTotalPrices() (int32, error) {
	result := int32(0)
	unitsPerSKU := ctx.scanHistory.GetTotalUnitsPerSKU()

	if len(unitsPerSKU) == 0 {
		return result, nil
	}

	for SKUName, count := range unitsPerSKU {
		SKU := ctx.catalog.GetSKUbyName(SKUName)

		if SKU == nil {
			return 0, fmt.Errorf("SKU=%s not found", string(SKUName))
		}

		totalPerSKU, err := SKU.GetOptimalCheckoutPrice(count)
		if err != nil {
			return 0, err
		}

		result += totalPerSKU
	}

	return result, nil
}

func (ctx *context) tryAddSKUToCheckout(skuInput rune) error {
	for _, currentSKU := range ctx.catalog.SKUs {
		if rune(currentSKU.Name[0]) == skuInput {
			if err := ctx.scanHistory.Add(skuInput); err != nil {
				return err
			} else {
				return nil
			}
		}
	}

	return errors.New("SKU not found")
}
