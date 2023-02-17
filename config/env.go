package config

import (
	"os"
)

//Mustget gets the values of an env variables identified by a Key
func MustGet(key string, def string) string {
	value := os.Getenv(key)
	if def != "" && value == "" {
		return def
	} else if def == "" && value == "" {
		panic(key + ":not found in env")
	}
	return value
}
