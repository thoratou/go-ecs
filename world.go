package ecs

type World interface {
    AddSystem(s System)
    RemoveSystem(s System)

    //create a new Entity
    //warning: the entity isn't added to the world, AddEntity must be called after this
    NewEntity() *Entity

    AddEntity(e Entity) error
    UpdateEntity(i EntityId) error
    RemoveEntity(i EntityId) bool
    GetEntity(i EntityId) (Entity, bool)

    Update(d Delay)
}

type Delay float32