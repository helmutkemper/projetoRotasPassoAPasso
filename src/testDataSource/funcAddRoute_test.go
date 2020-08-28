package testDataSource

import (
	"commomTypes"
	"testing"
)

func TestDataSource_AddRoute(t *testing.T) {
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

	if len(ds.dataList) != 1 {
		t.Fail()
	}

	if ds.dataList[0].Origin != "Recife" {
		t.Fail()
	}

	if ds.dataList[0].Destination != "Jaboatão dos Guararapes" {
		t.Fail()
	}

	origin = commomTypes.Origin("Recife")
	destination = commomTypes.Destination("Cabo de Santo Agostinho")
	price = commomTypes.Price(80)
	ds.AddRoute(origin, destination, price)

	if len(ds.dataList) != 2 {
		t.Fail()
	}

	if ds.dataList[0].Origin != "Recife" {
		t.Fail()
	}

	if ds.dataList[0].Destination != "Jaboatão dos Guararapes" {
		t.Fail()
	}

	if ds.dataList[1].Origin != "Recife" {
		t.Fail()
	}

	if ds.dataList[1].Destination != "Cabo de Santo Agostinho" {
		t.Fail()
	}

	dataList, err = ds.GetStretchByDestination("Jaboatão dos Guararapes")
	if err != nil {
		panic(err)
	}

	if dataList[0].Origin != "Recife" {
		t.Fail()
	}

	if dataList[0].Destination != "Jaboatão dos Guararapes" {
		t.Fail()
	}
}
