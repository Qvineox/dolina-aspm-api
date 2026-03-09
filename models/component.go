package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Component struct {
	UUID uuid.UUID `gorm:"column:uuid; primaryKey; type:uuid; default:uuid_generate_v7()"`
	PURL string    `gorm:"column:purl; uniqueIndex; not null; comment:Unique component PURL"`

	ComponentType ComponentType `gorm:"column:component_type; type:component_type"`

	IsPublic bool `gorm:"column:is_public; not null; index"`

	//Defects []Defect `gorm:"foreignKey:ComponentPURL; references:PURL; constraint:OnUpdate:CASCADE, OnDelete:SET NULL;"`
	//
	//Applications []Application `gorm:"many2many:application_components"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type ComponentType int32

const (
	ComponentTypeUnknown ComponentType = iota

	ComponentTypeLanguagePackage
	ComponentTypeContainerImage
	ComponentTypeOSPackage
	ComponentTypeBundle
	ComponentTypeBinaryExecutable
	ComponentTypeMobileApplication
)
