package logger_test

import (
	"testing"
	"time"

	"github.com/mrchuanxu/vito_infra/logger"
)

func Test_Log(t *testing.T) {
	defer logger.TransLogger.Sync() // flushes buffer, if any
	sugar := logger.TransLogger.Sugar()
	url := "trans"
	sugar.Infow("failed to fetch URL",
		// Structured context as loosely typed key-value pairs.
		"url", url,
		"attempt", 3,
		"backoff", time.Second,
	)
	sugar.Infof("Failed to fetch URL: %s", url)
}
