package ecs

type System interface {
    //world containing this system
    GetWorld() World

    Update(d Delay)

    UpdateRegistration(e Entity)
    Unregister(e Entity)
}
