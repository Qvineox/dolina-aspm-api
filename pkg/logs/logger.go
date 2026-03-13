package logs

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"

	"github.com/go-playground/validator/v10"
	"gitlab.domsnail.ru/dolina/dolina-aspm-api/pkg/cfg"
)

// wiki: https://gitlab.domsnail.ru/dolina/dolina-aspm-api/-/wikis/%D0%9B%D0%BE%D0%B3%D0%B8%D1%80%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D0%B5

// guide: https://www.dash0.com/guides/logging-in-go-with-slog
// guide: https://habr.com/ru/companies/slurm/articles/798207/?ysclid=mmnh7y0xzr879852594

var file *os.File

func NewLogger(opts cfg.LogsConfig) (*slog.Logger, error) {
	validate := validator.New()

	err := validate.Struct(opts)
	if err != nil {
		return nil, fmt.Errorf("validation error: %w", err)
	}

	var output io.Writer
	if opts.Filepath != "" {
		file, err = os.Create(opts.Filepath)
		if err != nil {
			return nil, fmt.Errorf("failed to create file '%s': %w", opts.Filepath, err)
		}

		output = file
	} else {
		output = os.Stdout
	}

	var handlerOpts = slog.HandlerOptions{
		AddSource: opts.AddSource,
		Level:     slog.Level(opts.Level),
	}

	// todo: add context handler for correlation_uuid, service

	var handler slog.Handler
	switch opts.Format {
	case "json":
		handler = slog.
			NewJSONHandler(output, &handlerOpts).
			WithAttrs([]slog.Attr{
				slog.String("host", opts.Host),
				slog.String("env", opts.Environment),
			})
	case "text":
		handler = slog.NewTextHandler(output, &handlerOpts)
	default:
		// sanity check
		return nil, fmt.Errorf("unknown format: %s", opts.Format)
	}

	return slog.New(&ContextHandler{handler}), nil
}

const (
	MetadataKey_CorrelationUUID = "correlation_uuid"
)

type ContextHandler struct {
	slog.Handler
}

func NewContextHandler(h slog.Handler) *ContextHandler {
	return &ContextHandler{Handler: h}
}

func (ch *ContextHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return ch.Handler.Enabled(ctx, level)
}

func (ch *ContextHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return NewContextHandler(ch.Handler.WithAttrs(attrs))
}

func (ch *ContextHandler) WithGroup(name string) slog.Handler {
	return NewContextHandler(ch.Handler.WithGroup(name))
}

func (ch *ContextHandler) Handle(ctx context.Context, r slog.Record) error {
	if id := ctx.Value(MetadataKey_CorrelationUUID); id != nil {
		if idStr, ok := id.(string); ok && idStr != "" {
			r.AddAttrs(slog.String("correlation_id", idStr))
		}
	}

	return ch.Handler.Handle(ctx, r)
}

func CloseFiles() {
	if file != nil {
		_ = file.Close()
	}
}
