package models

import "time"

type ContentType string

const (
	ContentTypeVideo ContentType = "video"
	ContentTypeText  ContentType = "article"
)

type Content struct {
	ID          int64       `gorm:"primaryKey,autoIncrement"`
	Provider    string      `gorm:"not null"`
	ProviderID  string      `gorm:"not null"`
	Title       string      `gorm:"not null"`
	Type        ContentType `gorm:"not null"`
	Tags        []string    `gorm:"serializer:json;null"`
	Views       int         `gorm:"default:0"`
	Likes       int         `gorm:"default:0"`
	ReadingTime int         `gorm:"default:0"`
	Reactions   int         `gorm:"default:0"`
	PublishedAt time.Time   `gorm:"not null"`
	Score       float64     `gorm:"default:0"`
}
