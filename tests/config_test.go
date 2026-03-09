package tests

import (
	"log/slog"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.domsnail.ru/dolina/dolina-aspm-api/pkg/cfg"
)

func TestConfig(t *testing.T) {
	t.Run("full static config test", func(t *testing.T) {
		const path = "mock_data/configs/full_config.yml"

		config, err := cfg.NewConfigFromFile(path)
		require.NoError(t, err)
		require.NotNil(t, config)

		require.EqualValues(t, "1.2.3.4", config.Database.Host)
		require.EqualValues(t, uint64(5432), config.Database.Port)
		require.EqualValues(t, "test_user1", config.Database.User)
		require.EqualValues(t, "test_pass2", config.Database.Pass)
		require.EqualValues(t, "Europe/Moscow", config.Database.Timezone)

		require.EqualValues(t, 4, config.Logging.Level)
	})

	t.Run("partial static config test with defaults", func(t *testing.T) {
		const path = "mock_data/configs/defaults_config.yml"

		config, err := cfg.NewConfigFromFile(path)
		require.NoError(t, err)
		require.NotNil(t, config)

		testDatabaseDefaults(t, config.Database)
		testLoggingDefaults(t, config.Logging)
	})

	t.Run("static config test with missing required values", func(t *testing.T) {
		const path = "mock_data/configs/invalid_config.yml"

		config, err := cfg.NewConfigFromFile(path)
		require.Error(t, err)
		require.NotNil(t, config)
	})

	t.Run("static config test with invalid logging values", func(t *testing.T) {
		const path = "mock_data/configs/invalid_logging_config.yml"

		config, err := cfg.NewConfigFromFile(path)
		require.Error(t, err)
		require.NotNil(t, config)
	})

	t.Run("static config test with no password", func(t *testing.T) {
		const path = "mock_data/configs/passwordless_config.yml"

		config, err := cfg.NewConfigFromFile(path)
		require.NoError(t, err)
		require.NotNil(t, config)
	})
}

func testDatabaseDefaults(t *testing.T, config cfg.DatabaseConfig) {
	require.EqualValues(t, "0.0.0.0", config.Host)
	require.EqualValues(t, uint64(5432), config.Port)
	require.EqualValues(t, "test_user", config.User)
	require.EqualValues(t, "test_pass", config.Pass)
	require.EqualValues(t, "Europe/Moscow", config.Timezone)
}

func testLoggingDefaults(t *testing.T, config cfg.LoggingConfig) {
	require.EqualValues(t, slog.LevelInfo, config.Level)
}
