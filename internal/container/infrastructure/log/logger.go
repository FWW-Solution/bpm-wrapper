package logger

import (
	"bpm-wrapper/internal/config"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Initialize(cfg *config.Config) *zap.SugaredLogger {
	zapOptions := []zap.Option{
		zap.AddStacktrace(zap.FatalLevel),
		zap.AddCallerSkip(1),
	}

	if cfg.IsVerbose {
		zapOptions = append(zapOptions, zap.IncreaseLevel(
			zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
				return lvl >= zap.DebugLevel
			}),
		))
	}

	logger, _ := zap.NewProduction(
		zapOptions...,
	)
	defer logger.Sync()
	sugar := logger.Sugar()
	return sugar
}
