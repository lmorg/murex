package utils

import "fmt"

// Exportable byte denominations
const (
	_  = iota
	KB = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

// Byte denominations as floats to make division cheaper (less casting at runtime)
const (
	_           = iota
	fKB float64 = 1 << (10 * iota)
	fMB
	fGB
	fTB
	fPB
	fEB
	//fZB
	//fYB
)

// HumanBytes converts n bytes into a human readable format
func HumanBytes(size uint64) (human string) {
	switch {
	//case size > YB*2:
	//	human = fmt.Sprintf("%.8f YB", float64(size)/fYB)
	//case size > ZB*2:
	//	human = fmt.Sprintf("%.8f ZB", float64(size)/fZB)
	case size > EB*2:
		human = fmt.Sprintf("%.6f EB", float64(size)/fEB)
	case size > PB*2:
		human = fmt.Sprintf("%.6f PB", float64(size)/fPB)
	case size > TB*2:
		human = fmt.Sprintf("%.4f TB", float64(size)/fTB)
	case size > GB*2:
		human = fmt.Sprintf("%.4f GB", float64(size)/fGB)
	case size > MB*2:
		human = fmt.Sprintf("%.2f MB", float64(size)/fMB)
	case size > KB*2:
		human = fmt.Sprintf("%.2f KB", float64(size)/fKB)
	default:
		human = fmt.Sprintf("%0d bytes", size)
	}
	return
}
