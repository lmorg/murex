package humannumbers

// ColumnLetter takes an int and converts it to an Excel-style column letter
// reference (eg `C` or `AB`)
func ColumnLetter(i int) string {
	var col string

	for i >= 0 {
		mod := i % 26
		col = string([]byte{byte(mod) + 65}) + col
		i = ((i - mod) / 26) - 1
	}

	return col
}
