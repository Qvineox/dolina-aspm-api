package models

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"gitlab.domsnail.ru/dolina/dolina-aspm-api/pkg/utils"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type ApplicationAsset struct {
	UUID     uuid.UUID `gorm:"column:uuid; primaryKey; type:uuid; default:uuid_generate_v7()" validate:"required,uuid"`
	SUID     string    `gorm:"column:suid; uniqueIndex:uidx_asset_ref; not null" validate:"required"`
	Revision string    `gorm:"column:revision; uniqueIndex:uidx_asset_ref; not null" validate:"required"` // e.g. ref, tag, branch

	ComponentPURL *string `gorm:"column:component_purl; uniqueIndex"`
	Name          string  `gorm:"column:name; index; not null" validate:"required"`
	Description   string  `gorm:"column:description"`
	URL           string  `gorm:"column:url" validate:"omitempty,url"`

	AssetType AssetType                   `gorm:"column:type; type:asset_type" validate:"required,oneof=repository image executable"`
	Labels    datatypes.JSONSlice[string] `gorm:"column:labels; type:jsonb; default:'[]'" validate:"omitempty,unique"`

	ApplicationID *uint32 `gorm:"column:application_id; index"`

	Defects    []AssetDefect `gorm:"foreignKey:AssetUUID; references:UUID; constraint:OnUpdate:CASCADE, OnDelete:SET NULL;"`
	Components []Component   `gorm:"many2many:asset_components"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (a ApplicationAsset) GetUUID() uuid.UUID {
	return a.UUID
}

func (a *ApplicationAsset) SetLabels(labels []string) {
	a.Labels = utils.Unique(labels)
}

func (a *ApplicationAsset) BeforeCreate(tx *gorm.DB) error {
	err := validate.Struct(a)
	if err != nil {
		return fmt.Errorf("validation error: %w", err)
	}

	return nil
}

func (a *ApplicationAsset) BeforeUpdate(tx *gorm.DB) error {
	err := validate.Struct(a)
	if err != nil {
		return fmt.Errorf("validation error: %w", err)
	}

	return nil
}

func (a *ApplicationAsset) TableName() string {
	return "application_assets"
}

type ApplicationAssetOptions struct {
	SUID     string
	Revision string

	Name        string
	Description string
	AssetType   AssetType

	URL    string
	Labels []string

	ComponentPURL *string
	ApplicationID *uint32
}

func NewApplicationAsset(opts ApplicationAssetOptions) (*ApplicationAsset, error) {
	uid, err := uuid.NewV7()
	if err != nil {
		return nil, fmt.Errorf("error generating uuid: %w", err)
	}

	app := &ApplicationAsset{
		UUID:          uid,
		Name:          opts.Name,
		Description:   opts.Description,
		Revision:      opts.Revision,
		ApplicationID: opts.ApplicationID,
		AssetType:     opts.AssetType,
		URL:           opts.URL,
		Labels:        utils.Unique(opts.Labels),
		ComponentPURL: opts.ComponentPURL,
	}

	if opts.SUID != "" {
		app.SUID = opts.SUID
	} else {
		app.SUID = generateSUID(opts.Name)
	}

	err = validate.Struct(app)
	if err != nil {
		return nil, fmt.Errorf("validation error: %w", err)
	}

	return app, nil
}

type AssetType string

const (
	AssetType_Unknown AssetType = "unknown"

	AssetType_Repository AssetType = "repository"
	AssetType_Image      AssetType = "image"
	AssetType_Executable AssetType = "executable"
)

var AssetTypes = [...]AssetType{
	AssetType_Unknown,
	AssetType_Repository,
	AssetType_Image,
	AssetType_Executable,
}

var AssetTypesEnums = map[AssetType]int32{
	AssetType_Unknown: 0,

	AssetType_Repository: 1,
	AssetType_Image:      2,
	AssetType_Executable: 3,
}

// todo: db.SetupJoinTable(&Person{}, "Addresses", &PersonAddress{})

// AssetComponent is a junction table
type AssetComponent struct {
	AssetID       uint32    `gorm:"column:asset_id; primaryKey"`
	ComponentUUID uuid.UUID `gorm:"column:component_uuid; type: uuid; primaryKey"`
}

func (a AssetComponent) TableName() string {
	return "asset_components"
}
