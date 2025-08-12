package crons

import (
	"github.com/adibhauzan/crons/internal/handlers"
)

type CronConfig struct {
	ClaimBlasting handlers.ClaimBlastingHandler
}

func (c *CronConfig) Setup() {
	c.SetupCron()
}

func (c *CronConfig) SetupCron() {
	c.ClaimBlasting.RegisterJob()
}
