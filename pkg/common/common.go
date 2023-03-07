package common

func Contains[T comparable](sources []T, target T) bool {
	for _, source := range sources {
		if source == target {
			return true
		}
	}
	return false
}
