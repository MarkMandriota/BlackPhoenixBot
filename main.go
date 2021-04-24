package main

import (
	"os"
	"reflect"

	dg "github.com/bwmarrin/discordgo"
)

var Config = struct {
	Token  string `env:"TOKEN"`
	Prefix string `env:"PREFIX"`
}{}

func init() {
	confV := reflect.ValueOf(&Config).Elem()
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
	s, err := dg.New(Config.Token)
	if err != nil {
		panic(err)
	}

	if err = s.Open(); err != nil {
		panic(err)
	}
	defer s.Close()

	<-make(chan struct{})
}
