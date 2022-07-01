package core

import (
	"github.com/allegro/bigcache/v3"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	ROBOT_DB     *gorm.DB
	ROBOT_CONFIG *viper.Viper
	ROBOT_LOGGER *zap.Logger
	ROBOT_CACHE  *bigcache.BigCache
)
