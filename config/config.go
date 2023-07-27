package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/walbety/csv-file-generator/canonical"
)

type Configs struct {
	Filename   string            `json:"filename"`
	TotalLines int               `json:"totalLines"`
	Fields     []canonical.Field `json:"fields"`
}

const (
	CONFIG_FILE_TYPE = "json"
	CONFIG_FILE_NAME = "config.json"
)

var Cfg Configs

func InitConfigs() error {
	viper.SetConfigFile(CONFIG_FILE_NAME)
	viper.SetConfigType(CONFIG_FILE_TYPE)

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Can't find the file json : ", err)
		return err
	}

	err = viper.Unmarshal(&Cfg)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
		return err
	}

	log.Infof("Configs initialized with values: %v", Cfg)
	return nil
}
