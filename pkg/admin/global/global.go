package global

import (
	"github.com/apache/dubbo-go-pixiu/pkg/admin/config"

	"github.com/spf13/viper"

	"go.uber.org/zap"

	"gorm.io/gorm"
)

var (
	VP     *viper.Viper
	DB     *gorm.DB
	CONFIG config.Server
	LOG    *zap.Logger
)
