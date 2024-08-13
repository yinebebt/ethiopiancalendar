package config

import (
	"context"
	"fmt"
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func InitConfig(path string) {
	viper.SetConfigFile(path)
	err := viper.ReadInConfig()
	if err != nil {
		log.Panic(context.Background(), fmt.Sprintf("Failed to read config: %v", err))
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Println(context.Background(), "Config file changed:", zap.String("file", e.Name))
	})

	viper.AutomaticEnv()

	bindEnv("server.host", "SERVER_HOST")
	bindEnv("server.port", "SERVER_PORT")
	bindEnv("server.schemes", "SERVER_SCHEMES")
}

func bindEnv(key, envVar string) {
	if err := viper.BindEnv(key, envVar); err != nil {
		log.Panic(context.Background(), fmt.Sprintf("failed to bind env variable %s to key %s: %v", envVar, key, err))
	}
}
