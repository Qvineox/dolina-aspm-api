package tests

import (
	"log/slog"

	"gitlab.domsnail.ru/dolina/dolina-aspm-api/models"
	"gitlab.domsnail.ru/dolina/dolina-aspm-api/pkg/cfg"
	"gitlab.domsnail.ru/dolina/dolina-aspm-api/pkg/db"
	"gitlab.domsnail.ru/dolina/dolina-aspm-api/pkg/logs"
	"gorm.io/gorm"
)

var orm *gorm.DB

const configPath = "config.yml"

func init() {
	config, err := cfg.NewConfigFromFile(configPath)
	if err != nil {
		panic(err)
	}

	logger, err := logs.NewLogger(config.Logging)
	if err != nil {
		panic(err)
	}

	slog.SetDefault(logger)

	orm, err = db.NewOrmClient(config.Database)
	if err != nil {
		panic(err)
	}

	prepareDatabase()
}

func prepareDatabase() {
	var err error

	orm.AllowGlobalUpdate = true

	err = clearDatabase()
	if err != nil {
		panic(err)
	}

	err = orm.Exec("CREATE EXTENSION IF NOT EXISTS pg_uuidv7;").Error
	if err != nil {
		panic(err)
	}

	err = migrateTypes()
	if err != nil {
		panic(err)
	}

	err = autoMigrate()
	if err != nil {
		panic(err)
	}

	orm.AllowGlobalUpdate = false
}

func migrateTypes() (err error) {
	err = db.MigrateComponentTypes(orm)
	if err != nil {
		return err
	}

	err = db.MigrateAssetTypes(orm)
	if err != nil {
		return err
	}

	err = db.MigrateDefectTypes(orm)
	if err != nil {
		return err
	}

	err = db.MigrateDefectStatuses(orm)
	if err != nil {
		return err
	}

	err = db.MigrateDefectStages(orm)
	if err != nil {
		return err
	}

	err = db.MigrateCVSSSeverities(orm)
	if err != nil {
		return err
	}

	return nil
}

func autoMigrate() (err error) {
	err = orm.AutoMigrate(
		&models.Application{},
		&models.Component{},
		&models.ApplicationAsset{},
		&models.Vulnerability{},
		&models.AssetDefect{},
		&models.Vendor{},
		&models.VendorScoring{},
		&models.Exploit{},
	)

	return
}

func clearDatabase() (err error) {
	err = orm.Exec("DROP SCHEMA public CASCADE; CREATE SCHEMA public;").Error
	return
}
