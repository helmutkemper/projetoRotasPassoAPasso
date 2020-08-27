package testDataSource

import (
	"commomTypes"
	"errors"
)

func (el *TestDataSource) GetStretchByDestination(
	destination commomTypes.Destination,
) (
	stretchList []commomTypes.RouteStretch,
	err error,
) {

	el.mutex.Lock()
	defer el.mutex.Unlock()

	stretchList = make([]commomTypes.RouteStretch, 0)

	for _, dataLine := range el.dataList {
		if dataLine.Destination == destination {
			stretchList = append(stretchList, dataLine)
		}
	}

	if len(stretchList) == 0 {
		err = errors.New(KErroStretchNotFound)
	}

	return
}
