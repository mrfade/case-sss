package xmlprovider

import (
	"context"
	"encoding/xml"
	"net/http"
	"time"

	"github.com/mrfade/case-sss/internal/core/models"
	"github.com/mrfade/case-sss/pkg/errors"
)

type XMLProvider struct {
	EndpointURL string
}

func NewXMLProvider(endpoint string) *XMLProvider {
	return &XMLProvider{
		EndpointURL: endpoint,
	}
}

type xmlFeed struct {
	Items []xmlItem `xml:"items>item"`
}

type xmlItem struct {
	ID         string   `xml:"id"`
	Headline   string   `xml:"headline"`
	Type       string   `xml:"type"`
	Stats      xmlStats `xml:"stats"`
	PubDate    string   `xml:"publication_date"`
	Categories []string `xml:"categories>category"`
}

type xmlStats struct {
	Views       int    `xml:"views"`
	Likes       int    `xml:"likes"`
	Duration    string `xml:"duration"`
	ReadingTime int    `xml:"reading_time"`
	Reactions   int    `xml:"reactions"`
	Comments    int    `xml:"comments"`
}

func (p *XMLProvider) FetchContents(ctx context.Context) ([]*models.Content, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, p.EndpointURL, nil)
	if err != nil {
		return nil, errors.ErrXMLProviderRequestFailed
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, errors.ErrXMLProviderRequestFailed
	}

	defer resp.Body.Close()

	var data xmlFeed
	if err := xml.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, errors.ErrXMLProviderDecodeFailed
	}

	var result []*models.Content
	for _, item := range data.Items {
		pubAt, err := time.Parse(time.DateOnly, item.PubDate)
		if err != nil {
			continue
		}

		result = append(result, &models.Content{
			Provider:    "xmlprovider",
			ProviderID:  item.ID,
			Title:       item.Headline,
			Type:        models.ContentType(item.Type),
			Tags:        item.Categories,
			Views:       item.Stats.Views,
			Likes:       item.Stats.Likes,
			ReadingTime: item.Stats.ReadingTime,
			Reactions:   item.Stats.Reactions,
			PublishedAt: pubAt,
		})
	}

	return result, nil
}
