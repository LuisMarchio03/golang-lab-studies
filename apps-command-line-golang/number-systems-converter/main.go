package main

import (
	"log"
	"number-systems-converter/app"
	"os"
)

func main() {
	application := app.App()
	err := application.Run(os.Args)
	if err != nil {
		log.Fatalf("%v", err)
	}
}
