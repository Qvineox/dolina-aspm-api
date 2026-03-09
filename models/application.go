package models

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Application struct {
	ID   uint32 `gorm:"column:id; primaryKey; autoIncrement"`
	SUID string `gorm:"column:suid; uniqueIndex; not null"`

	Name        string `gorm:"column:name; index; not null"`
	Description string `gorm:"column:description"`
	URL         string `gorm:"column:url"`

	Labels datatypes.JSONSlice[string] `gorm:"column:labels; type:jsonb"`

	ComponentPURL *string   `gorm:"column:component_purl; uniqueIndex"`
	Component     Component `gorm:"foreignKey:PURL; references:ComponentPURL"`

	// Defects []Defect `gorm:"foreignKey:ApplicationID; references:ID; constraint:OnUpdate:CASCADE, OnDelete:SET NULL;"`

	// RiskProfile // todo: add application risk profile

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (a Application) TableName() string {
	return "applications"
}

func (a Application) Validate() bool {
	return true
}
