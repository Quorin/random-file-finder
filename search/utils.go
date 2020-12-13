package search

// SliceContain search the slice and return if item is inside.
func SliceContain(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

func FindAnyExtension(extensions []string) bool {
	if len(extensions) == 0 {
		return false
	}

	return extensions[0] == AllExtensionsChar
}
