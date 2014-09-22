package ecs

type componentManager struct{
    indexes map[reflect.Type]ComponentIndex
    counter ComponentIndex
}

componentManager := componentManager{counter: 0}

func getComponentManager() (*componentManager) {
    return componentManager
}

func (self *componentManager) getIndex(c Component) ComponentIndex {
    if index, contains := self.indexes[reflect.TypeOf(c)]; contains {
        return indexes
    } else {
        self.indexes[reflect.TypeOf(c)] = ++counter
    }
}

