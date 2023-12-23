package map_util

func ToSlice[M ~map[K]V, K comparable, V any, E any](m M, fn func(k K, v V) E) []E {
	if m == nil {
		return nil
	}
	list := make([]E, 0, len(m))
	for k, v := range m {
		list = append(list, fn(k, v))
	}
	return list
}

// Keys 获取map的key列表
func Keys[M ~map[K]V, K comparable, V any](m M) []K {
	return KeysFunc(m, func(key K) K { return key })
}

func KeysFunc[M ~map[K]V, K comparable, V, E any](m M, fn func(key K) E) []E {
	return ToSlice(m, func(k K, _ V) E { return fn(k) })
}

// Values 获取map的value列表
func Values[M ~map[K]V, K comparable, V any](m M) []V {
	return ValuesFunc(m, func(value V) V { return value })
}

func ValuesFunc[M ~map[K]V, K comparable, V, E any](m M, fn func(value V) E) []E {
	return ToSlice(m, func(_ K, v V) E { return fn(v) })
}

func ContainsValue[M ~map[K]V, K comparable, V any](m M, fn func(v V) bool) bool {
	return ContainsFunc(m, func(_ K, v V) bool { return fn(v) })
}

func ContainsFunc[M ~map[K]V, K comparable, V any](m M, fn func(k K, v V) bool) bool {
	for k, v := range m {
		if fn(k, v) {
			return true
		}
	}
	return false
}

// Merge will merged all maps
func Merge[S ~map[K]V, K comparable, V any](maps ...S) S {
	if len(maps) == 0 {
		return nil
	}

	all := make(S, len(maps))
	for _, m := range maps {
		for k, v := range m {
			all[k] = v
		}
	}
	return all
}
