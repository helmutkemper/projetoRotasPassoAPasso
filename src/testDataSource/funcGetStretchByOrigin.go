package testDataSource

import (
	"commomTypes"
	"errors"
)

func (el *TestDataSource) GetStretchByOrigin(
	origin commomTypes.Origin,
) (
	dataList []commomTypes.RouteStretch,
	err error,
) {

	el.mutex.Lock()
	defer el.mutex.Unlock()

	dataList = make([]commomTypes.RouteStretch, 0)

	for _, line := range el.dataList {
		if line.Origin == origin {
			dataList = append(dataList, line)
		}
	}

	if len(dataList) == 0 {
		err = errors.New(KErroStretchNotFound)
	}

	return
}
