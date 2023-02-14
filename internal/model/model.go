package model

import "github.com/spf13/viper"

type AppConfig struct {
	DB DB
}

type DB struct {
	Host     string
	Port     string
	UserName string
	DbName   string
	Password string
	SslMode  string
}

func NewAppConfig(fileName string) (*AppConfig, error) {
	config, err := LoadConfigFile(fileName)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func LoadConfigFile(fileName string) (*AppConfig, error) {
	//Define path and extension of config file
	viper.SetConfigFile(fileName)
	//Read config file
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var appConfig AppConfig
	if err := viper.Unmarshal(&appConfig); err != nil {
		return nil, err
	}

	return &appConfig, nil
}
