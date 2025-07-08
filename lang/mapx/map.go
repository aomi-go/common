package mapx

func Keys[K comparable, V any](value map[K]V) []K {
	var keys = make([]K, 0)
	for k := range value {
		keys = append(keys, k)
	}
	return keys
}

func ContainsKey[K comparable, V any](value map[K]V, key K) bool {
	if value == nil {
		return false
	}
	if _, ok := value[key]; ok {
		return true
	}
	return false
}

func Values[K comparable, V any](value map[K]V) []V {
	var values = make([]V, 0)
	for _, v := range value {
		values = append(values, v)
	}
	return values
}

// PutAll merges the contents of v2 into v1.
func PutAll[K comparable, V any](v1, v2 map[K]V) {
	for k, v := range v2 {
		v1[k] = v
	}
}

// ComputeIfAbsent computes the value for the given key if it is not already present in the map.
func ComputeIfAbsent[K comparable, V any](m map[K]V, key K, createValueFunc func(K) V) V {
	if val, ok := m[key]; ok {
		return val
	}
	v := createValueFunc(key)
	m[key] = v
	return v
}
