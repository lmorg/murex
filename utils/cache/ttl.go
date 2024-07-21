package cache

import "time"

func Seconds(ttl int) time.Time {
	return time.Now().Add(time.Duration(ttl) * time.Second)
}

func Days(ttl int) time.Time {
	return time.Now().Add(time.Duration(ttl) * time.Second * 60 * 60 * 24)
}
