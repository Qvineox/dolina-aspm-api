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

	Type     DefectType                        `gorm:"column:type; not null; type: defect_type; index"`
	Statuses datatypes.JSONSlice[DefectStatus] `gorm:"column:status_list; type: jsonb"` // direct, indirect statuses

	AppliedRiskScore float32 `gorm:"column:applied_risk_score; type:numeric(4,2); index"`
	CVSSScore        float32 `gorm:"column:cvss_score; type:numeric(4,2); index"`

	AssetUUID     uuid.UUID  `gorm:"column:asset_uuid; not null; type:uuid; index"`
	ComponentUUID *uuid.UUID `gorm:"column:component_uuid; type:uuid; index"`

	VulnerabilityID *string `gorm:"column:vulnerability_id; index"`

	CWEs datatypes.JSONSlice[uint32] `gorm:"column:cwe_list; type:jsonb"`

	// CodeFragments
	// Activity
	//
	Stage *DefectStage `gorm:"column:stage; type:defect_stage; not null; default:open"` // todo: add default

	ReferenceURLs datatypes.JSONSlice[string] `gorm:"column:reference_url_list; type:json"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type DefectType string

const (
	DefectTypeUnknown DefectType = "unknown"

	DefectTypeSAST    DefectType = "sast"
	DefectTypeSecrets DefectType = "secrets"

	DefectTypeDAST DefectType = "dast"
	DefectTypeMAST DefectType = "mast"

	DefectTypePentest DefectType = "pentest"

	DefectTypeSCA DefectType = "sca"

	DefectTypeIAC DefectType = "iac"

	DefectTypeArchitecture DefectType = "architecture"
	DefectTypeCodeReview   DefectType = "code_review"
)

var DefectTypes = [...]DefectType{
	DefectTypeUnknown,

	DefectTypeSAST,
	DefectTypeSecrets,
	DefectTypeDAST,
	DefectTypeMAST,
	DefectTypePentest,
	DefectTypeSCA,
	DefectTypeIAC,
	DefectTypeArchitecture,
	DefectTypeCodeReview,
}

var DefectTypeEnums = map[DefectType]int32{
	DefectTypeUnknown: 0,

	DefectTypeSAST:         1,
	DefectTypeSecrets:      2,
	DefectTypeDAST:         3,
	DefectTypeMAST:         4,
	DefectTypePentest:      5,
	DefectTypeSCA:          6,
	DefectTypeIAC:          7,
	DefectTypeArchitecture: 8,
	DefectTypeCodeReview:   9,
}

type DefectStatus string

const (
	DefectStatusUnknown DefectStatus = "unknown"

	DefectStatusIndirect DefectStatus = "indirect"
	DefectStatusDirect   DefectStatus = "direct"

	DefectStatusDevOnly DefectStatus = "dev_only"
	DefectStatusCIOnly  DefectStatus = "ci_only"

	DefectStatusFixedByUpdate DefectStatus = "fixed_by_update"
	DefectStatusWillNotFix    DefectStatus = "will_not_fix"
	DefectStatusEndOfLife     DefectStatus = "end_of_life"

	DefectStatusHasExploit DefectStatus = "has_exploit"
)

var DefectsStatuses = [...]DefectStatus{
	DefectStatusUnknown,

	DefectStatusIndirect,
	DefectStatusDirect,
	DefectStatusDevOnly,
	DefectStatusCIOnly,
	DefectStatusFixedByUpdate,
	DefectStatusWillNotFix,
	DefectStatusEndOfLife,
	DefectStatusHasExploit,
}

var DefectStatusEnums = map[DefectStatus]int32{
	DefectStatusUnknown: 0,

	DefectStatusIndirect:      1,
	DefectStatusDirect:        2,
	DefectStatusDevOnly:       3,
	DefectStatusCIOnly:        4,
	DefectStatusFixedByUpdate: 5,
	DefectStatusWillNotFix:    6,
	DefectStatusEndOfLife:     7,
	DefectStatusHasExploit:    8,
}

type DefectStage string

const (
	DefectStage_Unknown DefectStage = "unknown"

	DefectStage_Open     DefectStage = "open"
	DefectStage_Assigned DefectStage = "assigned"

	DefectStage_UnderInvestigation      DefectStage = "under_investigation"
	DefectStage_RequestedAdditionalInfo DefectStage = "requested_additional_info"

	DefectStage_Confirmed DefectStage = "confirmed"

	DefectStage_Patching DefectStage = "patching"

	DefectStage_RiskAccepted DefectStage = "risk_accepted"
	DefectStage_Fixed        DefectStage = "fixed"

	DefectStage_Closed DefectStage = "closed"
)

var DefectsStages = [...]DefectStage{
	DefectStage_Unknown,

	DefectStage_Open,
	DefectStage_Assigned,
	DefectStage_UnderInvestigation,
	DefectStage_RequestedAdditionalInfo,
	DefectStage_Confirmed,
	DefectStage_Patching,
	DefectStage_RiskAccepted,
	DefectStage_Fixed,
	DefectStage_Closed,
}

var DefectsStagesEnums = map[DefectStage]int32{
	DefectStage_Unknown: 0,

	DefectStage_Open:                    1,
	DefectStage_Assigned:                2,
	DefectStage_UnderInvestigation:      3,
	DefectStage_RequestedAdditionalInfo: 4,
	DefectStage_Confirmed:               5,
	DefectStage_Patching:                6,
	DefectStage_RiskAccepted:            7,
	DefectStage_Fixed:                   8,
	DefectStage_Closed:                  9,
}
