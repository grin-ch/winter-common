package slice_util

func ToMap[S ~[]E, E any, K comparable, V any](s S, fn func(elem E) (K, V)) map[K]V {
	if s == nil {
		return nil
	}
	m := make(map[K]V, len(s))
	for _, v := range s {
		k, v := fn(v)
		m[k] = v
	}
	return m
}

// Distinct 去重
func Distinct[S ~[]E, E comparable](src S) S {
	return DistinctFunc(src, func(val E) E { return val })
}

func DistinctFunc[S ~[]E, E any, K comparable](src S, fn func(val E) K) S {
	if src == nil {
		return nil
	}
	if len(src) < 2 {
		dst := make(S, len(src))
		copy(dst, src)
		return dst
	}

	list := make(S, 0, len(src))
	m := make(map[K]struct{}, len(src))
	for _, v := range src {
		key := fn(v)
		if _, has := m[key]; !has {
			list = append(list, v)
			m[key] = struct{}{}
		}
	}
	return list
}

// Map 根据 []E1 生成 []E2
func Map[S ~[]E1, E1 any, E2 any](src S, fn func(elem E1) E2) []E2 {
	if src == nil {
		return nil
	}
	list := make([]E2, 0, len(src))
	for _, v := range src {
		list = append(list, fn(v))
	}
	return list
}

func Filter[S ~[]E, E any](s S, fn func(elem E) bool) S {
	if s == nil {
		return nil
	}
	f := make(S, 0, len(s))
	for _, v := range s {
		if fn(v) {
			f = append(f, v)
		}
	}
	return f
}

func Contains[S ~[]E, E comparable](s S, elem E) bool {
	return ContainsFunc(s, func(e E) bool { return e == elem })
}

func ContainsFunc[S ~[]E, E any](s S, fn func(e E) bool) bool {
	for _, v := range s {
		if fn(v) {
			return true
		}
	}
	return false
}
