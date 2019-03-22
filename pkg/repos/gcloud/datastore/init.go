package datastore

import (
	"github.com/klahssen/webapp/pkg/conf"
	"go.uber.org/zap"
)

var (
	logger *zap.SugaredLogger
)

func init() {
	logger = conf.GetLogger()
}
