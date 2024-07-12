package env

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
	"github.com/tinhminhtue/go-reused-lib/core/defined"
)

func LoadEnv() {
	if os.Getenv(defined.USE_LOCAL_SECERT_ENV) == "true" {
		viper.SetConfigName("local-secret") // name of config file (without extension)
	} else {
		viper.SetConfigName("default") // name of config file (without extension)
	}

	viper.SetConfigType("yaml") // REQUIRED if the config file does not have the extension in the name
	// viper.AddConfigPath("$HOME/.appname")  // call multiple times to add many search paths
	viper.AddConfigPath("./cfg") // optionally look for config in the working directory
	err := viper.ReadInConfig()  // Find and read the config file
	if err != nil {              // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_")) // Handle nested keys
	viper.AutomaticEnv()
	if os.Getenv("DEBUG") != "" {
		viper.Debug()
	}

}
