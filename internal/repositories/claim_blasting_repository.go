package repositories

import (
	"context"

	"github.com/adibhauzan/crons/internal/domain/entity"
	"gorm.io/gorm"
)

type ClaimBlastingRepository interface {
	GetAllClaimToBlasting(ctx context.Context) ([]entity.Claim, error)
	ClaimGetDocuments(ctx context.Context, number string) ([]entity.ClaimDocuments, error)
	UpdateStatusBlasting(ctx context.Context, number string) error
}

type claimBlastingRepository struct {
	db *gorm.DB
}

func NewClaimBlastingRepository(db *gorm.DB) ClaimBlastingRepository {
	return &claimBlastingRepository{
		db: db,
	}
}

func (r *claimBlastingRepository) GetAllClaimToBlasting(ctx context.Context) ([]entity.Claim, error) {
	var datas []entity.Claim

	err := r.db.WithContext(ctx).Raw(`EXEC sp_trx_claims_get_for_blasting`).Scan(&datas).Error
	if err != nil {
		return nil, err
	}

	return datas, nil
}

func (r *claimBlastingRepository) UpdateStatusBlasting(ctx context.Context, number string) error {
	err := r.db.WithContext(ctx).Exec(`EXEC sp_trx_claim_update_blasting_status @Number = ?`, number).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *claimBlastingRepository) ClaimGetDocuments(ctx context.Context, number string) ([]entity.ClaimDocuments, error) {
	var claimDocuments []entity.ClaimDocuments
	err := r.db.WithContext(ctx).Raw(`EXEC sp_trx_claimdocuments_get @Claims_Number = ?`, number).Scan(&claimDocuments).Error
	if err != nil {
		return nil, err
	}
	return claimDocuments, nil
}
