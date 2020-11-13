package setting

import (
	"gopkg.in/ini.v1"
)

var Conf = new(AppConfig)

// AppConfig 应用程序配置
type AppConfig struct {
	Release      bool `ini:"release"`
	Port         int  `ini:"port"`
	*SQLite3Config `ini:"sqlite3"` // unused !
}

// MySQLConfig 数据库配置
type SQLite3Config struct {
	DB       string `ini:"db"`
}

func Init(file string) error {
	return ini.MapTo(Conf, file)
}
