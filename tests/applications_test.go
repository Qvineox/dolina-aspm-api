package tests

import (
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.domsnail.ru/dolina/dolina-aspm-api/pkg/cfg"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func TestApplicationModel(t *testing.T) {
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

	t.Run("application create test", func(t *testing.T) {

	})
}
