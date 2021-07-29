// Package main for the test task
package main

import (
	"fmt"
	"github.com/SKorolchuk/go-checkout-kata/internal/pkg/checkout"
	"github.com/SKorolchuk/go-checkout-kata/internal/pkg/sku"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

const (
	itemsToScanFlagName              = "scan-series"
	SKUCatalogSourceLocationFlagName = "sku-source-file"
)

func main() {
	app := &cli.App{
		Name:  "checkout-cli",
		Usage: "Sample CLI implementation of Checkout process",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     itemsToScanFlagName,
				Value:    "",
				Usage:    "Specify list of SKU items to scan (for ex., AABCA)",
				Required: true,
			},
			&cli.StringFlag{
				Name:  SKUCatalogSourceLocationFlagName,
				Value: "test/.checkout_cli_data/skus.json",
				Usage: "Specify path to SKU json file",
			},
		},
		Action: checkoutCLIHandler,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func checkoutCLIHandler(c *cli.Context) error {
	SKUCatalogSourceLocation := c.String(SKUCatalogSourceLocationFlagName)
	fmt.Printf("SKU Catalog location: %s\r\n", SKUCatalogSourceLocation)

	itemsToScan := c.String(itemsToScanFlagName)
	fmt.Printf("Items to scan: %s\r\n", itemsToScan)

	var catalog sku.Catalog
	err := catalog.Load(SKUCatalogSourceLocation)
	if err != nil {
		return err
	}

	checkoutProcess, err := checkout.NewCheckout(&catalog)
	if err != nil {
		return err
	}

	if err := checkoutProcess.Scan(itemsToScan); err != nil {
		return err
	}

	totalPrice, err := checkoutProcess.GetTotalPrices()
	if err != nil {
		return err
	}

	fmt.Printf("Total Price: %d\r\n", totalPrice)

	return nil
}
