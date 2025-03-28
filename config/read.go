package config

import (
	"OpenTan/utils"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/kelseyhightower/envconfig"
	"github.com/spf13/viper"
)

const defaultFilePath = "./config.yml" // Correct when called from main.go
const envPrefix = "OpenTan"

var c Config

func init() {
	filePath := defaultFilePath
	viper.SetConfigFile(filePath)
	if utils.FileExist(filePath) {
		utils.PanicOnErr(viper.ReadInConfig())
		utils.PanicOnErr(viper.Unmarshal(&c))
	} else {
		fmt.Println("Config file not exist in ", filePath, ". Using environment variables.")
		utils.PanicOnErr(envconfig.Process(envPrefix, &c)) // Inject environment variables like `ENVPREFIX_XXX`into the c variable
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Printf("Config file changed: %s; ", e.Name)
		if err := viper.ReadInConfig(); err != nil {
			fmt.Println("Error reading config file:", err)
		}
		utils.PanicOnErr(viper.Unmarshal(&c))
		fmt.Println("Config file reloaded.")
	})
}

// Set manually sets the config and is not recommended to use in production
func Set(config Config) {
	c = config
}

func Get() Config {
	return c
}

func IsRelease() bool {
	return c.Mode == ModeRelease
}

func IsDebug() bool {
	return c.Mode == ModeDebug
}
