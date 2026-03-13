package tests

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMigrations(t *testing.T) {
	t.Run("types migration test", func(t *testing.T) {
		require.NoError(t, migrateTypes())
	})

	t.Run("schema migration test", func(t *testing.T) {
		require.NoError(t, autoMigrate())
	})

	orm.AllowGlobalUpdate = false
}
