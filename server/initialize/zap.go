package initialize

import "go.uber.org/zap"

func InitZap() *zap.SugaredLogger {
	logger, _ := zap.NewProduction()
	return logger.Sugar()
}
