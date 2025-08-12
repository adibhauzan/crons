package repositories

import (
	"context"

	"github.com/adibhauzan/crons/internal/domain/entity"
	"gorm.io/gorm"
)

type ApplicationSettingRepository interface {
	Get(ctx context.Context) (*entity.ApplicationSetting, error)
}

type applicationSettingRepository struct {
	db *gorm.DB
}

func NewApplicationSettingRepository(db *gorm.DB) *applicationSettingRepository {
	return &applicationSettingRepository{
		db: db,
	}
}

func (repo *applicationSettingRepository) Get(ctx context.Context) (*entity.ApplicationSetting, error) {
	var setting entity.ApplicationSetting
	err := repo.db.WithContext(ctx).Raw(`EXEC sp_setting_get`).Scan(&setting).Error
	if err != nil {
		return nil, err
	}
	return &setting, nil
}
