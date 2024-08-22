package global

import (
	"gorm.io/gorm"
	"gvb_server/config"
)

var (
	Config *config.Config
	DB     *gorm.DB
	//Log    *log.Logger
)
