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

func HumanBytes(size uint64) (human string) {
	switch {
	case size > EB*2:
		human = fmt.Sprintf("%d EB", size/EB)
	case size > PB*2:
		human = fmt.Sprintf("%d PB", size/PB)
	case size > TB*2:
		human = fmt.Sprintf("%d TB", size/TB)
	case size > GB*2:
		human = fmt.Sprintf("%d GB", size/GB)
	case size > MB*2:
		human = fmt.Sprintf("%d MB", size/MB)
	case size > KB*2:
		human = fmt.Sprintf("%d KB", size/KB)
	default:
		human = fmt.Sprintf("%d bytes", size)
	}
	return
}
