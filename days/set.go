package days

var exists interface{}

type set[T comparable] struct {
	m map[T]interface{}
}

func NewSet[T comparable](initalValues ...T) *set[T] {
	s := &set[T] {
		make(map[T]interface{}),
	}

	for _, value := range initalValues {
		s.Add(value)
	}

	return s
}

func (this *set[T]) Add(toAdd T) {
	this.m[toAdd] = exists
}

func (this *set[T]) Remove(toRemove T) {
	delete(this.m, toRemove)
}

func (this *set[T]) Contains(toCheck T) bool {
	_, contains := this.m[toCheck]
	return contains
}

func (this *set[T]) Length() int {
	return len(this.m)
}

func (this *set[T]) ToSlice() []T {
	result := make([]T, 0, len(this.m))

	for k, _ := range this.m {
		result = append(result, k)
	}

	return result
}
