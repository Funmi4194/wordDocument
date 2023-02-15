package config

import (
	"fmt"
	"reflect"

	"github.com/Funmi4194/myMod/enviroment"
)

func VerifyEnvironment(env enviroment.Env) error {
	//get the type of the argument
	t := reflect.TypeOf(env)
	if t == nil {
		return fmt.Errorf("env is nil")
	}
	//only allow struct type
	if t.Kind() != reflect.Struct {
		return fmt.Errorf("env is not a struct")
	}
	//verify each struct field tag
	for i := 0; i < t.NumField(); i++ {
		structField := t.Field(i)
		//get the field tag value
		tag := structField.Tag.Get(envTagName)
		if tag == "" {
			continue
		}
		//check if environment variable is set
		if MustGet(tag, " ") == " " {
			return fmt.Errorf("enviroment variable %s is not set", tag)
		}
	}
	return nil
}

func AppendEnvironment(env *enviroment.Env) {
	//get the type of argument
	t := reflect.TypeOf(*env)
	if t == nil {
		return
	}
	// only allow struct type
	if t.Kind() != reflect.Struct {
		return
	}
	// append each struct field tag
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		// get the field tag value
		tag := field.Tag.Get(envTagName)
		if tag == "" {
			continue
		}
		//append environment variable to const .ENV
		reflect.ValueOf(env).Elem().Field(i).SetString(MustGet(tag, ""))
	}
}
