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

func (s *set[T]) Copy() *set[T] {
	result := NewSet[T]()

	for ele := range s.Itterate() {
		result.Add(ele)
	}

	return result
}

func (s *set[T]) Add(toAdd T) {
	s.m[toAdd] = exists
}

func (s1 *set[T]) Combine(s2 *set[T]) *set[T] {
	if s1 == nil {
		return s2
	}

	if s2 == nil {
		return s1
	}

	result := s1.Copy()
	for ele := range s2.Itterate() {
		result.Add(ele)
	}

	return result
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

func (s *set[T]) Itterate() <-chan T {
	channel := make(chan T)

	go func() {
		for k, _ := range s.m {
			channel <- k
		}

		close(channel)
	}()

	return channel
}
