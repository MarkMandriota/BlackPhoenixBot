package main

import (
	"fmt"
	"math"
	"os"
	"reflect"
)

var config = struct {
	Token  string `env:"TOKEN"`
	Prefix string `env:"PREFIX"`
}{}

func init() {
	confV := reflect.ValueOf(&config).Elem()
	confT := confV.Type()

	for i := 0; i < confV.NumField(); i++ {
		fieldV := confV.Field(i)
		fieldT := confT.Field(i)

		if fieldT.Type.Kind() == reflect.String && fieldV.CanSet() {
			fieldV.SetString(os.Getenv(fieldT.Tag.Get("env")))
		}
	}
}

func main() {

}
