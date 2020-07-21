package main

import (
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"golang.org/x/sys/unix"
)

func main() {

	// Get current' pid
	id := unix.Getpid()

	//
	atom := zap.NewAtomicLevel()

	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.TimeKey = "timestamp"
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder

	logger := zap.New(zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderCfg),
		zapcore.Lock(os.Stdout),
		atom,
	))

	defer logger.Sync()

	atom.SetLevel(zap.DebugLevel)

	for {
		logger.Sugar().Infow("failed to fetch URL",
			"pid", id,
			"url", "http://example.com",
			"attempt", 3,
			"backoff", time.Second,
		)

		// Sleeping 5s for next print
		time.Sleep(5 * time.Second)
	}
}
