package history

import (
	"errors"
)

// ScanHistory represents checkout scan process
type ScanHistory interface {
	Add(skuName rune) error
	GetTotalUnitsPerSKU() map[rune]int32
	Clean() error
}

type context struct {
	scanHistory []rune
}

// NewScanHistory is used to construct ScanHistory instance
func NewScanHistory() ScanHistory {
	ctx := context{}
	ctx.scanHistory = make([]rune, 0)

	return &ctx
}

// Add is used to scan SKU item
func (ctx *context) Add(SKUName rune) error {
	if ctx.scanHistory == nil {
		return errors.New("history is not initialized")
	}

	ctx.scanHistory = append(ctx.scanHistory, SKUName)
	return nil
}

// Clean is used to purge scan history
func (ctx *context) Clean() error {
	ctx.scanHistory = make([]rune, 0)

	return nil
}

// GetTotalUnitsPerSKU returns dictionary of scanned units per SKU
func (ctx *context) GetTotalUnitsPerSKU() map[rune]int32 {
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
