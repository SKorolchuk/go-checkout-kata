package checkout

import (
	"github.com/SKorolchuk/go-checkout-kata/internal/pkg/processor"
)

// ICheckoutService is API for Checkout process
type ICheckoutService interface {
	Scan(item string) error
	GetTotalPrices() (int32, error)
}

type service struct {
}

var Instance ICheckoutService = &service{}

func (ctx *service) Scan(item string) error {
	for _, skuInput := range item {
		if err := processor.Instance.AddSKUToCheckout(skuInput); err != nil {
			return err
		}
	}

	return nil
}

func (ctx *service) GetTotalPrices() (int32, error) {
	return processor.Instance.GetTotalPrices()
}
