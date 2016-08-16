package core

import (
	"fmt"

	"github.com/spf13/viper"
	"github.com/asaskevich/govalidator"
)

const cfgName = "hb.yaml"

// Flags holds the CLI configuration values and flags
type Flags struct {
	Verbose      bool
	LogLevel     bool
	SlackWebHook string
}

func init() {
	viper.SetConfigFile(cfgName)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
}

func validateConfig() error {
	if !govalidator.IsURL(GetSlackWebHookUrl()) {
		return fmt.Errorf("Invalid WebHookUrl!")
	}
	return nil
}

func ReadConfig() error {
	err := viper.ReadInConfig()
	if err != nil {
	    return fmt.Errorf("Fatal error config file: %s \n", err)
	}
	return validateConfig()
}

func GetSlackWebHookUrl() string {
	return viper.GetString("slack.url")
}

func GetSlackBotName() string {
	return viper.GetString("slack.botname")
}

func GetSlackBotIcon() string {
	return viper.GetString("slack.boticon")
}

func GetSlackTarget() string {
	return viper.GetString("slack.target")
}
