package conf

import "go.uber.org/zap"

var (
	logger *zap.SugaredLogger
)

func init() {
	l, _ := zap.NewProduction()
	logger = l.Sugar()
}

//GetLogger returns zap sugar logger
func GetLogger() *zap.SugaredLogger {
	return logger
}
