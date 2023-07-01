package time

import (
	gotime "time"

	"github.com/undefinedlabs/go-mpatch"
)

func PatchTimeNow() {
	mpatch.PatchMethod(gotime.Now, func() gotime.Time {
		return gotime.Date(2023, 03, 06, 14, 27, 45, 3699162, gotime.UTC)
	})
}
