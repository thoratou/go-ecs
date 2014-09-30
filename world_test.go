package ecs

import (
	"bytes"
	"testing"
)

//components

type VoidComponent struct{}

type PositionComponent struct {
	x uint32
	y uint32
}

//systems

type ComponentIndexSystem struct {
	world    World
	entities map[EntityId]Entity
	result   map[EntityId]string
}

func (c *ComponentIndexSystem) GetWorld() World {
	return c.world
}

func (c *ComponentIndexSystem) UpdateRegistration(e Entity) {
	id := e.GetId()
	if _, exists := c.entities[id]; !exists {
		c.entities[id] = e
	}
}

func (c *ComponentIndexSystem) Unregister(e Entity) {
	id := e.GetId()
	if _, exists := c.entities[id]; exists {
		delete(c.entities, id)
	}
}

var voidIndex ComponentIndex = getComponentManager().getIndex(VoidComponent{})
var positionIndex ComponentIndex = getComponentManager().getIndex(PositionComponent{})

func (c *ComponentIndexSystem) Update(d Delay) {
	c.result = make(map[EntityId]string) //clear the result map
	for id, entity := range c.entities {
		buffer := bytes.NewBufferString("|")
		if _, exists := entity.GetComponent(positionIndex); exists {
			buffer.WriteString("position|")
		}
		if _, exists := entity.GetComponent(voidIndex); exists {
			buffer.WriteString("void|")
		}

		c.result[id] = buffer.String()
	}
}

func TestBasic(t *testing.T) {
	world := NewWorld()

	system := ComponentIndexSystem{
		world:    world,
		entities: make(map[EntityId]Entity),
		result:   make(map[EntityId]string),
	}

	world.AddSystem(&system)

	empty := world.NewEntity()
	world.AddEntity(empty)

	if len(system.entities) != 1 {
		t.Fatal("system should contain one entity")
	}

	if _, exists := system.entities[empty.GetId()]; !exists {
		t.Fatal("system should contain an entity with id: ", empty.GetId())
	}

	if len(system.result) != 0 {
		t.Fatal("result should be empty due to no Update() call")
	}

	world.Update(0.0)

	if len(system.entities) != 1 {
		t.Fatal("system should contain one entity")
	}

	if _, exists := system.entities[empty.GetId()]; !exists {
		t.Fatal("system should contain an entity with id: ", empty.GetId())
	}

	if len(system.result) != 1 {
		t.Fatal("result should contain one result after Update() call")
	}

	result, exists := system.result[empty.GetId()]
	if !exists {
		t.Fatal("wrong id in result map")
	}
	if result != "|" {
		t.Fatal("result gap, expected: |, received: ", result)
	}
}

func TestAddComponents(t *testing.T) {
	world := NewWorld()

	system := ComponentIndexSystem{
		world:    world,
		entities: make(map[EntityId]Entity),
		result:   make(map[EntityId]string),
	}

	world.AddSystem(&system)

	voidAndPosition := world.NewEntity()

	if err := voidAndPosition.AddComponent(VoidComponent{}); err != nil {
		t.Fatal("unexpected error at component addition: ", err)
	}

	if err := voidAndPosition.AddComponent(PositionComponent{1, 2}); err != nil {
		t.Fatal("unexpected error at component addition: ", err)
	}

	world.AddEntity(voidAndPosition)

	world.Update(0.0)

	if len(system.entities) != 1 {
		t.Fatal("system should contain one entity")
	}

	if _, exists := system.entities[voidAndPosition.GetId()]; !exists {
		t.Fatal("system should contain an entity with id: ", voidAndPosition.GetId())
	}

	if len(system.result) != 1 {
		t.Fatal("result should contain one result after Update() call")
	}

	result, exists := system.result[voidAndPosition.GetId()]
	if !exists {
		t.Fatal("wrong id in result map")
	}
	if result != "|position|void|" {
		t.Fatal("result gap, expected: |position|void|, received: ", result)
	}
}

func TestUpdateEntity(t *testing.T) {
	world := NewWorld()

	system := ComponentIndexSystem{
		world:    world,
		entities: make(map[EntityId]Entity),
		result:   make(map[EntityId]string),
	}

	world.AddSystem(&system)

	voidAndPosition := world.NewEntity()

	if err := voidAndPosition.AddComponent(VoidComponent{}); err != nil {
		t.Fatal("unexpected error at component addition: ", err)
	}

	if err := voidAndPosition.AddComponent(PositionComponent{1, 2}); err != nil {
		t.Fatal("unexpected error at component addition: ", err)
	}

	world.AddEntity(voidAndPosition)

	world.Update(0.0)

	if len(system.entities) != 1 {
		t.Fatal("system should contain one entity")
	}

	if _, exists := system.entities[voidAndPosition.GetId()]; !exists {
		t.Fatal("system should contain an entity with id: ", voidAndPosition.GetId())
	}

	if len(system.result) != 1 {
		t.Fatal("result should contain one result after Update() call")
	}

	{
		result, exists := system.result[voidAndPosition.GetId()]
		if !exists {
			t.Fatal("wrong id in result map")
		}
		if result != "|position|void|" {
			t.Fatal("result gap, expected: |position|void|, received: ", result)
		}
	}

	//update the entity by removing a component
	voidAndPosition.RemoveComponent(voidIndex)
	world.UpdateEntity(voidAndPosition)

	//not recomputed yet
	{
		result, exists := system.result[voidAndPosition.GetId()]
		if !exists {
			t.Fatal("wrong id in result map")
		}
		if result != "|position|void|" {
			t.Fatal("result gap, expected: |position|void|, received: ", result)
		}
	}

	world.Update(0.0)

	//result updated
	{
		result, exists := system.result[voidAndPosition.GetId()]
		if !exists {
			t.Fatal("wrong id in result map")
		}
		if result != "|position|" {
			t.Fatal("result gap, expected: |position|, received: ", result)
		}

	}
}

func TestRemoveEntities(t *testing.T) {
	world := NewWorld()

	system := ComponentIndexSystem{
		world:    world,
		entities: make(map[EntityId]Entity),
		result:   make(map[EntityId]string),
	}

	world.AddSystem(&system)

	voidAndPosition := world.NewEntity()

	if err := voidAndPosition.AddComponent(VoidComponent{}); err != nil {
		t.Fatal("unexpected error at component addition: ", err)
	}

	if err := voidAndPosition.AddComponent(PositionComponent{1, 2}); err != nil {
		t.Fatal("unexpected error at component addition: ", err)
	}

	world.AddEntity(voidAndPosition)

	world.Update(0.0)

	if len(system.entities) != 1 {
		t.Fatal("system should contain one entity")
	}

	if _, exists := system.entities[voidAndPosition.GetId()]; !exists {
		t.Fatal("system should contain an entity with id: ", voidAndPosition.GetId())
	}

	if len(system.result) != 1 {
		t.Fatal("result should contain one result after Update() call")
	}

	{
		result, exists := system.result[voidAndPosition.GetId()]
		if !exists {
			t.Fatal("wrong id in result map")
		}
		if result != "|position|void|" {
			t.Fatal("result gap, expected: |position|void|, received: ", result)
		}
	}

	world.RemoveEntity(voidAndPosition)

	if voidAndPosition.GetId() != UnknownId {
		t.Fatal("entity ID should have been resetted")
	}

	if len(system.result) != 1 {
		t.Fatal("result should contain one result after Update() call")
	}

	if _, exists := system.result[voidAndPosition.GetId()]; exists {
		t.Fatal("as the id should be set to Unknown, the data should not be accessible anymore")
	}

	world.Update(0.0)

	if len(system.entities) != 0 {
		t.Fatal("system should be empty")
	}

	if len(system.result) != 0 {
		t.Fatal("result should be empty")
	}
}
