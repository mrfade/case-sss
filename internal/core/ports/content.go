package ports

import (
	"context"

	"github.com/mrfade/case-sss/internal/core/models"
	"github.com/mrfade/case-sss/pkg/request"
	"github.com/mrfade/case-sss/pkg/scorer"
)

type ContentProvider interface {
	FetchContents(ctx context.Context) ([]*models.Content, error)
}

type ContentRepository interface {
	Create(ctx context.Context, content *models.Content) (*models.Content, error)
	Update(ctx context.Context, content *models.Content) (*models.Content, error)
	Delete(ctx context.Context, id int64) error
	FindByID(ctx context.Context, id int64) (*models.Content, error)
	FindByProviderID(ctx context.Context, provider, providerID string) (*models.Content, error)
	FindAll(ctx context.Context, request *request.Request) ([]*models.Content, error)
}

type ContentService interface {
	Create(ctx context.Context, content *models.Content) (*models.Content, error)
	Update(ctx context.Context, content *models.Content) (*models.Content, error)
	Delete(ctx context.Context, id int64) error
	FindByID(ctx context.Context, id int64) (*models.Content, error)
	FindAll(ctx context.Context, request *request.Request) ([]*models.Content, error)
	SyncContents(ctx context.Context, scorer scorer.Scorer) error
}
