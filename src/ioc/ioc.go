package ioc

import (
	"fmt"
)

var components = make(map[string]interface{})

func Provide(t string) (interface{}, error) {
	r := components[t]
	if r == nil {
		return nil, fmt.Errorf("指定名称的组件不存在: %s", t)
	}
	return r, nil
}

func PutIn(t string, component interface{}) error {
	existComponent := components[t]
	if existComponent != nil {
		return fmt.Errorf("指定名称的组件已存在: %s", t)
	} else {
		components[t] = component
		return nil
	}
}
