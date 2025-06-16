package global

import (
	"github.com/spf13/viper"

	"go.uber.org/zap"

	"github.com/dubbo-go-pixiu/pixiu-api/pkg/api/config"
	"gorm.io/gorm"
)

var (
	VP     *viper.Viper
	DB     *gorm.DB
	CONFIG config.Server
	LOG    *zap.Logger
)
