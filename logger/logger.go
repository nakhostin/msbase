package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// error logger
var ErrorLogger *zap.SugaredLogger

// var ConsoleLog bool

var levelMap = map[string]zapcore.Level{
	"debug":  zapcore.DebugLevel,
	"info":   zapcore.InfoLevel,
	"warn":   zapcore.WarnLevel,
	"error":  zapcore.ErrorLevel,
	"dpanic": zapcore.DPanicLevel,
	"panic":  zapcore.PanicLevel,
	"fatal":  zapcore.FatalLevel,
}

func GetLoggerLevel(lvl string) zapcore.Level {
	if level, ok := levelMap[lvl]; ok {
		return level
	}
	return zapcore.InfoLevel
}

func Debug(args ...interface{}) {
	ErrorLogger.Debug(args...)
}

func Debugf(template string, args ...interface{}) {
	ErrorLogger.Debugf(template, args...)
}

func Info(args ...interface{}) {
	ErrorLogger.Info(args...)
}

func Infof(template string, args ...interface{}) {
	ErrorLogger.Infof(template, args...)
}

func Warn(args ...interface{}) {
	ErrorLogger.Warn(args...)
}

func Warnf(template string, args ...interface{}) {
	ErrorLogger.Warnf(template, args...)
}

func Error(args ...interface{}) {
	ErrorLogger.Error(args...)
}

func Errorf(template string, args ...interface{}) {
	ErrorLogger.Errorf(template, args...)
}

func DPanic(args ...interface{}) {
	ErrorLogger.DPanic(args...)
}

func DPanicf(template string, args ...interface{}) {
	ErrorLogger.DPanicf(template, args...)
}

func Panic(args ...interface{}) {
	ErrorLogger.Panic(args...)
}

func Panicf(template string, args ...interface{}) {
	ErrorLogger.Panicf(template, args...)
}

func Fatal(args ...interface{}) {
	ErrorLogger.Fatal(args...)
}

func Fatalf(template string, args ...interface{}) {
	ErrorLogger.Fatalf(template, args...)
}
