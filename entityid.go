package ecs

type EntityId uint32

const (
	UnknownId EntityId = EntityId(^uint32(0)) //0xFFFFFFFF
)

type entityIdStack struct{ vec []EntityId }

func newEntityIdStack() *entityIdStack {
	return &entityIdStack{
		vec: make([]EntityId, 0),
	}
}

func (s *entityIdStack) Empty() bool {
	return len(s.vec) == 0
}

func (s *entityIdStack) Peek() EntityId {
	return s.vec[len(s.vec)-1]
}

func (s *entityIdStack) Len() int {
	return len(s.vec)
}

func (s *entityIdStack) Put(i EntityId) {
	s.vec = append(s.vec, i)
}

func (s *entityIdStack) Pop() EntityId {
	d := s.vec[len(s.vec)-1]
	s.vec = s.vec[:len(s.vec)-1]
	return d
}
