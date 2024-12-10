package log

import (
	"context"
	"github.com/jackc/pgx"
	"io"
	"log"
	"os"
)

type Logger interface {
	Printf(ctx context.Context, format string, args ...interface{})
	Infof(ctx context.Context, format string, args ...interface{})
	Debugf(ctx context.Context, format string, args ...interface{})
	Errorf(ctx context.Context, format string, args ...interface{})
	Panicf(ctx context.Context, format string, args ...interface{})
	Log(level pgx.LogLevel, msg string, data map[string]interface{})
}

type LoggerConfig struct {
	Writer io.Writer
	Prifix string
	Flags  int
}

type logger struct {
	logger *log.Logger
}

func New(cfg *LoggerConfig) Logger {
	return &logger{
		logger: log.New(
			cfg.Writer,
			cfg.Prifix,
			cfg.Flags,
		),
	}
}

func GetLogger() Logger {
	return loggerInstance
}

var loggerInstance *logger

func init() {
	loggerInstance = &logger{
		logger: log.New(
			os.Stdout,
			"",
			log.Ldate|log.Ltime|log.Lshortfile,
		),
	}
}

func With(cfg *LoggerConfig) Logger {
	return &logger{
		logger: log.New(
			merge(cfg.Writer, loggerInstance.logger.Writer()),
			merge(cfg.Prifix, loggerInstance.logger.Prefix()),
			merge(cfg.Flags, loggerInstance.logger.Flags()),
		),
	}
}

func merge[T any](t1, t2 T) T {
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}
	return t2
}

func (l logger) Printf(_ context.Context, format string, args ...interface{}) {
	l.logger.Printf(format, args...)
}

func (l logger) Infof(_ context.Context, format string, args ...interface{}) {
	l.logger.Printf(format, args...)
}

func (l logger) Debugf(_ context.Context, format string, args ...interface{}) {
	l.logger.Printf(format, args...)
}

func (l logger) Errorf(_ context.Context, format string, args ...interface{}) {
	l.logger.Printf(format, args...)
}

func (l logger) Panicf(_ context.Context, format string, args ...interface{}) {
	l.logger.Panicf(format, args...)
}

func (l logger) Log(_ pgx.LogLevel, _ string, _ map[string]interface{}) {}
