package main

import "log"

func main() {
	c := config{
		addr: ":8080",
	}
	app := &application{
		config: c,
	}

	log.Fatal(app.run())
}
