package ecs

type World interface {
	AddSystem(s System)
	RemoveSystem(s System)

	//create a new Entity
	//warning: the entity isn't added to the world, AddEntity must be called after this
	NewEntity() Entity

	AddEntity(e Entity) error
	UpdateEntity(e Entity) error
	RemoveEntity(e Entity) bool
	GetEntity(i EntityId) (Entity, bool)

	Update(d Delay)
}

type Delay float32

type world struct {
	entityCounter EntityId
	releasedIds   entityIdStack
	entities      map[EntityId]Entity
	systems       []System
}

func NewWorld() World {
	return &world{
		entityCounter: 0,
		entities:      make(map[EntityId]Entity),
		systems:       make([]System, 0),
	}
}

func (w *world) AddSystem(s System) {
	w.systems = append(w.systems, s)
}

func (w *world) RemoveSystem(s System) {
	for i, item := range w.systems {
		if item == s {
			w.systems = append(w.systems[:i], w.systems[i+1:]...)
			break
		}
	}
}

func (w *world) NewEntity() Entity {
	return &entity{
		components: make(map[ComponentIndex]Component),
		id:         UnknownId,
		world:      w,
	}
}

func (w *world) AddEntity(e Entity) error {
	oldId := e.GetId()

	if oldId != UnknownId {
		return NewError("entity with id ", oldId, " already added to world")
	}

	newId := w.allocateNewEntityId()
	e.SetId(newId)

	w.entities[newId] = e
	w.updateSystemsWithEntity(e)
	return nil
}

func (w *world) UpdateEntity(e Entity) error {
	id := e.GetId()
	if _, exists := w.entities[id]; !exists {
		return NewError("entity with id ", id, " does not exist into the world, cannot update it")
	} else {
		w.updateSystemsWithEntity(e)
		return nil
	}
}

func (w *world) RemoveEntity(e Entity) bool {
	id := e.GetId()
	_, exists := w.entities[id]

	if exists {
		w.removeEntityFromSystems(e)
		delete(w.entities, id)
		w.releaseEntityId(id)
		e.SetId(UnknownId)
	}

	return exists
}

func (w *world) GetEntity(i EntityId) (Entity, bool) {
	entity, exists := w.entities[i]
	return entity, exists
}

func (w *world) Update(d Delay) {
	for _, s := range w.systems {
		s.Update(d)
	}
}

func (w *world) updateSystemsWithEntity(e Entity) {
	for _, s := range w.systems {
		s.UpdateRegistration(e)
	}
}

func (w *world) removeEntityFromSystems(e Entity) {
	for _, s := range w.systems {
		s.Unregister(e)
	}
}

func (w *world) allocateNewEntityId() EntityId {
	if w.releasedIds.Empty() {
		id := w.entityCounter
		w.entityCounter++
		return id
	} else {
		id := w.releasedIds.Pop()
		return id
	}
}

func (w *world) releaseEntityId(id EntityId) {
	w.releasedIds.Put(id)
}
