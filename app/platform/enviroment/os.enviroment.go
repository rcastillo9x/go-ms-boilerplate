package enviroment

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func GetEnvVariable(key string) string {
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

	value, ok := viper.Get(key).(string)

	if !ok {
		log.Fatalf("Invalid type assertion")
	}

	return value
}

func GetHTTPPort(key string) string {
	return fmt.Sprintf(":%s", GetEnvVariable(key))
}
