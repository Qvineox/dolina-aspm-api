package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Component struct {
	UUID uuid.UUID `gorm:"column:uuid; primaryKey; type:uuid; default:uuid_generate_v7()"`
	PURL string    `gorm:"column:purl; uniqueIndex:uidx_component_purl; not null; comment:Unique component package URL"`

	ComponentType ComponentType `gorm:"column:component_type; type:component_type; index"`

	IsPublic bool `gorm:"column:is_public; not null; index"`

	//Defects []Defect `gorm:"foreignKey:ComponentPURL; references:PURL; constraint:OnUpdate:CASCADE, OnDelete:SET NULL;"`
	//
	//Applications []Application `gorm:"many2many:application_components"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
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
