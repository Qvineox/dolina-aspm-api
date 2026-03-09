package tests

import (
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.domsnail.ru/dolina/dolina-aspm-api/models"
	"gitlab.domsnail.ru/dolina/dolina-aspm-api/pkg/cfg"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func TestAssetsModel(t *testing.T) {
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

	MigrateTypes(t, orm)
	AutoMigrate(t, orm)

	var application *models.Application

	t.Run("application create test", func(t *testing.T) {
		var opts = models.ApplicationOptions{
			Name:          "test_app2",
			Description:   "test_app2 description",
			URL:           "https://test.example.com",
			Labels:        nil,
			ComponentPURL: nil,
		}

		application, err = models.NewApplication(opts)
		require.NoError(t, err)
		require.NotNil(t, application)

		require.EqualValues(t, "test_app2", application.SUID)

		err = orm.Create(&application).Error
		require.NoError(t, err)
	})

	t.Run("application assets create error test", func(t *testing.T) {
		var asset *models.ApplicationAsset

		asset, err = models.NewApplicationAsset(models.ApplicationAssetOptions{
			Revision:      "main",
			Description:   "test_app2_asset1 description",
			URL:           "",
			Labels:        nil,
			ComponentPURL: nil,
			ApplicationID: nil,
		})

		require.Error(t, err)
		require.Nil(t, asset)

		asset, err = models.NewApplicationAsset(models.ApplicationAssetOptions{
			Revision:      "main",
			Name:          "test_app2_asset1",
			Description:   "test_app2_asset1 description",
			URL:           "",
			Labels:        nil,
			ComponentPURL: nil,
			ApplicationID: nil,
		})

		require.Error(t, err)
		require.Nil(t, asset)

		asset, err = models.NewApplicationAsset(models.ApplicationAssetOptions{
			Revision:      "main",
			Name:          "test_app2_asset1",
			Description:   "test_app2_asset1 description",
			AssetType:     models.AssetType_Repository,
			URL:           "1231231dawefasfwear",
			Labels:        nil,
			ComponentPURL: nil,
			ApplicationID: nil,
		})

		require.Error(t, err)
		require.Nil(t, asset)
	})

	t.Run("application assets create test", func(t *testing.T) {
		var asset *models.ApplicationAsset

		asset, err = models.NewApplicationAsset(models.ApplicationAssetOptions{
			Revision:      "main",
			Name:          "test_app2_asset1",
			Description:   "test_app2_asset1 description",
			AssetType:     models.AssetType_Repository,
			URL:           "",
			Labels:        nil,
			ComponentPURL: nil,
			ApplicationID: nil,
		})

		require.NoError(t, err)
		require.NotNil(t, asset)
	})

}
