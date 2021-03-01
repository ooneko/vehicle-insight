package main

import (
	"log"

	"ooneko.github.com/vehicle-insight/cmd/app"
)

func main() {
	cmd := app.NewAPIServerCommand()

	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
