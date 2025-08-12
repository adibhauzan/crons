package configs

import (
	"time"

	"github.com/robfig/cron/v3"
)

func NewCronConfig() *cron.Cron {
	location, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		location = time.Local
	}

	return cron.New(
		cron.WithLocation(location),
		cron.WithSeconds(),
		cron.WithChain(
			cron.Recover(cron.DefaultLogger),
		),
	)
}
