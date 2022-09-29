package commons

import (
	"os"

	"github.com/pkg/errors"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logger *zap.Logger
	atomic zap.AtomicLevel
)

func init() {
	atomic = zap.NewAtomicLevel()
	atomic.SetLevel(zap.InfoLevel)
	logger = zap.New(zapcore.NewTee(zapcore.NewCore(zapcore.NewJSONEncoder(zapcore.EncoderConfig{
		MessageKey:   "message",
		LevelKey:     "severity",
		EncodeLevel:  zapcore.LowercaseLevelEncoder,
		TimeKey:      "time",
		EncodeTime:   zapcore.RFC3339TimeEncoder,
		CallerKey:    "caller",
		EncodeCaller: zapcore.FullCallerEncoder,
	}), zapcore.Lock(os.Stdout), atomic)))
}

// GetLogger returns the shared *VpnbeastLogger
func GetLogger() *zap.Logger {
	return logger
}

// SetLogLevel function sets the log level of the *zap.Logger using zap.AtomicLevel dynamically
func SetLogLevel(level string) (*zap.Logger, error) {
	if logger != nil {
		parsedLevel, err := zapcore.ParseLevel(level)
		if err != nil {
			return nil, err
		}

		atomic.SetLevel(parsedLevel)
		return logger, nil
	}

	return nil, errors.New("logger not initialized yet")
}
