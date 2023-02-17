package config

import environ "github.com/Funmi4194/myMod/environment"

const (
	envTagName = "env"
)

//Global varables
var (
	//instance of Env struct
	Env = new(environ.Env)
)
