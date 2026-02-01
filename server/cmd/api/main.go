package main

import (
	"log"

	"github.com/saianand32/Quipify/internal/config"
)

func main() {
	c := config.New()
	app := &application{c}
	log.Fatal(app.run())
}
