package testDataSource

import (
	"commomTypes"
	"sync"
)

type TestDataSource struct {
	dataList []commomTypes.RouteStretch
	mutex    sync.Mutex
}
