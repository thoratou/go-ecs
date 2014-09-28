package ecs

import (
	"reflect"
)

type Aspect interface {
	One(args ...interface{}) Aspect
	All(args ...interface{}) Aspect
	Exclude(args ...interface{}) Aspect
	Match(e Entity) bool
}

type aspect struct {
	oneList     []ComponentIndex
	allList     []ComponentIndex
	excludeList []ComponentIndex
}

func NewAspect() Aspect {
	return &aspect{
		oneList:     make([]ComponentIndex, 0),
		allList:     make([]ComponentIndex, 0),
		excludeList: make([]ComponentIndex, 0),
	}
}

func (a *aspect) One(args ...interface{}) Aspect {
	a.add(&a.oneList, args)
	return a
}

func (a *aspect) All(args ...interface{}) Aspect {
	a.add(&a.allList, args)
	return a
}

func (a *aspect) Exclude(args ...interface{}) Aspect {
	a.add(&a.excludeList, args)
	return a
}

func (a *aspect) Match(e Entity) bool {
	hasOneMatch := true
	allMatch := true
	hasExcludeMatch := false

	if len(a.oneList) != 0 {
		hasOneMatch = false
		for _, one := range a.oneList {
			if _, exists := e.GetComponent(one); exists {
				hasOneMatch = true
			}
		}
	}

	if len(a.allList) != 0 {
		for _, all := range a.allList {
			if _, exists := e.GetComponent(all); !exists {
				allMatch = false
			}
		}
	}

	if len(a.excludeList) != 0 {
		for _, exclude := range a.excludeList {
			if _, exists := e.GetComponent(exclude); exists {
				hasExcludeMatch = true
			}
		}
	}

	return hasOneMatch && allMatch && !hasExcludeMatch
}

func (a *aspect) add(ids *[]ComponentIndex, args []interface{}) {
	for _, arg := range args {
		switch typedArg := arg.(type) {
		case ComponentIndex:
			*ids = append(*ids, typedArg)
		case reflect.Type:
			*ids = append(*ids, getComponentManager().getIndexFromType(typedArg))
		default:
			*ids = append(*ids, getComponentManager().getIndex(typedArg))
		}
	}
}
