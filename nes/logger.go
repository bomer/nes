package nes

import (
	"fmt"
	"log/slog"
)

// Debugf is capitalized so any file in or out of the nes package can use it
func Debugf(format string, args ...any) {
	if slog.Default().Enabled(nil, slog.LevelDebug) {
		slog.Debug(fmt.Sprintf(format, args...))
	}
}
