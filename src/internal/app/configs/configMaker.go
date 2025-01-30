package configs

import (
	"log"
	"os"
	"sync"

	"github.com/pelletier/go-toml/v2"
)

type ApplicationConfg struct {
	API_SERVER_PORT int
}

var applicationConfigOnce sync.Once
var ApplicationConfig_Single *ApplicationConfg

func GetApplicationConfigInstance() *ApplicationConfg {
	applicationConfigOnce.Do(func() {
		ApplicationConfig_Single = &ApplicationConfg{}

		config_data, err := os.ReadFile("./_config.toml")
		if err != nil {
			log.Fatalf("Cant parse config data : %s", err)
		}
		err = toml.Unmarshal(config_data, &ApplicationConfig_Single)
		if err != nil {
			log.Fatalf("Error with unmarshal config data : %s", err)
		}
	})

	return ApplicationConfig_Single
}
