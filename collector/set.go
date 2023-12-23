package collector

type Set[T comparable] interface {
	Add(elems ...T)
	Del(elems ...T)
	Has(elem T) bool
	Size() int
	Range(fn func(elem T) bool)
	ToSlice() []T
}

func NewSet[T comparable](elems ...T) Set[T] {
	s := new(set[T])
	s.m = make(map[T]struct{}, len(elems))
	s.Add(elems...)
	return s
}

type set[T comparable] struct {
	m map[T]struct{}
}

// Add implements Set.
func (s *set[T]) Add(elems ...T) {
	for _, v := range elems {
		s.m[v] = struct{}{}
	}
}

// Del implements Set.
func (s *set[T]) Del(elems ...T) {
	for _, v := range elems {
		delete(s.m, v)
	}
}

// Has implements Set.
func (s *set[T]) Has(elem T) bool {
	_, has := s.m[elem]
	return has
}

// Size implements Set.
func (s *set[T]) Size() int {
	return len(s.m)
}

// Range implements Set.
func (s *set[T]) Range(fn func(elem T) bool) {
	for k := range s.m {
		if !fn(k) {
			break
		}
	}
}

// ToSlice implements Set.
func (s *set[T]) ToSlice() []T {
	list := make([]T, 0, len(s.m))
	for k := range s.m {
		list = append(list, k)
	}
	return list
}
