package days

var exists = struct{}{}

type set struct {
	m map[int]struct{}
}

func NewSet() *set {
	s := &set{}
	s.m = make(map[int]struct{})
	return s
}

func (s *set) Add(toAdd int) {
	s.m[toAdd] = exists
}

func (s *set) Remove(toRemove int) {
	delete(s.m, toRemove)
}

func (s *set) Contains(toCheck int) bool {
	_, contains := s.m[toCheck]
	return contains
}

