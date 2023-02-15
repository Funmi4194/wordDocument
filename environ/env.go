package environ

import (
	"log"

	"github.com/Funmi4194/myMod/config"
	"github.com/Funmi4194/myMod/enviroment"
	"github.com/joho/godotenv"
)

func Environ() {
	env := config.MustGet("ENV_PATH", ".env")
	log.Printf("Loading %s file\n", env)
	if err := godotenv.Load(env); err != nil {
		if err := godotenv.Load(); err != nil {
			log.Printf("Error verifying environment varaiables: %s\n", err)
		}
	}

	if err := config.VerifyEnvironment(enviroment.Env{}); err != nil {
		log.Fatalf("Error verifying environment variables: %s\n", err)
	}
	config.AppendEnvironment(config.Env)

}
