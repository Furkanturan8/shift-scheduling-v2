package logger

import (
	"github.com/google/wire"
	"shift-scheduling-V2/config"
	"shift-scheduling-V2/pkg/logger"
)

var Set = wire.NewSet(
	NewLoggerAplication,
)

// NewHandler Constructor
func NewLoggerAplication(cfg *config.Configuration) logger.Logger {
	return logger.NewApiLogger(cfg)
}
