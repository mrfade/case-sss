package services

import (
	"context"

	"github.com/mrfade/case-sss/internal/core/models"
	"github.com/mrfade/case-sss/internal/core/ports"
	"github.com/mrfade/case-sss/pkg/errors"
	"github.com/mrfade/case-sss/pkg/request"
	"github.com/mrfade/case-sss/pkg/scorer"
)

type ContentService struct {
	repo      ports.ContentRepository
	providers []ports.ContentProvider
}

func NewContentService(
	repo ports.ContentRepository,
	providers ...ports.ContentProvider,
) ports.ContentService {
	return &ContentService{
		repo,
		providers,
	}
}

func (service *ContentService) FindByID(ctx context.Context, id int64) (*models.Content, error) {
	return service.repo.FindByID(ctx, id)
}

func (service *ContentService) Create(ctx context.Context, content *models.Content) (*models.Content, error) {
	return service.repo.Create(ctx, content)
}

func (service *ContentService) Update(ctx context.Context, content *models.Content) (*models.Content, error) {
	return service.repo.Update(ctx, content)
}

func (service *ContentService) Delete(ctx context.Context, id int64) error {
	return service.repo.Delete(ctx, id)
}

func (service *ContentService) FindAll(ctx context.Context, request *request.Request) ([]*models.Content, int64, error) {
	return service.repo.FindAll(ctx, request)
}

func (service *ContentService) SyncContents(ctx context.Context, scorer scorer.Scorer) error {
	for _, provider := range service.providers {
		contents, err := provider.FetchContents(ctx)
		if err != nil {
			return err
		}

		for _, content := range contents {
			existingContent, err := service.repo.FindByProviderID(ctx, content.Provider, content.ProviderID)
			if err != nil {
				if err == errors.ErrNotFound {
					content.Score = scorer.Score(content)
					_, createErr := service.repo.Create(ctx, content)
					if createErr != nil {
						return createErr
					}
				}
			} else {
				content.ID = existingContent.ID
				content.Score = scorer.Score(content)
				_, updateErr := service.repo.Update(ctx, content)
				if updateErr != nil {
					return updateErr
				}
			}
		}
	}

	return nil
}
