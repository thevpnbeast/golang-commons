package commons

import (
	"log"
	"os"
	"strings"

	errors "github.com/pkg/errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type VpnbeastLogger struct {
	Logger *zap.Logger
}

var logger *VpnbeastLogger

func NewLogger() (*VpnbeastLogger, error) {
	return &VpnbeastLogger{Logger: zap.New(zapcore.NewTee(zapcore.NewCore(zapcore.NewJSONEncoder(zapcore.EncoderConfig{
		MessageKey:   "message",
		LevelKey:     "severity",
		EncodeLevel:  zapcore.LowercaseLevelEncoder,
		TimeKey:      "time",
		EncodeTime:   zapcore.RFC3339TimeEncoder,
		CallerKey:    "caller",
		EncodeCaller: zapcore.FullCallerEncoder,
	}), zapcore.Lock(os.Stdout), zap.InfoLevel)))}, nil
}

func init() {
	var err error
	logger, err = NewLogger()
	if err != nil {
		panic(err)
	}
}

// GetLogger returns the shared *VpnbeastLogger
func GetLogger() *VpnbeastLogger {
	return logger
}

func (l *VpnbeastLogger) Log(logLevel, message string) {
	level, err := zapcore.ParseLevel(strings.ToLower(logLevel))
	if err != nil {
		log.Println(errors.Wrapf(err, "unable to parse log level %s", strings.ToLower(logLevel)))
		return
	}

	defer func() {
		_ = l.Logger.Sync()
	}()

	l.Logger.Log(level, message)
}
