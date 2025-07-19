package services

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"

	"github.com/mrfade/case-sss/internal/core/models"
	"github.com/mrfade/case-sss/internal/core/ports"
	"github.com/mrfade/case-sss/pkg/errors"
	"github.com/mrfade/case-sss/pkg/request"
	"github.com/mrfade/case-sss/pkg/scorer"
)

type ContentService struct {
	repo      ports.ContentRepository
	cacher    ports.Cacher
	providers []ports.ContentProvider
}

func NewContentService(
	repo ports.ContentRepository,
	cacher ports.Cacher,
	providers ...ports.ContentProvider,
) ports.ContentService {
	return &ContentService{
		repo,
		cacher,
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
	const cacheTTL = 10 * time.Minute

	type cacheData struct {
		Contents     []*models.Content
		TotalRecords int64
	}

	cacheKey := generateCacheKey(request)
	if contents, total, ok := service.getFromCache(cacheKey); ok {
		return contents, total, nil
	}

	// Simulate a delay to mimic processing time
	time.Sleep(time.Second)

	contents, totalRecords, err := service.repo.FindAll(ctx, request)
	if err != nil {
		return nil, 0, err
	}

	data := &cacheData{
		Contents:     contents,
		TotalRecords: totalRecords,
	}
	service.setToCache(cacheKey, data, cacheTTL)

	return contents, totalRecords, nil
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

func generateCacheKey(request *request.Request) string {
	hash := sha256.Sum256([]byte(request.String()))
	return fmt.Sprintf("contents-%x", hash)
}

func (service *ContentService) getFromCache(key string) ([]*models.Content, int64, bool) {
	cached, err := service.cacher.Get(key)
	if err != nil {
		return nil, 0, false
	}

	var data struct {
		Contents     []*models.Content
		TotalRecords int64
	}

	str, ok := cached.(string)
	if !ok {
		return nil, 0, false
	}

	if err := json.Unmarshal([]byte(str), &data); err != nil {
		return nil, 0, false
	}

	return data.Contents, data.TotalRecords, true
}

func (service *ContentService) setToCache(key string, data any, ttl time.Duration) {
	value, err := json.Marshal(data)
	if err != nil {
		return
	}

	service.cacher.Set(key, value, ttl)
}
