package utils

func ContainsEmptyString(strings ...string) bool {
	for _, s := range strings {
		if s == "" {
			return true
		}
	}

	return false
}