package logger

import (
	"sync"

	"go.uber.org/zap"
)

var (
	onceLogger sync.Once
	// TransLogger markLogger
	TransLogger *zap.Logger
)

func init() {
	initLogger := func() {
		TransLogger, _ = zap.NewDevelopment()
	}
	onceLogger.Do(initLogger)
}
