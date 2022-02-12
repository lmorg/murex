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
