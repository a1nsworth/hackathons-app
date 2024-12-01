package log

import (
	"log/slog"
	"os"
)

func NewLogger() (log *slog.Logger) {
	log = slog.New(slog.NewJSONHandler(os.Stdout, nil))
	return
}
