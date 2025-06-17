package main

import (
	"errors"
	"log/slog"
	"os"
	"strings"

	"github.com/fabienjuif/console-slog"
)

func main() {
	logger := slog.New(
		console.NewHandler(os.Stderr, &console.HandlerOptions{
			Level:      slog.LevelDebug,
			AddSource:  true,
			TimeFormat: "15:04:05.000000",
			ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
				if v, ok := a.Value.Any().(*slog.Source); ok {
					file := v.File
					parts := strings.Split(v.File, "/")
					if len(parts) > 0 {
						file = parts[len(parts)-1]
					}
					return slog.Attr{
						Key: slog.SourceKey,
						Value: slog.AnyValue(&slog.Source{
							Function: v.Function,
							File:     file,
							Line:     v.Line,
						}),
					}
				}
				return a
			},
		}),
	)
	slog.SetDefault(logger)
	slog.Info("Hello world!", "foo", "bar")
	slog.Debug("Debug message")
	slog.Warn("Warning message")
	slog.Error("Error message", "err", errors.New("the error"))

	logger = logger.With("foo", "bar").
		WithGroup("the-group").
		With("bar", "baz")

	logger.Info("group info", "attr", "value")
}
