package main

import "github.com/samk-dev/bourbon"

type application struct {
	App *bourbon.Bourbon
}

func main() {
	b := initApp()
	b.App.ListenAndServe()
}
