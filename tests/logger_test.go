package tests

import (
	"context"
	"log/slog"
	"os"
	"path/filepath"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"gitlab.domsnail.ru/dolina/dolina-aspm-api/pkg/cfg"
	"gitlab.domsnail.ru/dolina/dolina-aspm-api/pkg/logs"
)

func TestLogger(t *testing.T) {
	dir := t.TempDir()
	filePath := filepath.Join(dir, "test.jsonl")

	t.Run("test invalid logging configs", func(t *testing.T) {
		hostname, _ := os.Hostname()

		opt := cfg.LogsConfig{
			Host:        hostname,
			Environment: "test1",
			Level:       int(slog.LevelInfo),
			AddSource:   false,
			Format:      "json",
			Filepath:    "",
		}

		var err error

		_, err = logs.NewLogger(opt)
		require.Error(t, err)

		opt = cfg.LogsConfig{
			Environment: "dev",
			Level:       int(slog.LevelInfo),
			AddSource:   false,
			Format:      "json",
			Filepath:    "",
		}

		_, err = logs.NewLogger(opt)
		require.Error(t, err)

		opt = cfg.LogsConfig{
			Environment: "test",
			Level:       -12,
			AddSource:   false,
			Format:      "json",
			Filepath:    "",
		}

		_, err = logs.NewLogger(opt)
		require.Error(t, err)

		opt = cfg.LogsConfig{
			Environment: "test",
			AddSource:   false,
			Format:      "aboba",
			Filepath:    "",
		}

		_, err = logs.NewLogger(opt)
		require.Error(t, err)

		opt = cfg.LogsConfig{
			Environment: "test",
			Format:      "aboba",
		}

		_, err = logs.NewLogger(opt)
		require.Error(t, err)
	})

	t.Run("test logging config", func(t *testing.T) {
		hostname, _ := os.Hostname()

		opt := cfg.LogsConfig{
			Host:        hostname,
			Environment: "dev",
			Level:       int(slog.LevelInfo),
			AddSource:   false,
			Format:      "json",
			Filepath:    "",
		}

		logger, err := logs.NewLogger(opt)
		defer logs.CloseFiles()

		require.NoError(t, err)
		require.NotNil(t, logger)
	})

	t.Run("test file logging config", func(t *testing.T) {
		hostname, _ := os.Hostname()

		opt := cfg.LogsConfig{
			Host:        hostname,
			Environment: "dev",
			Level:       int(slog.LevelInfo),
			AddSource:   false,
			Format:      "json",
			Filepath:    filePath,
		}

		logger, err := logs.NewLogger(opt)
		defer logs.CloseFiles()

		require.NoError(t, err)
		require.NotNil(t, logger)

		logger.InfoContext(context.Background(), "info message 1")
		logger.InfoContext(context.Background(), "info message 2")
		logger.InfoContext(context.Background(), "info message 3")

		stat, err := os.Stat(filePath)
		require.NoError(t, err)
		require.NotZero(t, stat.Size())
		require.EqualValues(t, filepath.Base(filePath), stat.Name())
	})

	t.Run("test context logging", func(t *testing.T) {
		opt := cfg.LogsConfig{
			Environment: "dev",
			Level:       int(slog.LevelInfo),
			AddSource:   false,
			Format:      "json",
		}

		logger, err := logs.NewLogger(opt)
		require.NoError(t, err)
		require.NotNil(t, logger)

		ctx := context.WithValue(context.Background(), logs.MetadataKey_CorrelationUUID, uuid.New().String())

		logger.InfoContext(ctx, "info message 1")
		logger.InfoContext(ctx, "info message 2")
		logger.InfoContext(ctx, "info message 3")
	})
}
