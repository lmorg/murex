package man

import "time"

func cacheTtl() time.Time {
	return time.Now().Add(time.Hour * 24 * 31)
}
