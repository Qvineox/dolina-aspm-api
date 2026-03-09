package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type ApplicationAsset struct {
	UUID     uuid.UUID `gorm:"column:uuid; primaryKey; type:uuid; default:uuid_generate_v7()"`
	SUID     string    `gorm:"column:suid; uniqueIndex:uidx_asset_ref; not null"`
	Revision string    `gorm:"column:revision; uniqueIndex:uidx_asset_ref; not null"` // e.g. ref, tag, branch

	Name        string `gorm:"column:name; index; not null"`
	Description string `gorm:"column:description"`
	URL         string `gorm:"column:url"`

	AssetType AssetType `gorm:"column:type; type:asset_type"`

	Labels datatypes.JSONSlice[string] `gorm:"column:labels; type:jsonb"`

	ComponentPURL *string   `gorm:"column:component_purl; uniqueIndex"`
	Component     Component `gorm:"foreignKey:PURL; references:ComponentPURL"`

	// Defects []Defect `gorm:"foreignKey:ApplicationID; references:ID; constraint:OnUpdate:CASCADE, OnDelete:SET NULL;"`

	Components []Component `gorm:"many2many:asset_components"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (a ApplicationAsset) TableName() string {
	return "application_assets"
}

func (a ApplicationAsset) Validate() bool {
	return true
}

type AssetType uint32

const (
	AssetType_Unknown AssetType = iota

	AssetType_Repository
	AssetType_Image
	AssetType_Executable
)

// todo: db.SetupJoinTable(&Person{}, "Addresses", &PersonAddress{})

// AssetComponent is a junction table
type AssetComponent struct {
	AssetID       uint32    `gorm:"column:asset_id; primaryKey"`
	ComponentUUID uuid.UUID `gorm:"column:component_uuid; type: uuid; primaryKey"`
}

func (a AssetComponent) TableName() string {
	return "application_assets_components"
}
