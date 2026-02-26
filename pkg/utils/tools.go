package utils

func ContainsMaps(parent map[string]string, subset map[string]string) bool {
	if len(subset) == 0 {
		return true
	}
	if len(parent) == 0 {
		return false
	}

	for k, v := range subset {
		parentVal, exists := parent[k]
		if !exists || parentVal != v {
			return false
		}
	}
	return true
}

func MergeMaps(base map[string]string, overrides map[string]string) map[string]string {
	result := make(map[string]string)

	for k, v := range base {
		result[k] = v
	}

	for k, v := range overrides {
		result[k] = v
	}

	return result
}
