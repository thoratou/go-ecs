package ecs

import (
	"reflect"
)

type Component interface{}

//index for component type
//used to avoid performance issues due to reflect computing
type ComponentIndex uint32

type componentManager struct {
	indexes map[reflect.Type]ComponentIndex
	counter ComponentIndex
}

var componentManagerInst *componentManager = nil

func getComponentManager() *componentManager {
	if componentManagerInst == nil {
		componentManagerInst = &componentManager{
			indexes: make(map[reflect.Type]ComponentIndex),
			counter: 0,
		}
	}
	return componentManagerInst
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
		m.indexes[t] = index
		return index
	}
}
