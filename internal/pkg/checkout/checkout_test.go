package checkout_test

import (
	"fmt"
	"github.com/SKorolchuk/go-checkout-kata/internal/pkg/checkout"
	"github.com/SKorolchuk/go-checkout-kata/internal/pkg/sku"
	"github.com/stretchr/testify/assert"
	"testing"
)

type testCase struct {
	scanInput      string
	expectedResult int32
}

func TestCheckoutFlow(t *testing.T) {
	nameFormat := func(testCase *testCase) string {
		return fmt.Sprintf("Checkout should return total price equal to %d for %s input", testCase.expectedResult, testCase.scanInput)
	}

	testCases := []testCase{
		{
			scanInput:      "AABBCAA",
			expectedResult: 245,
		},
		{
			scanInput:      "A",
			expectedResult: 50,
		},
		{
			scanInput:      "B",
			expectedResult: 30,
		},
		{
			scanInput:      "CBA",
			expectedResult: 100,
		},
		{
			scanInput:      "",
			expectedResult: 0,
		},
	}

	for _, testCase := range testCases {
		t.Run(nameFormat(&testCase), func(t *testing.T) {
			catalog, err := getTestCatalog()
			assert.NoError(t, err)
			if catalog == nil {
				t.Error("Failed to read test json file")
			}

			underTest, err := checkout.NewCheckout(catalog)
			assert.NoError(t, err)

			err = underTest.Scan(testCase.scanInput)
			assert.NoError(t, err)

			actualResult, err := underTest.GetTotalPrices()

			assert.NoError(t, err)
			assert.Equal(t, testCase.expectedResult, actualResult)
		})
	}
}

func TestCheckoutError(t *testing.T) {
	catalog, err := getTestCatalog()
	assert.NoError(t, err)

	if catalog == nil {
		t.Error("Failed to read test json file")
	}

	underTest, err := checkout.NewCheckout(catalog)
	assert.NoError(t, err)

	t.Run("Checkout should throw error for incorrect SKU names", func(t *testing.T) {
		assert.EqualError(t, underTest.Scan("%"), "SKU not found")
	})
}

func getTestCatalog() (*sku.Catalog, error) {
	var catalog sku.Catalog

	if err := catalog.Load("../../../test/.checkout_test_data/skus.json"); err != nil {
		return nil, err
	}

	return &catalog, nil
}
