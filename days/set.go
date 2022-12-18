package days

var exists interface{}

type set[T comparable] struct {
	m map[T]interface{}
}

func NewSet[T comparable](initalValues ...T) *set[T] {
	s := &set[T]{make(map[T]interface{})}

	for _, value := range initalValues {
		s.Add(value)
	}

	return s
}

func (s *set[T]) Add(toAdd T) {
	s.m[toAdd] = exists
}

func (s *set[T]) AddSlice(toAdd []T) {
	for _, element := range toAdd {
		s.Add(element)
	}
}

func (s *set[T]) Remove(toRemove T) {
	delete(s.m, toRemove)
}

func (s *set[T]) Contains(toCheck T) bool {
	_, contains := s.m[toCheck]
	return contains
}

func (s *set[T]) Length() int {
	return len(s.m)
}

func (s *set[T]) ToSlice() []T {
	result := make([]T, 0, s.Length())

	for k, _ := range s.m {
		result = append(result, k)
	}

	return result
}
