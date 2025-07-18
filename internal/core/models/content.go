package models

import "time"

type ContentType string

const (
	ContentTypeVideo ContentType = "video"
	ContentTypeText  ContentType = "article"
)

type Content struct {
	ID          string      `gorm:"primaryKey"`
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
