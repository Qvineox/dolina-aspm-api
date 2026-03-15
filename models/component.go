package models

import (
	"database/sql/driver"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/package-url/packageurl-go"
	"gitlab.domsnail.ru/dolina/dolina-aspm-api/pkg/utils"
	"gorm.io/gorm"
)

type Component struct {
	UUID uuid.UUID `gorm:"column:uuid; primaryKey; type:uuid; default:uuid_generate_v7()" validate:"required,uuid"`
	PURL *PURL     `gorm:"column:purl; uniqueIndex:uidx_component_purl; comment:Unique component package URL"`

	Name    string `gorm:"column:name; not null" validate:"required,min=1"`
	Version string `gorm:"column:version"`

	ComponentType ComponentType `gorm:"column:component_type; type:component_type; index" validate:"required"`

	IsPublic bool `gorm:"column:is_public; not null; index"`

	//Defects []Defect `gorm:"foreignKey:ComponentPURL; references:PURL; constraint:OnUpdate:CASCADE, OnDelete:SET NULL;"`

	Applications      []Application      `gorm:"foreignKey:ComponentPURL; references:PURL; constraint:OnUpdate:CASCADE, OnDelete:SET NULL"`
	ApplicationAssets []ApplicationAsset `gorm:"many2many:asset_components" validate:"omitempty,dive"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (c *Component) GetUUID() uuid.UUID {
	return c.UUID
}

func (c *Component) SetAssets(assets []ApplicationAsset) {
	c.ApplicationAssets = utils.UniqueByUUID(assets)
}

func (c *Component) BeforeCreate(tx *gorm.DB) error {
	err := validate.Struct(c)
	if err != nil {
		return fmt.Errorf("validation error: %w", err)
	}

	return nil
}

func (c *Component) BeforeUpdate(tx *gorm.DB) error {
	err := validate.Struct(c)
	if err != nil {
		return fmt.Errorf("validation error: %w", err)
	}

	return nil
}

func (c *Component) BeforeDelete(tx *gorm.DB) error {
	err := tx.Model(c).Association("ApplicationAssets").Unscoped().Clear()
	if err != nil {
		return fmt.Errorf("association delete error: %w", err)
	}

	return nil
}

func (c *Component) TableName() string {
	return "components"
}

type ComponentType string

const (
	ComponentTypeUnknown ComponentType = "unknown"

	ComponentTypeLanguagePackage   ComponentType = "language_package"
	ComponentTypeContainerImage    ComponentType = "container_image"
	ComponentTypeOSPackage         ComponentType = "os_package"
	ComponentTypeBundle            ComponentType = "bundle"
	ComponentTypeBinaryExecutable  ComponentType = "binary_executable"
	ComponentTypeMobileApplication ComponentType = "mobile_application"
)

var ComponentTypes = [...]ComponentType{
	ComponentTypeUnknown,
	ComponentTypeLanguagePackage,
	ComponentTypeContainerImage,
	ComponentTypeOSPackage,
	ComponentTypeBundle,
	ComponentTypeBinaryExecutable,
	ComponentTypeMobileApplication,
}

var ComponentTypesEnums = map[ComponentType]int32{
	ComponentTypeUnknown: 0,

	ComponentTypeLanguagePackage:   1,
	ComponentTypeContainerImage:    2,
	ComponentTypeOSPackage:         3,
	ComponentTypeBundle:            4,
	ComponentTypeBinaryExecutable:  5,
	ComponentTypeMobileApplication: 6,
}

type ComponentOptions struct {
	PURL    string
	Name    string
	Version string

	IsPublic bool

	ComponentType ComponentType
}

func NewComponent(opts ComponentOptions) (*Component, error) {
	uid, err := uuid.NewV7()
	if err != nil {
		return nil, fmt.Errorf("error generating uuid: %w", err)
	}

	component := &Component{
		Name:              opts.Name,
		Version:           opts.Version,
		UUID:              uid,
		ComponentType:     opts.ComponentType,
		IsPublic:          opts.IsPublic,
		ApplicationAssets: []ApplicationAsset{},
	}

	if opts.PURL != "" {
		purl, err := packageurl.FromString(opts.PURL)
		if err != nil {
			return nil, fmt.Errorf("error parsing purl: %w", err)
		}

		component.PURL = &PURL{purl}
	}

	err = validate.Struct(component)
	if err != nil {
		return nil, fmt.Errorf("validation error: %w", err)
	}

	return component, nil
}

type PURL struct {
	packageurl.PackageURL
}

func (p *PURL) Value() (driver.Value, error) {
	if len(p.PackageURL.Type) == 0 {
		return nil, nil
	}

	return p.ToString(), nil // e.g. "pkg:maven/org.apache.commons/io@1.3.4"
}

func (p *PURL) Scan(value any) error {
	if value == nil {
		p.PackageURL = packageurl.PackageURL{}
		return nil
	}

	var s string

	switch t := value.(type) {
	case []byte:
		s = string(value.([]byte))
	case string:
		s = value.(string)
	default:
		return fmt.Errorf("unsupported type %T", t)
	}

	if s == "" {
		p.PackageURL = packageurl.PackageURL{}
		return nil
	}

	var err error
	p.PackageURL, err = packageurl.FromString(s)
	if err != nil {
		return fmt.Errorf("package url malformed: %w", err)
	}

	return nil
}
