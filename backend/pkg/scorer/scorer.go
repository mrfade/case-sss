package scorer

import (
	"time"

	"github.com/mrfade/case-sss/internal/core/models"
)

type Scorer interface {
	Score(content *models.Content) float64
}

type DefaultScorer struct{}

func (s DefaultScorer) Score(c *models.Content) float64 {
	base := 0.0
	multiplier := 1.0
	ageBonus := 0.0
	interactionBonus := 0.0

	switch c.Type {
	case models.ContentTypeVideo:
		base = float64(c.Views)/1000 + float64(c.Likes)/100
		multiplier = 1.5
		if c.Views > 0 {
			interactionBonus = (float64(c.Likes) / float64(c.Views)) * 10
		}
	case models.ContentTypeText:
		base = float64(c.ReadingTime) + float64(c.Reactions)/50
		multiplier = 1.0
		if c.ReadingTime > 0 {
			interactionBonus = (float64(c.Reactions) / float64(c.ReadingTime)) * 5
		}
	}

	daysAgo := int(time.Since(c.PublishedAt).Hours() / 24)
	switch {
	case daysAgo <= 7:
		ageBonus = 5
	case daysAgo <= 30:
		ageBonus = 3
	case daysAgo <= 90:
		ageBonus = 1
	default:
		ageBonus = 0
	}

	final := (base * multiplier) + ageBonus + interactionBonus
	c.Score = final
	return final
}
