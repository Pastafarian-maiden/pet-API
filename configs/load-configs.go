package configs

import (
	"github.com/spf13/viper"
)

type Config struct {
	Username string `mapstructure:USERNAME_BASIC`
	Password string `mapstructure:PASSWORD_BASIC`
	Address  string `mapstructure:ADDR_API`
}

func LoadConfig() (config Config, err error) {
	viper.AddConfigPath("./")
	viper.SetConfigName("config")
	viper.SetConfigType("json")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return

	// content, err := ioutil.ReadFile("./config.json")
	// if err != nil {
	// 	log.Fatal("Error when opening file: ", err)
	// }

	// var env map[string]interface{}
	// err = json.Unmarshal(content, &env)
	// if err != nil {
	// 	log.Fatal("Error during Unmarshal(): ", err)
	// }

	// return &Configs{
	// 	env: env,
	// }
}
