package postgre

import (
	"context"

	"github.com/mrfade/case-sss/internal/core/models"
	"github.com/mrfade/case-sss/pkg/request"
)

type ContentRepository struct {
	db *DB
}

func NewContentRepository(db *DB) *ContentRepository {
	return &ContentRepository{
		db,
	}
}

func (repo *ContentRepository) FindByID(ctx context.Context, id int64) (*models.Content, error) {
	var content models.Content
	result := repo.db.DB.WithContext(ctx).First(&content, id)
	if result.Error != nil {
		return nil, FromGormError(result.Error)
	}

	return &content, nil
}

func (repo *ContentRepository) Create(ctx context.Context, content *models.Content) (*models.Content, error) {
	result := repo.db.DB.WithContext(ctx).Create(content)
	if result.Error != nil {
		return nil, FromGormError(result.Error)
	}

	return content, nil
}

func (repo *ContentRepository) Update(ctx context.Context, content *models.Content) (*models.Content, error) {
	result := repo.db.DB.WithContext(ctx).Save(content)
	if result.Error != nil {
		return nil, FromGormError(result.Error)
	}

	return content, nil
}

func (repo *ContentRepository) Delete(ctx context.Context, id int64) error {
	result := repo.db.DB.WithContext(ctx).Delete(&models.Content{}, id)
	if result.Error != nil {
		return FromGormError(result.Error)
	}

	return nil
}

func (repo *ContentRepository) FindByProviderID(ctx context.Context, provider, providerID string) (*models.Content, error) {
	var content models.Content
	result := repo.db.DB.WithContext(ctx).Where("provider = ? AND provider_id = ?", provider, providerID).First(&content)
	if result.Error != nil {
		return nil, FromGormError(result.Error)
	}

	return &content, nil
}

func (repo *ContentRepository) FindAll(ctx context.Context, request *request.Request) ([]*models.Content, error) {
	var contents []*models.Content
	query := repo.db.DB.WithContext(ctx).Model(&models.Content{})

	if request.PageNumber > 0 && request.PageSize > 0 {
		query = query.Offset((request.PageNumber - 1) * request.PageSize).Limit(request.PageSize)
	}

	if err := query.Find(&contents).Error; err != nil {
		return nil, FromGormError(err)
	}

	return contents, nil
}
