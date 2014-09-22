package ecs

type world struct{
    entityCounter EntityId
    entities map[EntityId]Entity
    systems []Systems
}

func NewWorld() *World {
    return &world{entityCounter: 0}
}

func (w *world) AddSystem(s System) {
    w.systems = append(w.systems, s)
}

func (w *world) RemoveSystem(s System) {
    delete(w.systems, s)
}

func (w *world) NewEntity() *Entity {
    return &entity{id: ++w.entityCounter, world: w}
}

func (w *world) AddEntity(e Entity) error {
    Entity id = e.GetId()
    if _, contains := w.entities[id]; contains {
        return NewError("entity with id ", id, " already added to world")
    }
    w.entities[id] = e
    updateSystemsWithEntity(e)
    return nil
}

func (w *world) UpdateEntity(e Entity) error {
    Entity id = e.GetId()
    if _, contains := w.entities[id]; !contains {
        return NewError("entity with id ", id, " does not exist into the world, cannot update it")
    } else {
        w.updateSystemsWithEntity(e)
        return nil
    }
}

func (w *world) RemoveEntity(i EntityId) bool {
    _, contains := w.entities[i]

    if contains {
        w.removeEntityFromSystems(e)
    }

    delete(w.entities, i)
    return contains
}

func (w world) GetEntity(i EntityId) (Entity, bool) {
    return w.entities[i]
}

func (w world) Update(d Delay) {
    for _, s := range w.systems {
        s.Update(d)
    }
}

func (w world) updateSystemsWithEntity(e Entity) {
    for _, s := range w.systems {
        s.UpdateRegistration(e)
    }
}

func (w world) removeEntityFromSystems(e Entity) {
    for _, s := range w.systems {
        s.Unregister(e)
    }
}