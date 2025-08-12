package repositories

import (
	"context"

	"github.com/adibhauzan/crons/internal/domain/entity"
	"gorm.io/gorm"
)

type documentsRepository struct {
	db *gorm.DB
}

type DocumentsRepository interface {
	Create(ctx context.Context, data *entity.Document) error
	GetAll(ctx context.Context, filter, text, sortBy, sortDirection string, page, limit int) ([]entity.Document, error)
	GetByID(ctx context.Context, id int) (*entity.Document, error)
	Update(ctx context.Context, data *entity.Document) error
	Delete(ctx context.Context, id int, user string) error
	GetNormal(ctx context.Context, filter, text, sortBy, sortDirection string, page, limit int) ([]entity.Document, error)
	GetClaim(ctx context.Context, filter, text, sortBy, sortDirection string, page, limit int) ([]entity.Document, error)
	GetRefund(ctx context.Context, filter, text, sortBy, sortDirection string, page, limit int) ([]entity.Document, error)
}

func NewDocumentsRepository(db *gorm.DB) *documentsRepository {
	return &documentsRepository{
		db: db,
	}
}

func (repo *documentsRepository) Create(ctx context.Context, data *entity.Document) error {
	err := repo.db.WithContext(ctx).Exec(`EXEC sp_mst_documents_insert @Name = ?, @User = ?, @isClaim =?, @isRefund = ?`, data.Name, data.CreatedBy, data.IsClaim, data.IsRefund).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *documentsRepository) GetAll(ctx context.Context, filter, text, sortBy, sortDirection string, page, limit int) ([]entity.Document, error) {
	var documents []entity.Document
	err := repo.db.WithContext(ctx).Raw(`EXEC sp_mst_documents_get @Filter = ?, @Text = ?, @SortBy = ?, @SortDirection = ?, @PageNumber = ?, @PageSize = ?`, filter, text, sortBy, sortDirection, page, limit).Scan(&documents).Error
	if err != nil {
		return nil, err
	}
	return documents, nil
}

func (repo *documentsRepository) GetByID(ctx context.Context, id int) (*entity.Document, error) {
	var document entity.Document
	err := repo.db.WithContext(ctx).Raw(`EXEC sp_mst_documents_get_by_id @ID = ?`, id).Scan(&document).Error
	if err != nil {
		return nil, err
	}
	return &document, nil
}

func (repo *documentsRepository) Update(ctx context.Context, data *entity.Document) error {
	err := repo.db.WithContext(ctx).Exec(`EXEC sp_mst_documents_update @ID = ?, @Name = ?, @User = ?, @isClaim =?, @isRefund = ?`, data.ID, data.Name, data.UpdatedBy, data.IsClaim, data.IsRefund).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *documentsRepository) Delete(ctx context.Context, id int, user string) error {
	err := repo.db.WithContext(ctx).Exec(`EXEC sp_mst_documents_delete @ID = ?, @User = ?`, id, user).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *documentsRepository) GetNormal(ctx context.Context, filter, text, sortBy, sortDirection string, page, limit int) ([]entity.Document, error) {
	var documents []entity.Document
	err := repo.db.WithContext(ctx).Raw(`EXEC sp_mst_documents_get_normal @Filter = ?, @Text = ?, @SortBy = ?, @SortDirection = ?, @PageNumber = ?, @PageSize = ?`, filter, text, sortBy, sortDirection, page, limit).Scan(&documents).Error
	if err != nil {
		return nil, err
	}
	return documents, nil
}

func (repo *documentsRepository) GetClaim(ctx context.Context, filter, text, sortBy, sortDirection string, page, limit int) ([]entity.Document, error) {
	var documents []entity.Document
	err := repo.db.WithContext(ctx).Raw(`EXEC sp_mst_documents_get_claim @Filter = ?, @Text = ?, @SortBy = ?, @SortDirection = ?, @PageNumber = ?, @PageSize = ?`, filter, text, sortBy, sortDirection, page, limit).Scan(&documents).Error
	if err != nil {
		return nil, err
	}
	return documents, nil
}

func (repo *documentsRepository) GetRefund(ctx context.Context, filter, text, sortBy, sortDirection string, page, limit int) ([]entity.Document, error) {
	var documents []entity.Document
	err := repo.db.WithContext(ctx).Raw(`EXEC sp_mst_documents_get_refund @Filter = ?, @Text = ?, @SortBy = ?, @SortDirection = ?, @PageNumber = ?, @PageSize = ?`, filter, text, sortBy, sortDirection, page, limit).Scan(&documents).Error
	if err != nil {
		return nil, err
	}
	return documents, nil
}
