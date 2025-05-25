package main

import (
	"github.com/henning-konigs/fem-go-crud/internal/app"
)

func main() {
	app, err := app.NewApplication()

	if err != nil {
		panic(err)
	}

	app.Logger.Println("We are running our app")
}
