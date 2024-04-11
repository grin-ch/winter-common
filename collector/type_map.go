package collector

import "sync"

// TypeMap is a generic wrapper type of sync.Map
type TypeMap[k comparable, V any] struct {
	m sync.Map
}

func (tm *TypeMap[K, V]) CompareAndDelete(key K, old V) bool {
	return tm.m.CompareAndDelete(key, old)
}

func (tm *TypeMap[K, V]) CompareAndSwap(key K, old V, new V) bool {
	return tm.m.CompareAndSwap(key, old, new)
}

func (tm *TypeMap[K, V]) Delete(key K) {
	tm.m.Delete(key)
}

func (tm *TypeMap[K, V]) Load(key K) (val V, loaded bool) {
	value, loaded := tm.m.Load(key)
	if !loaded {
		return
	}
	return value.(V), true
}

func (tm *TypeMap[K, V]) LoadAndDelete(key K) (val V, loaded bool) {
	value, loaded := tm.m.LoadAndDelete(key)
	if !loaded {
		return
	}
	return value.(V), true
}

func (tm *TypeMap[K, V]) LoadOrStore(key K, value V) (val V, loaded bool) {
	actual, loaded := tm.m.LoadOrStore(key, value)
	if !loaded {
		return
	}
	return actual.(V), true
}

func (tm *TypeMap[K, V]) Range(f func(k K, v V) bool) {
	tm.m.Range(func(key, value any) bool {
		return f(key.(K), value.(V))
	})
}

func (tm *TypeMap[K, V]) Store(key K, value V) {
	tm.m.Store(key, value)
}

func (tm *TypeMap[K, V]) Swap(key K, value V) (previous V, loaded bool) {
	old, loaded := tm.m.Swap(key, value)
	if !loaded {
		return
	}
	return old.(V), true
}
