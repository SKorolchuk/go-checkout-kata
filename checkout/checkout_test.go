package checkout_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"tasks/checkout"
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
		t.Run(fmt.Sprintf("Checkout should return total price equal to %d for %s input", testCase.expectedResult, testCase.scanInput), func(t *testing.T) {
			context := getTestContext(t)
			if context == nil {
				t.Error("Failed to read test json file")
			}

			if err := context.Scan(testCase.scanInput); err != nil {
				t.Error(err)
			}

			actualResult, err := context.GetTotalPrices()
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
	context := getTestContext(t)
	if context == nil {
		t.Error("Failed to read test json file")
	}

	t.Run("Checkout should throw error for incorrect SKU names", func(t *testing.T) {
		if err := context.Scan("%"); err == nil {
			t.Error("Scan method should check SKU name")
		}
	})
}

func getTestContext(t *testing.T) *checkout.CheckoutContext {
	file, err := os.Open("test_skus.json")
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

	var context checkout.CheckoutContext

	json.Unmarshal([]byte(bytes), &context)

	return &context
}
