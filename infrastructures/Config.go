package infrastructures

import (
	"io/ioutil"

	"github.com/fsnotify/fsnotify"

	log "github.com/sirupsen/logrus"
	config "github.com/spf13/viper"
)

// SetConfig init configuration
func SetConfig() {
	config.SetConfigName("App")
	config.SetConfigType("yaml")
	config.AddConfigPath("./configurations")

	err := config.ReadInConfig()
	if err != nil {
		log.Fatal("config error: ", err)
	}

	config.WatchConfig()
	config.OnConfigChange(func(e fsnotify.Event) {
		log.Warn("Config file changed:", e.Name)
	})

	log.AddHook(NewLogHook().
		SetFormatType(config.GetString("log.format_output")).
		SetLogLevel(config.GetInt("log.level")).
		SetRotateLog(config.GetString("log.rotate")).
		SetLogType(config.GetString("log.type")),
	)

	log.SetOutput(ioutil.Discard)
}
