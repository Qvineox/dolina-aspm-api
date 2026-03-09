package models

import (
	"fmt"
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Application struct {
	ID   uint32 `gorm:"column:id; primaryKey; autoIncrement"`
	SUID string `gorm:"column:suid; uniqueIndex:uidx_application_slug; not null; comment:String unique ID" validate:"required,min=4,max=16"`

	Name        string `gorm:"column:name; index; not null" validate:"required"`
	Description string `gorm:"column:description"`
	URL         string `gorm:"column:url" validate:"omitempty,url"`

	Labels datatypes.JSONSlice[string] `gorm:"column:labels; type:jsonb; default:'[]'"`

	ComponentPURL *string   `gorm:"column:component_purl; uniqueIndex:uidx_application_purl"`
	Component     Component `gorm:"foreignKey:PURL; references:ComponentPURL"`

	Assets []ApplicationAsset `gorm:"foreignKey:ApplicationID; references:ID; constraint:OnUpdate:CASCADE, OnDelete:SET NULL;"`

	// RiskProfile // todo: add application risk profile

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (a *Application) BeforeCreate(tx *gorm.DB) error {
	err := validate.Struct(a)
	if err != nil {
		return fmt.Errorf("validation error: %w", err)
	}

	return nil
}

func (a *Application) BeforeUpdate(tx *gorm.DB) error {
	err := validate.Struct(a)
	if err != nil {
		return fmt.Errorf("validation error: %w", err)
	}

	return nil
}

func (a *Application) TableName() string {
	return "applications"
}

type ApplicationOptions struct {
	SUID        string
	Name        string
	Description string

	URL    string
	Labels []string

	ComponentPURL *string
}

func NewApplication(opts ApplicationOptions) (*Application, error) {
	app := &Application{
		Name:          opts.Name,
		Description:   opts.Description,
		URL:           opts.URL,
		Labels:        opts.Labels,
		ComponentPURL: opts.ComponentPURL,
	}

	if opts.SUID != "" {
		app.SUID = opts.SUID
	} else {
		app.SUID = generateSUID(opts.Name)
	}

	err := validate.Struct(app)
	if err != nil {
		return nil, fmt.Errorf("validation error: %w", err)
	}

	return app, nil
}
