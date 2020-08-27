package testDataSource

import (
	"commomTypes"
)

// Adiciona uma nova rota a fonte de dados
func (el *TestDataSource) AddRoute(
	origin commomTypes.Origin,
	destination commomTypes.Destination,
	price commomTypes.Price,
) {

	el.mutex.Lock()
	defer el.mutex.Unlock()

	if len(el.dataList) == 0 {
		el.dataList = make([]commomTypes.RouteStretch, 0)
	}

	el.dataList = append(el.dataList, commomTypes.RouteStretch{
		Origin:      origin,
		Destination: destination,
		Price:       price,
	})
}
