package terrain

import (
	"time"
)

func GetSeed() int64 {
	return time.Now().UTC().UnixNano()
}