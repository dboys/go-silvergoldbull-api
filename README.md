# go-silvergoldbull-api
Golang client for Silver Gold Bull web service

# NAME

go-silvergoldbull-api - Golang client for Silver Gold Bull web service (https://silvergoldbull.com/)

# INSTALLATION

To install this module, run the following commands:

	$ go get github.com/dboys/go-silvergoldbull-api

# TESTING

To test this module, run the following commands:

	$ export SILVERGOLDBULL_API_KEY=<API key>
	$ go test -v

# SYNOPSIS
```go
import sgb "github.com/dboys/go-silvergoldbull-api"

func main() {
	client := sgb.New("<API key>")

	// get available currencies
	curr, err := client.GetCurrencyList()
	if err != nil {
	 	...
	}

	// get available payment methods
	pm, err := client.GetPaymentList()
	if err != nil {
	 	...
	}

	// get available shipping methods
	ship, err := client.GetShippingList()
	if err != nil {
	 	...
	}

	// get available products
	p, err := client.GetProductList()
	if err != nil {
	 	...
	}

	// get product by id
	p, err := client.GetProduct(<product id>)
	if err != nil {
	 	...
	}

	// get available orders
	o, err := client.GetOrderList()
	if err != nil {
	 	...
	}

	// get order by id
	p, err := client.GetOrder(<order id>)
	if err != nil {
	 	...
	}

	//create order
	addr := &sgb.Address{
		Email: "sales@silvergoldbull.com",
		Postcode: "T2P 	5C5",
		Region:    "AB",
		City:      "Calgary",
		FirstName: "John",
		LastName:  "Smith",
		Country:   "CA",
		Street:    "888 - 3 ST SW, 10 FLOOR - WEST TOWER",
		Phone:     "+1 (403) 668 8648",
	}

	items := []*sgb.Item{
		{
			ID:       "2706",
			QTY:      1,
			BidPrice: 468.37,
		},
		{
			ID:       "2580",
			QTY:      1,
			BidPrice: 2,
		},
	}

	order := &sgb.Order{
		Currency:      "USD",
		PaymentMethod: "paypall",
		ShipMethod:    "1YR_STORAGE",
		Declaration:   "TEST",
		Items:         items,
		Shipping:      addr,
		Billing:       addr,
	}

	bytes, err := client.CreateOrder(order)
	if err != nil {
		...
	}

	//quote
	quote := &sgb.Quote{
	 	Currency:      "USD",
	 	PaymentMethod: "paypall",
	 	ShipMethod:    "1YR_STORAGE",
	 	Declaration:   "TEST",
	 	Items:         items,
	 	Shipping:      addr,
	 	Billing:       addr,
	}

	bytes, err := client.Quote(quote)
	if err != nil {
		...
	}
}
```

# SEE ALSO

- [SilverGoldBull API docs](https://silvergoldbull.com/api-docs)