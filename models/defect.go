package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type AssetDefect struct {
	UUID uuid.UUID `gorm:"column:uuid; primaryKey; type:uuid; default:uuid_generate_v7()"`

	Title       string `gorm:"column:title; not null"` // todo: add tsvector index
	Description string `gorm:"column:description"`

	ApplicationID uint32     `gorm:"column:application_id; index"`
	ComponentUUID *uuid.UUID `gorm:"column:component_uuid; index"`

	Type     DefectType                        `gorm:"column:type; not null; type: defect_type; index"`
	Statuses datatypes.JSONSlice[DefectStatus] `gorm:"column:status_list; type: jsonb"` // direct, indirect statuses

	AppliedRiskScore float32 `gorm:"column:applied_risk_score; type:numeric(4,2); index"`
	CVSSScore        float32 `gorm:"column:applied_risk_score; type:numeric(4,2); index"`

	Vulnerability   Vulnerability `gorm:"foreignKey:ID; references:VulnerabilityID; index"`
	VulnerabilityID *string

	CWEs datatypes.JSONSlice[uint32] `gorm:"column:cwe_list; type:jsonb"`

	// CodeFragments
	// Activity

	Stage DefectStage `gorm:"column:stage; type:defect_stage; not null"` // todo: add default

	ReferenceURLs datatypes.JSONSlice[string] `gorm:"column:reference_url_list; type:json"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type DefectType int32

const (
	DefectTypeUnknown DefectType = iota

	DefectTypeSAST
	DefectTypeSecrets

	DefectTypeDAST
	DefectTypeMAST

	DefectTypePentest

	DefectTypeSCA

	DefectTypeIAC

	DefectTypeArchitecture
	DefectTypeCodeReview
)

type DefectStatus int32

const (
	DefectStatusUnknown DefectStatus = iota

	DefectStatusIndirect
	DefectStatusDirect

	DefectStatusDevOnly
	DefectStatusCIOnly

	DefectStatusFixedByUpdate
	DefectStatusWillNotFix
	DefectStatusEndOfLife

	DefectStatusHasExploit
)

type DefectStage int32

const (
	DefectStage_Unknown DefectStatus = iota

	DefectStage_Open
	DefectStage_Assigned

	DefectStage_UnderInvestigation
	DefectStage_RequestedAdditionalInfo

	DefectStage_Confirmed

	DefectStage_Patching

	DefectStage_RiskAccepted
	DefectStage_Fixed

	DefectStageClosed
)
