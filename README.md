# Checkout Kata Solution in Go

## Overview

Implement the code for a checkout system that handles pricing schemes such as _**"pineapples cost 50, three pineapples cost 130."**_

Implement the code for a supermarket checkout that calculates the total price of a number of items.

In a normal supermarket, things are identified using Stock Keeping Units, or SKUs. In our store, we’ll use individual letters of the alphabet (A, B, C, and so on). Our goods are priced individually.

In addition, some items are multi-priced: buy n of them, and they’ll cost you y. For example, item **A** might cost 50 individually, but this week we have a special offer: buy three **A**s and they’ll cost you 130.

In fact the prices are:

|SKU|Unit Price|Special Price|
|---|---|---|
|A|50|3 for 130|
|B|30|2 for 45|
|C|20||
|D|15||

The checkout accepts items in any order, so that if we scan a B, an A, and another B, we’ll recognize the two Bs and price them at 45 (for a total price so far of 95).

The pricing changes frequently, so pricing should be independent of the checkout.pricing changes frequently, so pricing should be independent of the checkout. The interface to the checkout could look like:
```csharp
interface ICheckout
{
    void Scan(string item);
    int GetTotalPrice();
}
```

## Structure

|Folder|Description|
|---|---|
|cmd/checkout-cli|Sample CLI implementation of Checkout process|
|internal/pkg/checkout|Folder contains package with solution of Checkout API|
|internal/pkg/history|Folder contains API for managing scan history|
|internal/pkg/sku|Folder contains SKU catalog models and methods|
|test|SKU catalog samples which used in unit tests and CLI default configuration|
|bin (after `make build`)|Location for executable `checkout-cli` binary file|

## Commands

|Command|Description|
|---|---|
|make fmt|Command will re-format Go files|
|make vet|Command will lint Go files|
|make clean|Command will clean Go package cache and `bin` folder|
|make build|Command will clean up, format and lint code and then build `checkout-cli` binary|
|make run|Command will run `checkout-cli` application|
|make tests|Command will run all unit tests|
|make lint|Command will run staticcheck tool|
|make tidy|Command will sync Go modules|
|make deps-cleancache|Command will clean up Go modules cache|
|make list|Command will list Go modules|

## How To

### Hot to run

```bash
~ make run-example
```

```bash
~ make build
~ ./bin/checkout-cli --scan-series AABCDBBA
```

NOTE: `--scan-series` argument accepts series of SKU names presented in SKU catalog file
