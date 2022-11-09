package global

import (
	"aixinge/config"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	DB     *gorm.DB
	CONFIG config.Server
	VP     *viper.Viper
	LOG    *zap.Logger
)
