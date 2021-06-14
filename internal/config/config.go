package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type CuckooSetting struct {
	File     string
	RpcPort  uint
	Capacity uint
}

func Read() *CuckooSetting {
	var cuckooConf CuckooSetting
	v := viper.New()
	v.SetConfigName("config")
	v.AddConfigPath(".")
	v.SetConfigType("json")
	err := v.ReadInConfig()
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}
	if err := v.Unmarshal(&cuckooConf); err != nil {
		fmt.Println(err)
	}
	return &cuckooConf
}
