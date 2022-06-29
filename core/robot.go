package core

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	ROBOT_DB     *gorm.DB
	ROBOT_CONFIG *viper.Viper
	ROBOT_LOGGER *zap.Logger
)
