package ecs

type Component interface{
    //index for component type
    //only one index value per component type should be provided, even for different instances
    GetIndex() ComponentIndex
}

//index for component type
//used to avoid performance issues due to reflect computing
type ComponentIndex uint32