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

	Labels datatypes.JSONSlice[string] `gorm:"column:labels; type:json"`

	ComponentPURL *string   `gorm:"column:component_purl; uniqueIndex"`
	Component     Component `gorm:"foreignKey:PURL; references:ComponentPURL"`

	Defects []Defect `gorm:"foreignKey:ApplicationID; references:ID; constraint:OnUpdate:CASCADE, OnDelete:SET NULL;"`

	Components []Component `gorm:"many2many:application_components"`
	Artifacts  []Component `gorm:"many2many:application_components"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (a Application) Validate() bool {
	return true
}

// todo: db.SetupJoinTable(&Person{}, "Addresses", &PersonAddress{})

type ApplicationComponent struct {
	ApplicationID uint32 `gorm:"column:application_id; primaryKey"`
	ComponentID   uint32 `gorm:"column:component_id; primaryKey"`
}
