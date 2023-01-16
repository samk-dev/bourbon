package main

import (
	"log"
	"os"

	"github.com/samk-dev/bourbon"
)

func initApp() *application {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// init bourbon
	bourbon := &bourbon.Bourbon{}
	err = bourbon.New(path)
	if err != nil {
		log.Fatal(err)
	}

	bourbon.AppName = "testapp"

	bourbon.InfoLog.Println("Debug is set to", bourbon.Debug)

	app := &application{
		App: bourbon,
	}

	return app
}
