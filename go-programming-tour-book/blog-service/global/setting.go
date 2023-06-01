package global

import (
	"github.com/Temptation1/go_blog/go-programming-tour-book/blog-service/pkg/logger"
	"github.com/Temptation1/go_blog/go-programming-tour-book/blog-service/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	JWTSetting      *setting.JWTSettingS

	Logger *logger.Logger
)
