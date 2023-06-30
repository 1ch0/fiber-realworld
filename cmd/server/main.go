package main

import (
	"github.com/1ch0/fiber-realworld/cmd/server/app"
	"github.com/1ch0/fiber-realworld/pkg/server/utils/log"
)

func main() {
	cmd := app.NewAPIServerCommand()
	if err := cmd.Execute(); err != nil {
		log.Logger.Fatalln(err)
	}
}
