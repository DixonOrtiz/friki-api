package utils

import (
	"time"

	"github.com/undefinedlabs/go-mpatch"
)

func PatchTimeNow() {
	mpatch.PatchMethod(time.Now, func() time.Time {
		return time.Date(2023, 03, 06, 14, 27, 45, 3699162, time.UTC)
	})
}
