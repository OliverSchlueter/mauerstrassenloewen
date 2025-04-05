package cmd

import "log/slog"

func main() {
	slog.SetLogLoggerLevel(slog.LevelDebug)

	slog.Info("Hello, world!")
}
