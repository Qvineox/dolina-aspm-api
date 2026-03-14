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

	var _ uuid.UUID

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

		//componentUUID = component.UUID
	})
}
