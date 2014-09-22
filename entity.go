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