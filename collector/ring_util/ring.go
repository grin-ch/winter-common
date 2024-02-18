package ring_util

import "container/ring"

type Ring[T any] struct {
	r *ring.Ring
}

func NewRing[T any](list []T) *Ring[T] {
	r := ring.New(len(list))
	for _, v := range list {
		r.Value = v
		r = r.Next()
	}
	return &Ring[T]{
		r: r,
	}
}

func (ring *Ring[T]) CurValue() (T, bool) {
	v, ok := ring.r.Value.(T)
	return v, ok
}

func (ring *Ring[T]) Turn() T {
	ring.Next()
	v, _ := ring.CurValue()
	return v
}

func (ring *Ring[T]) Next() {
	ring.r = ring.r.Next()
}

func (ring *Ring[T]) Prev() {
	ring.r = ring.r.Prev()
}

func (ring *Ring[T]) Move(n int) {
	ring.r = ring.r.Move(n)
}

func (ring *Ring[T]) Link(r *Ring[T]) {
	ring.r = ring.r.Link(r.r)
}

func (ring *Ring[T]) Unlink(n int) {
	ring.r = ring.r.Unlink(n)
}

func (ring *Ring[T]) Len() int {
	return ring.r.Len()
}

func (ring *Ring[T]) Do(fn func(val T)) {
	ring.r.Do(func(a any) {
		fn(a.(T))
	})
}
