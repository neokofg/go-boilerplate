package logs

import "go.uber.org/zap"

func InitZap() *zap.SugaredLogger {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()

	return sugar
}
