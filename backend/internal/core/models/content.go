package models

import "time"

type ContentType string

const (
	ContentTypeVideo ContentType = "video"
	ContentTypeText  ContentType = "article"
)

type Content struct {
	ID          int64       `gorm:"primaryKey,autoIncrement" json:"id"`
	Provider    string      `gorm:"not null" json:"provider"`
	ProviderID  string      `gorm:"not null" json:"provider_id"`
	Title       string      `gorm:"not null" json:"title"`
	Type        ContentType `gorm:"not null" json:"type"`
	Tags        []string    `gorm:"serializer:json;null" json:"tags"`
	Views       int         `gorm:"-:all" json:"-"`
	Likes       int         `gorm:"-:all" json:"-"`
	ReadingTime int         `gorm:"-:all" json:"-"`
	Reactions   int         `gorm:"-:all" json:"-"`
	PublishedAt time.Time   `gorm:"not null" json:"published_at"`
	Score       float64     `gorm:"default:0" json:"score"`
}
