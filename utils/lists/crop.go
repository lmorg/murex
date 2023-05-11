package lists

import "strings"

func CropPartial(list []string, partial string) []string {
	var items []string
	for i := range list {
		if strings.HasPrefix(list[i], partial) {
			items = append(items, list[i][len(partial):])
		}
	}

	return items
}

func CropPartialMapKeys(m map[string]string, partial string) map[string]string {
	cropped := make(map[string]string)
	for key, val := range m {
		if strings.HasPrefix(key, partial) {
			cropped[key[len(partial):]] = val
		}
	}

	return cropped
}
