package ecs

import (
	"reflect"
	"testing"
)

type DummyComponent1 struct{}
type DummyComponent2 struct{}
type DummyComponent3 struct{}

func TestEmptyAspect(t *testing.T) {
	emptyAspect := NewAspect()

	world := NewWorld()

	emptyEntity := world.NewEntity()

	if !emptyAspect.Match(emptyEntity) {
		t.Fatal("By default, an empty aspect should match all entities")
	}

	dummyEntity := world.NewEntity()
	dummyEntity.AddComponent(DummyComponent1{})

	if !emptyAspect.Match(emptyEntity) {
		t.Fatal("By default, an empty aspect should match all entities")
	}

	dummyEntity.AddComponent(DummyComponent2{})
	dummyEntity.AddComponent(DummyComponent3{})

	if !emptyAspect.Match(emptyEntity) {
		t.Fatal("By default, an empty aspect should match all entities")
	}
}

func TestOneAspect(t *testing.T) {
	oneAspect := NewAspect().One(DummyComponent1{})

	world := NewWorld()

	emptyEntity := world.NewEntity()

	if oneAspect.Match(emptyEntity) {
		t.Fatal("The entity should not match the aspect")
	}

	dummyEntity := world.NewEntity()
	dummyEntity.AddComponent(DummyComponent1{})

	if !oneAspect.Match(dummyEntity) {
		t.Fatal("The entity should match the aspect")
	}

	dummyEntity.AddComponent(DummyComponent2{})
	dummyEntity.AddComponent(DummyComponent3{})

	if !oneAspect.Match(dummyEntity) {
		t.Fatal("The entity should match the aspect")
	}
}

func TestAspectParameters(t *testing.T) {
	oneAspectFromComponent := NewAspect().One(DummyComponent1{})
	oneAspectFromType := NewAspect().One(reflect.TypeOf(DummyComponent1{}))
	oneAspectFromId := NewAspect().One(getComponentManager().getIndex(DummyComponent1{}))

	world := NewWorld()

	dummyEntity := world.NewEntity()
	dummyEntity.AddComponent(DummyComponent1{})

	if !oneAspectFromComponent.Match(dummyEntity) {
		t.Fatal("The entity should match the aspect")
	}

	if !oneAspectFromType.Match(dummyEntity) {
		t.Fatal("The entity should match the aspect")
	}

	if !oneAspectFromId.Match(dummyEntity) {
		t.Fatal("The entity should match the aspect")
	}
}
