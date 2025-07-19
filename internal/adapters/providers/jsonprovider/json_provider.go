package jsonprovider

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/mrfade/case-sss/internal/core/models"
	"github.com/mrfade/case-sss/pkg/errors"
)

type JSONProvider struct {
	EndpointURL string
}

func NewJSONProvider(endpoint string) *JSONProvider {
	return &JSONProvider{
		EndpointURL: endpoint,
	}
}

type jsonContent struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Type    string `json:"type"`
	Metrics struct {
		Views    int    `json:"views"`
		Likes    int    `json:"likes"`
		Duration string `json:"duration"`
	} `json:"metrics"`
	PublishedAt string   `json:"published_at"`
	Tags        []string `json:"tags"`
}

type jsonResponse struct {
	Contents []jsonContent `json:"contents"`
}

func (p *JSONProvider) FetchContents(ctx context.Context) ([]*models.Content, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, p.EndpointURL, nil)
	if err != nil {
		return nil, errors.ErrJSONProviderRequestFailed
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, errors.ErrJSONProviderRequestFailed
	}

	defer resp.Body.Close()

	var data jsonResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, errors.ErrJSONProviderDecodeFailed
	}

	var result []*models.Content
	for _, item := range data.Contents {
		pubAt, err := time.Parse(time.RFC3339, item.PublishedAt)
		if err != nil {
			continue
		}

		result = append(result, &models.Content{
			Provider:    "jsonprovider",
			ProviderID:  item.ID,
			Title:       item.Title,
			Type:        models.ContentType(item.Type),
			Tags:        item.Tags,
			Views:       item.Metrics.Views,
			Likes:       item.Metrics.Likes,
			PublishedAt: pubAt,
		})
	}

	return result, nil
}
