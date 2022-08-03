package configs

import (
	"encoding/json"
	"golang-api-structure/constants"
	"log"
	"os"
)

// Env the 3ºparam it's optional and defines the type of encoding, in this case is json
type Env struct {
	SrvAddr string `json:"SRV_ADDR"`
	SrvPort string `json:"SRV_PORT"`
}

func InitDefaultEnv() Env {
	return Env{
		SrvAddr: "0.0.0.0",
		SrvPort: "8081",
	}
}

func GetEnvConfigs() Env {
	env := InitDefaultEnv()

	file, err := os.ReadFile(constants.EnvFileName)

	if err != nil {
		log.Println("ERROR:", err)
		log.Println("Using default env:", env)
		return env
	}

	json.Unmarshal([]byte(file), &env)

	return env
}
