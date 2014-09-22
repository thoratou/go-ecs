package ecs

import (
	"reflect"
)

type entity struct {
    components map[ComponentIndex]Component
    id EntityId
    world World
}

func (e entity) GetId() EntityId {
    return e.id
}

func (e entity) GetWorld() World {
    return e.world
}

func (e *entity) AddComponent(c Component) error {
    i := getComponentManager().getIndex(c)
    if _, contains := e.components[i]; contains {
        return NewError("component with type ", reflect.TypeOf(c), "already registered to entity ", id)
    }
    return nil
}

func (e entity) GetComponent(i ComponentIndex) (Component, bool){
    return e.components[i]

func (e entity) RemoveComponent(i ComponentIndex) bool {
    _, contains := e.components[i]
    delete(e.components, i)
    return contains
}

func (e entity) GetComponentFromType(t reflect.Type) (Component, bool) {
    return e.GetComponent(getComponentManager().getIndexFromType(t))
}

func (e *entity) RemoveComponentFromType(t reflect.Type) bool{
    return e.RemoveComponent(getComponentManager().getIndexFromType(t))
}
