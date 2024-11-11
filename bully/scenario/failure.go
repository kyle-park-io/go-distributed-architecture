package scenario

import (
	"math"
	"time"
)

func NoResponse() {
	time.Sleep(time.Duration(math.MaxInt64))
}
