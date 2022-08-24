package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

func init() {
	var err error

	config := zap.NewProductionConfig()
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.StacktraceKey = ""            
	config.EncoderConfig = encoderConfig
 
	log, err = config.Build(zap.AddCallerSkip(1))
 
	//log, err = zap.NewProduction()
	if err != nil {
		panic(err)
	}
	
	defer log.Sync()
}

func Info(message string, fields ...interface{}) {
	log.Sugar().Infow(message, fields)
	// log.Info(message, fields...)
}