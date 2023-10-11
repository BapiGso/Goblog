package tools

import (
	"log/slog"
	"os"
)

func LogInit() *slog.Logger {
	logger := slog.New(
		slog.NewTextHandler(
			os.Stdout, &slog.HandlerOptions{
				AddSource: false,
				Level:     slog.LevelInfo,
				ReplaceAttr: func(_ []string, a slog.Attr) slog.Attr {
					return a
				},
			},
		),
	)
	slog.SetDefault(logger)
	return slog.Default()
}
