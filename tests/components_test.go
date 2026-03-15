package tests

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"gitlab.domsnail.ru/dolina/dolina-aspm-api/models"
)

func TestComponentModel(t *testing.T) {
	t.Run("component create error test", func(t *testing.T) {
		component, err := models.NewComponent(models.ComponentOptions{})

		require.Error(t, err)
		require.Nil(t, component)

		component, err = models.NewComponent(models.ComponentOptions{
			Name:     "test_component1",
			IsPublic: false,
		})

		require.Error(t, err)
		require.Nil(t, component)

		component, err = models.NewComponent(models.ComponentOptions{
			Name:     "test_component2",
			IsPublic: false,
		})

		require.Error(t, err)
		require.Nil(t, component)

		component, err = models.NewComponent(models.ComponentOptions{
			PURL:     "not_valid_purl",
			Name:     "test_component3",
			IsPublic: false,
		})

		require.Error(t, err)
		require.Nil(t, component)
	})

	var componentUUID uuid.UUID

	t.Run("component create test", func(t *testing.T) {
		component, err := models.NewComponent(models.ComponentOptions{
			ComponentType: models.ComponentTypeLanguagePackage,
			Name:          "test_component1",
			IsPublic:      false,
		})

		require.NoError(t, err)
		require.NotNil(t, component)

		component, err = models.NewComponent(models.ComponentOptions{
			Name:          "test_component1",
			IsPublic:      false,
			ComponentType: models.ComponentTypeLanguagePackage,
		})

		require.NoError(t, err)
		require.NotNil(t, component)

		component, err = models.NewComponent(models.ComponentOptions{
			PURL:          "pkg:maven/org.apache.commons/io@1.3.4",
			Name:          "test_component1",
			IsPublic:      false,
			ComponentType: models.ComponentTypeLanguagePackage,
		})

		require.NoError(t, err)
		require.NotNil(t, component)

		require.EqualValues(t, "pkg:maven/org.apache.commons/io@1.3.4", component.PURL.String())
		require.EqualValues(t, "maven", component.PURL.Type)
		require.EqualValues(t, "org.apache.commons", component.PURL.Namespace)
		require.EqualValues(t, "io", component.PURL.Name)
		require.EqualValues(t, "1.3.4", component.PURL.Version)

		err = orm.Create(&component).Error

		require.NoError(t, err)
		require.NotZero(t, component.UUID.String())
		require.NotNil(t, component.CreatedAt)
		require.NotNil(t, component.UpdatedAt)

		componentUUID = component.UUID
	})

	t.Run("components query test", func(t *testing.T) {
		var component_ *models.Component

		require.NoError(t, orm.First(&component_, "uuid = ?", componentUUID).Error)

		require.EqualValues(t, "pkg:maven/org.apache.commons/io@1.3.4", component_.PURL.String())
		require.EqualValues(t, "maven", component_.PURL.Type)
		require.EqualValues(t, "org.apache.commons", component_.PURL.Namespace)
		require.EqualValues(t, "io", component_.PURL.Name)
		require.EqualValues(t, "1.3.4", component_.PURL.Version)

		require.NotZero(t, component_.UUID.String())
		require.NotNil(t, component_.CreatedAt)
		require.NotNil(t, component_.UpdatedAt)
	})

	var assetUUID uuid.UUID
	var assetSUID string

	t.Run("asset component creation test", func(t *testing.T) {
		asset, err := models.NewApplicationAsset(models.ApplicationAssetOptions{
			Revision:    "main",
			Name:        "test_asset3",
			Description: "test_asset3 description",
			AssetType:   models.AssetType_Repository,
		})

		require.NoError(t, err)
		require.NotNil(t, asset)
		require.NoError(t, orm.Create(&asset).Error)

		assetSUID = asset.SUID
		assetUUID = asset.UUID

		component, err := models.NewComponent(models.ComponentOptions{
			ComponentType: models.ComponentTypeLanguagePackage,
			Name:          "test_asset3_component1",
			IsPublic:      false,
		})

		require.NoError(t, err)
		require.NotNil(t, component)

		component.ApplicationAssets = append(component.ApplicationAssets, *asset)
		require.NoError(t, orm.Create(&component).Error)

		component, err = models.NewComponent(models.ComponentOptions{
			Name:          "test_asset3_component2",
			IsPublic:      false,
			ComponentType: models.ComponentTypeLanguagePackage,
		})

		require.NoError(t, err)
		require.NotNil(t, component)

		component.ApplicationAssets = append(component.ApplicationAssets, *asset)
		require.NoError(t, orm.Create(&component).Error)

		component, err = models.NewComponent(models.ComponentOptions{
			Name:          "test_asset3_component2",
			IsPublic:      false,
			ComponentType: models.ComponentTypeLanguagePackage,
			PURL:          "pkg:maven/org.apache.commons/test@1.4.5",
		})

		require.NoError(t, err)
		require.NotNil(t, component)

		component.ApplicationAssets = append(component.ApplicationAssets, *asset)
		require.NoError(t, orm.Create(&component).Error)
	})

	t.Run("asset components query test", func(t *testing.T) {
		var asset_ *models.ApplicationAsset

		require.NoError(t, orm.Preload("Components").First(&asset_, "suid = ?", assetSUID).Error)
		require.Len(t, asset_.Components, 3)
	})

	t.Run("asset delete test", func(t *testing.T) {
		require.NoError(t, orm.Where("suid = ?", assetSUID).Delete(&models.ApplicationAsset{}).Error)

		var components []models.ApplicationAsset

		err := orm.
			Model(&models.ApplicationAsset{}).
			Where("application_asset_uuid = ?", assetUUID).
			Association("Components").
			Find(&components)

		require.NoError(t, err)
		require.Empty(t, components)
	})
}
