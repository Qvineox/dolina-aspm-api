package tests

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/require"
	"gitlab.domsnail.ru/dolina/dolina-aspm-api/models"
)

var validate = validator.New()

func TestApplicationModel(t *testing.T) {
	var err error

	t.Run("application create error test", func(t *testing.T) {
		application, err := models.NewApplication(models.ApplicationOptions{
			Description:   "test_app1 description",
			URL:           "https://test.example.com",
			Labels:        nil,
			ComponentPURL: nil,
		})

		require.Error(t, err)
		require.Nil(t, application)

		err = orm.Create(&application).Error
		require.Error(t, err)

		application, err = models.NewApplication(models.ApplicationOptions{
			Name:          "tst",
			Description:   "test_app1 description",
			URL:           "https://test.example.com",
			Labels:        nil,
			ComponentPURL: nil,
		})

		require.Error(t, err)
		require.Nil(t, application)

		err = orm.Create(&application).Error
		require.Error(t, err)

		var opts = models.ApplicationOptions{
			Name:          "test_app1$F",
			Description:   "test_app1 description",
			URL:           "https://test.example.com",
			Labels:        nil,
			ComponentPURL: nil,
		}

		application, err = models.NewApplication(opts)
		require.NoError(t, err)
		require.NotNil(t, application)

		require.EqualValues(t, "test_app1_f", application.SUID)
	})

	var appID uint32
	var appSUID string

	t.Run("application create test", func(t *testing.T) {
		var opts = models.ApplicationOptions{
			Name:          "test_app1",
			Description:   "test_app1 description",
			URL:           "https://test.example.com",
			Labels:        nil,
			ComponentPURL: nil,
		}

		application, err := models.NewApplication(opts)
		require.NoError(t, err)
		require.NotNil(t, application)

		require.EqualValues(t, "test_app1", application.SUID)

		err = orm.Create(&application).Error
		require.NoError(t, err)

		require.EqualValues(t, "test_app1", application.SUID)
		require.NotZero(t, application.ID)
		require.NotNil(t, application.CreatedAt)
		require.NotNil(t, application.UpdatedAt)

		appID = application.ID
		appSUID = application.SUID
	})

	t.Run("application find test", func(t *testing.T) {
		app := &models.Application{}

		err = orm.First(&app, appID).Error
		require.NoError(t, err)
		require.NotNil(t, app)
		require.Equal(t, "test_app1", app.SUID)
		require.Equal(t, "test_app1", app.Name)

		err = orm.First(&app, "suid = ?", appSUID).Error
		require.NoError(t, err)
		require.NotNil(t, app)
		require.Equal(t, "test_app1", app.SUID)
		require.Equal(t, "test_app1", app.Name)
	})

	t.Run("application update test", func(t *testing.T) {
		app := &models.Application{}

		err = orm.First(&app, appID).Error
		require.NoError(t, err)
		require.NotNil(t, app)
		require.Equal(t, "test_app1", app.SUID)
		require.Equal(t, "test_app1", app.Name)

		updatedAt := app.UpdatedAt
		createdAt := app.CreatedAt

		app.Description = "test_app1 new description"

		err = orm.Save(&app).Error
		require.NoError(t, err)
		require.NotNil(t, app)

		err = orm.First(&app, appID).Error
		require.NoError(t, err)
		require.NotNil(t, app)
		require.Equal(t, "test_app1", app.SUID)
		require.Equal(t, "test_app1", app.Name)

		require.Equal(t, "test_app1 new description", app.Description)
		require.NotEqual(t, updatedAt, app.UpdatedAt)
		require.Equal(t, createdAt, app.CreatedAt)
	})

	t.Run("application labels test", func(t *testing.T) {
		app := &models.Application{}

		err = orm.First(&app, appID).Error
		require.NoError(t, err)
		require.NotNil(t, app)

		app.SetLabels([]string{"test_label1", "test_label2"})
		require.Len(t, app.Labels, 2)

		app.SetLabels([]string{"test_label1"})
		require.Len(t, app.Labels, 1)

		app.SetLabels([]string{"test_label1", "test_label2", "test_label2"})
		require.Len(t, app.Labels, 2)

		require.NoError(t, validate.Struct(app))

		app.Labels = []string{"test_label1", "test_label2", "test_label3", "test_label3"}
		require.Error(t, validate.Struct(app))

		app.SetLabels([]string{"test_label1", "test_label2", "test_label2"})
		require.Len(t, app.Labels, 2)

		err = orm.Save(&app).Error
		require.NoError(t, err)

		err = orm.First(&app, appID).Error
		require.NoError(t, err)
		require.Len(t, app.Labels, 2)
	})

	t.Run("application update error test", func(t *testing.T) {
		app := &models.Application{}

		err = orm.First(&app, appID).Error
		require.NoError(t, err)
		require.NotNil(t, app)
		require.Equal(t, "test_app1", app.SUID)
		require.Equal(t, "test_app1", app.Name)

		app.SUID = ""

		err = orm.Save(&app).Error
		require.Error(t, err)
	})
}
