package global

import (
	"github.com/spf13/viper"

	"go.uber.org/zap"

	"github.com/apache/dubbo-go-pixiu/pkg/admin/config"
	"gorm.io/gorm"
)

var (
	VP     *viper.Viper
	DB     *gorm.DB
	CONFIG config.Server
	LOG    *zap.Logger
)
