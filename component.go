package ecs

import (
	"reflect"
)

type Component interface {
	//index for component type
	//only one index value per component type should be provided, even for different instances
	GetIndex() ComponentIndex
}

//index for component type
//used to avoid performance issues due to reflect computing
type ComponentIndex uint32

type componentManager struct {
	indexes map[reflect.Type]ComponentIndex
	counter ComponentIndex
}

var componentManagerInst componentManager = componentManager{counter: 0}

func getComponentManager() *componentManager {
	return &componentManagerInst
}

func (m *componentManager) getIndex(c Component) ComponentIndex {
	return m.getIndexFromType(reflect.TypeOf(c))
}

func (m *componentManager) getIndexFromType(t reflect.Type) ComponentIndex {
	if index, exists := m.indexes[t]; exists {
		return index
	} else {
		index := m.counter
		m.counter++
		m.indexes[t] = m.counter
		return index
	}
}
