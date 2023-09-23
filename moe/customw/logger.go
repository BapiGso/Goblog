package customw

import (
	"log/slog"
	"os"
)

var newLog = func() *slog.Logger {
	logger := slog.New(
		slog.NewTextHandler(
			os.Stdout, &slog.HandlerOptions{
				AddSource: true,
				Level:     slog.LevelInfo,
				ReplaceAttr: func(_ []string, a slog.Attr) slog.Attr {
					return a
				},
			},
		),
	)
	slog.SetDefault(logger)
	return slog.Default()
}()
