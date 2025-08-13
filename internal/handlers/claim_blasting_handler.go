// handlers/claim_blasting_handler.go
package handlers

import (
	"context"

	"github.com/adibhauzan/crons/internal/services"
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
)

type ClaimBlastingHandler interface {
	RegisterJob() error
}

type claimBlastingHandler struct {
	claimBlastingService services.ClaimBlastingService
	cron                 *cron.Cron
	logger               *logrus.Logger
}

func NewClaimBlastingHandler(claimBlastingService services.ClaimBlastingService, cron *cron.Cron, logger *logrus.Logger) *claimBlastingHandler {
	return &claimBlastingHandler{
		claimBlastingService: claimBlastingService,
		cron:                 cron,
		logger:               logger,
	}
}

func (h *claimBlastingHandler) RegisterJob() error {
	_, err := h.cron.AddFunc("0 0 20 * * *", func() {
		h.logger.Info("Starting Blasting Email Claim Cron")
		if err := h.claimBlastingService.BlastingEmailClaim(context.Background()); err != nil {
			h.logger.Errorf("Failed to blast email claims: %v", err)
		} else {
			h.logger.Info("Blasting Email Claim Cron completed successfully")
		}
	})
	return err
}
