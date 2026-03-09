package tests

import (
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.domsnail.ru/dolina/dolina-aspm-api/models"
	"gitlab.domsnail.ru/dolina/dolina-aspm-api/pkg/cfg"
	"gitlab.domsnail.ru/dolina/dolina-aspm-api/pkg/db"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func TestMigrations(t *testing.T) {
	config, err := cfg.NewConfigFromFile("config.yml")
	require.NoError(t, err)

	var orm *gorm.DB
	orm, err = gorm.Open(postgres.Open(config.Database.Postgres()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.LogLevel(config.Logging.Level)),
	})

	require.NoError(t, err)
	require.NotNil(t, orm)

	orm.AllowGlobalUpdate = true
	err = orm.Exec("DROP SCHEMA public CASCADE; CREATE SCHEMA public;").Error
	require.NoError(t, err)

	err = orm.Exec("CREATE EXTENSION IF NOT EXISTS pg_uuidv7;").Error
	require.NoError(t, err)

	t.Run("types migration test", func(t *testing.T) {
		MigrateTypes(t, orm)
	})

	t.Run("schema migration test", func(t *testing.T) {
		AutoMigrate(t, orm)
	})
}

func MigrateTypes(t *testing.T, orm *gorm.DB) {
	var err error

	err = db.MigrateComponentTypes(orm)
	require.NoError(t, err)

	err = db.MigrateAssetTypes(orm)
	require.NoError(t, err)

	err = db.MigrateDefectTypes(orm)
	require.NoError(t, err)

	err = db.MigrateDefectStatuses(orm)
	require.NoError(t, err)

	err = db.MigrateDefectStages(orm)
	require.NoError(t, err)

	err = db.MigrateCVSSSeverities(orm)
	require.NoError(t, err)
}

func AutoMigrate(t *testing.T, orm *gorm.DB) {
	var err error

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

	require.NoError(t, err)
}
