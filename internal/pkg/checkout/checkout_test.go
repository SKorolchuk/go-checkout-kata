package checkout_test

import (
	"encoding/json"
	"fmt"
	"github.com/SKorolchuk/go-checkout-kata/internal/pkg/checkout"
	"github.com/SKorolchuk/go-checkout-kata/internal/pkg/sku"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
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
			catalog := getTestCatalog(t)
			if catalog == nil {
				t.Error("Failed to read test json file")
			}

			underTest, err := checkout.New(catalog)
			if err != nil {
				t.Error(err)
			}

			if err := underTest.Scan(testCase.scanInput); err != nil {
				t.Error(err)
			}

			actualResult, err := underTest.GetTotalPrices()
			if err != nil {
				t.Error(err)
			}

			assert.Equal(t, testCase.expectedResult, actualResult)
		})
	}
}

func TestCheckoutError(t *testing.T) {
	catalog := getTestCatalog(t)
	if catalog == nil {
		t.Error("Failed to read test json file")
	}

	underTest, err := checkout.New(catalog)
	if err != nil {
		t.Error(err)
	}

	t.Run("Checkout should throw error for incorrect SKU names", func(t *testing.T) {
		err := underTest.Scan("%")
		if assert.Error(t, err) {
			assert.EqualError(t, err, "SKU not found")
		}
	})
}

func getTestCatalog(t *testing.T) *sku.Catalog {
	file, err := os.Open("../../../test/.checkout_test_data/skus.json")
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

	var catalog sku.Catalog

	if err := json.Unmarshal(bytes, &catalog); err != nil {
		t.Error(err)
	}

	return &catalog
}
