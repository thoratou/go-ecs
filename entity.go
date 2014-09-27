package ecs

import (
	"reflect"
)

type Entity interface {
	//global entity id (should be unique in all the world)
	GetId() EntityId

	//world containing this entity
	GetWorld() World

	//component interfaces
	AddComponent(c Component) error

	GetComponent(i ComponentIndex) (Component, bool)
	RemoveComponent(i ComponentIndex) bool

	GetComponentFromType(t reflect.Type) (Component, bool)
	RemoveComponentFromType(t reflect.Type) bool
}

type EntityId uint32

type entity struct {
	components map[ComponentIndex]Component
	id         EntityId
	world      World
}

func (e *entity) GetId() EntityId {
	return e.id
}

func (e *entity) GetWorld() World {
	return e.world
}

func (e *entity) AddComponent(c Component) error {
	i := getComponentManager().getIndex(c)
	if _, exists := e.components[i]; exists {
		return NewError("component with type ", reflect.TypeOf(c), "already registered to entity ", i)
	}
	return nil
}

func (e *entity) GetComponent(i ComponentIndex) (Component, bool) {
	component, exists := e.components[i]
	return component, exists
}

func (e *entity) RemoveComponent(i ComponentIndex) bool {
	_, exists := e.components[i]
	delete(e.components, i)
	return exists
}

func (e *entity) GetComponentFromType(t reflect.Type) (Component, bool) {
	component, exists := e.GetComponent(getComponentManager().getIndexFromType(t))
	return component, exists
}

func (e *entity) RemoveComponentFromType(t reflect.Type) bool {
	return e.RemoveComponent(getComponentManager().getIndexFromType(t))
}
