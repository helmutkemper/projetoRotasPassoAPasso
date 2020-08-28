package testDataSource

import (
	"commomTypes"
	"fmt"
)

func ExampleTestDataSource_GetStretchByOrigin() {
	var err error
	var origin commomTypes.Origin
	var destination commomTypes.Destination
	var price commomTypes.Price
	var dataList []commomTypes.RouteStretch

	ds := TestDataSource{}

	origin = commomTypes.Origin("Recife")
	destination = commomTypes.Destination("Jaboatão dos Guararapes")
	price = commomTypes.Price(20)
	ds.AddRoute(origin, destination, price)

	origin = commomTypes.Origin("Recife")
	destination = commomTypes.Destination("Cabo de Santo Agostinho")
	price = commomTypes.Price(80)
	ds.AddRoute(origin, destination, price)

	dataList, err = ds.GetStretchByOrigin("Recife")
	if err != nil {
		panic(err)
	}

	for _, stretch := range dataList {
		fmt.Printf("origin: %v\n", stretch.Origin)
		fmt.Printf("destination: %v\n", stretch.Destination)
		fmt.Printf("price: %v\n\n", stretch.Price)
	}

	// Output:
	// origin: Recife
	// destination: Jaboatão dos Guararapes
	// price: 20
	//
	// origin: Recife
	// destination: Cabo de Santo Agostinho
	// price: 80
}
