package logger

import (
	"GoPolyglot/libs/common"
	"GoPolyglot/libs/configs"
	"github.com/outreach-golang/logger"
	"go.uber.org/zap"
)

var (
	GLogger *zap.Logger
)

func InitLogger() error {
	var (
		err  error
		conf = configs.GConfig.Log
	)
	GLogger, err = logger.NewLogger(
		logger.ServerName(configs.GConfig.Server.AppName),
		logger.SaveLogAddr(logger.SaveLogForm(conf.Storage)),

		logger.AccessKeyID(configs.GConfig.Log.AccessKeyID),
		logger.AccessKeySecret(configs.GConfig.Log.AccessKeySecret),
		logger.LogStore(configs.GConfig.Log.LogStore),
		logger.Endpoint(configs.GConfig.Log.Endpoint),
		logger.Project(configs.GConfig.Log.Project),
		logger.Source(common.LocalIP()+":"+configs.GConfig.Server.Port),
		logger.LooKAddr(configs.GConfig.Log.LookAddr),

		//logger.DingHost(configs.GConfig.Dingtalk.Host),
		//logger.DingWebhook(configs.GConfig.Dingtalk.Webhook),
	)

	if err != nil {
		return err
	}

	return nil
}
