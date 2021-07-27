package checkout_test

import (
	"encoding/json"
	"fmt"
	"github.com/SKorolchuk/go-checkout-kata/internal/pkg/checkout"
	"github.com/SKorolchuk/go-checkout-kata/internal/pkg/processor"
	"github.com/SKorolchuk/go-checkout-kata/internal/pkg/sku"
	"io/ioutil"
	"os"
	"testing"
)

func TestCheckoutFlow(t *testing.T) {
	testCases := []struct {
		scanInput      string
		expectedResult int32
	}{
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
		testCaseName := fmt.Sprintf("Checkout should return total price equal to %d for %s input", testCase.expectedResult, testCase.scanInput)

		t.Run(testCaseName, func(t *testing.T) {
			catalog := getTestCatalog(t)
			if catalog == nil {
				t.Error("Failed to read test json file")
			}

			if err := processor.Instance.CleanCheckoutHistory(); err != nil {
				t.Error(err)
			}
			if err := processor.Instance.SetSKUCatalog(catalog); err != nil {
				t.Error(err)
			}
			if err := checkout.Instance.Scan(testCase.scanInput); err != nil {
				t.Error(err)
			}

			actualResult, err := checkout.Instance.GetTotalPrices()
			if err != nil {
				t.Error(err)
			}

			if testCase.expectedResult != actualResult {
				t.Errorf("%d != %d", testCase.expectedResult, actualResult)
			}
		})
	}
}

func TestCheckoutError(t *testing.T) {
	catalog := getTestCatalog(t)
	if catalog == nil {
		t.Error("Failed to read test json file")
	}

	t.Run("Checkout should throw error for incorrect SKU names", func(t *testing.T) {
		if err := processor.Instance.CleanCheckoutHistory(); err != nil {
			t.Error(err)
		}
		if err := processor.Instance.SetSKUCatalog(catalog); err != nil {
			t.Error(err)
		}
		if err := checkout.Instance.Scan("%"); err == nil {
			t.Error("Scan method should check SKU name")
		}
	})
}

func getTestCatalog(t *testing.T) *sku.SKUCollection {
	file, err := os.Open("../../../test/.checkout_test_data/test_skus.json")
	defer func() {
		if err := file.Close(); err != nil {
			t.Error(err)
		}
	}()

	if err != nil {
		t.Error(err)
	}

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		t.Error(err)
	}

	var catalog sku.SKUCollection

	if err := json.Unmarshal(bytes, &catalog); err != nil {
		t.Error(err)
	}

	return &catalog
}
