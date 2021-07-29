package sku

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type testCase struct {
	data           SKU
	units          int32
	expectedResult int32
}

func TestGetOptimalCheckoutPrice(t *testing.T) {
	nameFormat := func(testCase *testCase) string {
		return fmt.Sprintf("Optimal price should be %d for %d units of %s SKU", testCase.expectedResult, testCase.units, testCase.data.Name)
	}

	testCases := []testCase{
		{
			data: SKU{
				Name: "A",
				Prices: []PricePerUnit{
					{Price: 10, Units: 1},
					{Price: 15, Units: 2},
				},
			},
			units:          3,
			expectedResult: 25,
		},
		{
			data: SKU{
				Name: "B",
				Prices: []PricePerUnit{
					{Price: 75, Units: 1},
					{Price: 200, Units: 4},
				},
			},
			units:          7,
			expectedResult: 425,
		},
		{
			data: SKU{
				Name: "C",
				Prices: []PricePerUnit{
					{Price: 25, Units: 1},
				},
			},
			units:          5,
			expectedResult: 125,
		},
		{
			data: SKU{
				Name: "D",
				Prices: []PricePerUnit{
					{Price: 7, Units: 1},
					{Price: 10, Units: 2},
					{Price: 27, Units: 5},
				},
			},
			units:          15,
			expectedResult: 77,
		},
	}

	for _, testCase := range testCases {
		t.Run(nameFormat(&testCase), func(t *testing.T) {
			actualResult, err := testCase.data.GetOptimalCheckoutPrice(testCase.units)

			assert.NoError(t, err)
			assert.Equal(t, testCase.expectedResult, actualResult)
		})
	}

	t.Run("Optimal price should be 0", func(t *testing.T) {
		t.Run("for emtpy SKU", func(t *testing.T) {
			emptySKU := SKU{Name: "A", Prices: nil}

			actualResult, err := emptySKU.GetOptimalCheckoutPrice(1)

			assert.NoError(t, err)
			assert.Equal(t, int32(0), actualResult)
		})

		t.Run("for zero units", func(t *testing.T) {
			testSKU := SKU{Name: "A", Prices: []PricePerUnit{{Price: 1, Units: 1}}}

			actualResult, err := testSKU.GetOptimalCheckoutPrice(0)

			assert.NoError(t, err)
			assert.Equal(t, int32(0), actualResult)
		})
	})

	t.Run("Returns error when", func(t *testing.T) {
		t.Run("SKU Price is 0", func(t *testing.T) {
			emptySKU := SKU{
				Name: "A",
				Prices: []PricePerUnit{
					{Price: 0, Units: 0},
				},
			}

			actualResult, err := emptySKU.GetOptimalCheckoutPrice(1)

			assert.EqualError(t, err, "invalid PricePerUnit")
			assert.Equal(t, int32(0), actualResult)
		})
	})
}
