package utils

import "fmt"

const (
	KB uint64 = 1024
	MB uint64 = KB * 1024
	GB uint64 = MB * 1024
	TB uint64 = GB * 1024
	PB uint64 = TB * 1024
	EB uint64 = PB * 1024
)

const (
	fKB float64 = 1024
	fMB float64 = fKB * 1024
	fGB float64 = fMB * 1024
	fTB float64 = fGB * 1024
	fPB float64 = fTB * 1024
	fEB float64 = fPB * 1024
)

// HumanBytes converts n bytes into a human readable format
func HumanBytes(size uint64) (human string) {
	switch {
	case size > EB*2:
		human = fmt.Sprintf("%.8f EB", float64(size)/fEB)
	case size > PB*2:
		human = fmt.Sprintf("%.8f PB", float64(size)/fPB)
	case size > TB*2:
		human = fmt.Sprintf("%.6f TB", float64(size)/fTB)
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
