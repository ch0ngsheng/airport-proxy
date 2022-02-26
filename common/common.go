package common

import "net/http"

// HandlerMap 机场名称和对象的映射
var HandlerMap = map[string]IAirPort{}

// IAirPort 不同机场实现的接口
type IAirPort interface {
	Name() string
	// RemoveRateNode 移除订阅中的倍率节点
	RemoveRateNode() http.HandlerFunc
}
