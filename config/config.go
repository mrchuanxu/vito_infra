package config

import (
	"strconv"
	"sync"

	"github.com/fsnotify/fsnotify"
	"github.com/prometheus/common/log"

	"github.com/spf13/viper"
	"github.com/VitoChueng/vito_infra/logger"
)

var config *viper.Viper
var once sync.Once
var watchOnce sync.Once

type IOnConfigChange interface {
	OnConfigChange()
}

var mutexConfigWatcher sync.Mutex
var mapConfigWatcherTriget = map[IOnConfigChange]struct{}{}

func Init(env string, cfgPath string) {
	v := viper.New()
	v.SetConfigType("yaml")
	v.SetConfigName("conf")
	if cfgPath == "" {
		cfgPath = "./"
	}
	v.AddConfigPath(cfgPath)
	if err := v.ReadInConfig(); err != nil {
		logger.TransLogger.Sugar().Errorf("ReadInConfig has err [%v]", err)
		panic(err)
	}
	config = v
	watcherConfig()
}

func GetConfig() *viper.Viper {
	once.Do(func() {
		if config == nil {
			Init("conf", "")
		}
	})

	return config
}

func SetConfig(c *viper.Viper) {
	config = c
}

func GetInt(s string, defaultValue int) int {
	v := GetConfig().GetString(s)

	if v != "" {
		if r, err := strconv.Atoi(v); err == nil {
			return r
		}

	}
	return defaultValue

}

func GetInt64(s string, defaultValue int64) int64 {
	v := GetConfig().GetString(s)

	if v != "" {
		if r, err := strconv.ParseInt(v, 10, 64); err == nil {
			return r
		}

	}
	return defaultValue

}

func GetString(s, defaultValue string) string {
	v := GetConfig().GetString(s)

	if v == "" {
		return defaultValue

	}
	return v

}

func watcherConfig() {
	watchOnce.Do(func() {
		if config != nil {
			config.WatchConfig()
			config.OnConfigChange(onConfigChange)
		}
	})
}

func onConfigChange(event fsnotify.Event) {
	log.Infof("on config change, name:%v, event:%v", event.Name, event.Op)
	mutexConfigWatcher.Lock()
	defer mutexConfigWatcher.Unlock()

	for k, _ := range mapConfigWatcherTriget {
		k.OnConfigChange()
	}
}

func AddConfigWatcherTriger(iocc IOnConfigChange) {
	mutexConfigWatcher.Lock()
	defer mutexConfigWatcher.Unlock()

	if iocc == nil {
		log.Warn("add config watcher triger fail, nil function pointer")
		return
	}

	for k, _ := range mapConfigWatcherTriget {
		if k == iocc {
			log.Warnf("repeated add  config watcher triger %v", iocc)
			return
		}
	}

	mapConfigWatcherTriget[iocc] = struct{}{}
}

func DelConfigWatcherTriger(iocc IOnConfigChange) {
	mutexConfigWatcher.Lock()
	defer mutexConfigWatcher.Unlock()

	delete(mapConfigWatcherTriget, iocc)
}
