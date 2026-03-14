package tests

import (
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.domsnail.ru/dolina/dolina-aspm-api/models"
)

func TestAssetsModel(t *testing.T) {
	var err error

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

	t.Run("application assets labels test", func(t *testing.T) {
		asset, err := models.NewApplicationAsset(models.ApplicationAssetOptions{
			Revision:    "main",
			Name:        "test_app2_asset1",
			Description: "test_app2_asset1 description",
			AssetType:   models.AssetType_Repository,
			Labels:      []string{"test_label1", "test_label2", "test_label2"},
		})

		require.NoError(t, err)
		require.Len(t, asset.Labels, 2)

		asset, err = models.NewApplicationAsset(models.ApplicationAssetOptions{
			Revision:    "main",
			Name:        "test_app2_asset1",
			Description: "test_app2_asset1 description",
			AssetType:   models.AssetType_Repository,
		})

		require.NoError(t, err)

		asset.SetLabels([]string{"test_label1", "test_label2"})
		require.Len(t, asset.Labels, 2)

		asset.SetLabels([]string{"test_label1"})
		require.Len(t, asset.Labels, 1)

		asset.SetLabels([]string{"test_label1", "test_label2", "test_label2"})
		require.Len(t, asset.Labels, 2)

		require.NoError(t, validate.Struct(asset))

		asset.Labels = []string{"test_label1", "test_label2", "test_label3", "test_label3"}
		require.Error(t, validate.Struct(asset))

		asset.SetLabels([]string{"test_label1", "test_label2", "test_label2"})
		require.Len(t, asset.Labels, 2)
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
			ApplicationID: &application.ID,
		})

		require.NoError(t, err)
		require.NotNil(t, asset)
		require.NoError(t, orm.Create(&asset).Error)

		asset, err = models.NewApplicationAsset(models.ApplicationAssetOptions{
			Revision:      "main",
			Name:          "test_app2_asset2",
			Description:   "test_app2_asset2 description",
			AssetType:     models.AssetType_Image,
			URL:           "https://test.example.com",
			Labels:        nil,
			ComponentPURL: nil,
			ApplicationID: &application.ID,
		})

		require.NoError(t, err)
		require.NotNil(t, asset)
		require.NoError(t, orm.Create(&asset).Error)

		var nonExistingApplicationID uint32 = 2
		asset, err = models.NewApplicationAsset(models.ApplicationAssetOptions{
			Revision:      "main",
			Name:          "test_app2_asset2",
			Description:   "test_app2_asset2 description",
			AssetType:     models.AssetType_Repository,
			URL:           "",
			Labels:        nil,
			ComponentPURL: nil,
			ApplicationID: &nonExistingApplicationID,
		})

		require.NoError(t, err)
		require.NotNil(t, asset)

		require.Error(t, orm.Create(&asset).Error)
		require.Error(t, orm.Save(&asset).Error)
	})

	t.Run("application assets query test", func(t *testing.T) {
		var application_ *models.Application

		require.NoError(t, orm.Preload("ApplicationAssets").First(&application_, "id = ?", application.ID).Error)
		require.Len(t, application_.Assets, 2)

		var assets []models.ApplicationAsset

		require.NoError(t, orm.Find(&assets, "application_id = ?", application.ID).Error)
		require.Len(t, assets, 2)
	})

	t.Run("application delete test", func(t *testing.T) {
		require.NoError(t, orm.Delete(&application).Error)

		var assets []models.ApplicationAsset

		require.NoError(t, orm.Find(&assets, "application_id = ?", application.ID).Error)
		require.Empty(t, assets)
	})
}
