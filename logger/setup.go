package logger

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func Setup() {

	fileName := "log/web/" + time.Now().Format("2006.01.02") + ".log"

	level := GetLoggerLevel("debug")
	syncWriter := zapcore.AddSync(&lumberjack.Logger{
		Filename:  fileName,
		MaxSize:   1 << 30, //1G
		LocalTime: true,
		Compress:  true,
	})

	encoder := zap.NewProductionEncoderConfig()
	encoder.TimeKey = "timestamp"
	encoder.CallerKey = "module"
	encoder.EncodeTime = zapcore.ISO8601TimeEncoder
	enc := zapcore.NewJSONEncoder(encoder)

	core := zapcore.NewCore(enc, syncWriter, zap.NewAtomicLevelAt(level))
	zaplogger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(2))
	ErrorLogger = zaplogger.Sugar()

}
