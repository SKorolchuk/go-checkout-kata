package checkout

// PricePerUnit represents price per specific amount
type PricePerUnit struct {
	Units int32 `json:"units"`
	Price int32 `json:"price"`
}

// Sku represents collection of prices for SKU item
type Sku struct {
	SkuName       string         `json:"name"`
	PricesPerUnit []PricePerUnit `json:"pricesPerUnit"`
}

// CheckoutContext represents context for Checkout processing
type CheckoutContext struct {
	SKUs        []Sku  `json:"skus"`
	scanHistory []rune `json:"-"`
}
