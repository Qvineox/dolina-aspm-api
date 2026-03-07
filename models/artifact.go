package models

import "github.com/google/uuid"

type Artifact struct {
	UUID uuid.UUID `gorm:"column:uuid; primaryKey; type:uuid; default:uuid_generate_v7()"`

	Name         string `gorm:"column:name; not null"`
	ArtifactPURL string `gorm:"column:purl; not null"` // todo: add tsvector index
}
