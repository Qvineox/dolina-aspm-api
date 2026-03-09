package db

import (
	"fmt"
	"strings"

	"gitlab.domsnail.ru/dolina/dolina-aspm-api/models"
	"gorm.io/gorm"
)

func MigrateComponentTypes(orm *gorm.DB) error {
	var types []string
	for _, t := range models.ComponentTypes {
		types = append(types, fmt.Sprintf("'%s'", t))
	}

	err := orm.Exec(fmt.Sprintf(`
			DO $$ BEGIN
				CREATE TYPE component_type AS ENUM (
					%s
				);
				EXCEPTION WHEN duplicate_object THEN null;
    		END $$;
		`, strings.Join(types, ","))).Error

	if err != nil {
		return fmt.Errorf("failed to migrate type 'component_type': %w", err)
	}

	return err
}

func MigrateAssetTypes(orm *gorm.DB) error {
	var types []string
	for _, t := range models.AssetTypes {
		types = append(types, fmt.Sprintf("'%s'", t))
	}

	err := orm.Exec(fmt.Sprintf(`
			DO $$ BEGIN
				CREATE TYPE asset_type AS ENUM (
					%s
				);
				EXCEPTION WHEN duplicate_object THEN null;
    		END $$;
		`, strings.Join(types, ","))).Error

	if err != nil {
		return fmt.Errorf("failed to migrate type 'asset_type': %w", err)
	}

	return err
}

func MigrateDefectTypes(orm *gorm.DB) error {
	var types []string
	for _, t := range models.DefectTypes {
		types = append(types, fmt.Sprintf("'%s'", t))
	}

	err := orm.Exec(fmt.Sprintf(`
			DO $$ BEGIN
				CREATE TYPE defect_type AS ENUM (
					%s
				);
				EXCEPTION WHEN duplicate_object THEN null;
    		END $$;
		`, strings.Join(types, ","))).Error

	if err != nil {
		return fmt.Errorf("failed to migrate type 'defect_type': %w", err)
	}

	return err
}

func MigrateDefectStatuses(orm *gorm.DB) error {
	var types []string
	for _, t := range models.DefectsStatuses {
		types = append(types, fmt.Sprintf("'%s'", t))
	}

	err := orm.Exec(fmt.Sprintf(`
			DO $$ BEGIN
				CREATE TYPE defect_status AS ENUM (
					%s
				);
				EXCEPTION WHEN duplicate_object THEN null;
    		END $$;
		`, strings.Join(types, ","))).Error

	if err != nil {
		return fmt.Errorf("failed to migrate type 'defect_status': %w", err)
	}

	return err
}

func MigrateDefectStages(orm *gorm.DB) error {
	var types []string
	for _, t := range models.DefectsStages {
		types = append(types, fmt.Sprintf("'%s'", t))
	}

	err := orm.Exec(fmt.Sprintf(`
			DO $$ BEGIN
				CREATE TYPE defect_stage AS ENUM (
					%s
				);
				EXCEPTION WHEN duplicate_object THEN null;
    		END $$;
		`, strings.Join(types, ","))).Error

	if err != nil {
		return fmt.Errorf("failed to migrate type 'defect_stage': %w", err)
	}

	return err
}

func MigrateCVSSSeverities(orm *gorm.DB) error {
	var types []string
	for _, t := range models.Severities {
		types = append(types, fmt.Sprintf("'%s'", t))
	}

	err := orm.Exec(fmt.Sprintf(`
			DO $$ BEGIN
				CREATE TYPE cvss_severity AS ENUM (
					%s
				);
				EXCEPTION WHEN duplicate_object THEN null;
    		END $$;
		`, strings.Join(types, ","))).Error

	if err != nil {
		return fmt.Errorf("failed to migrate type 'cvss_severity': %w", err)
	}

	return err
}
