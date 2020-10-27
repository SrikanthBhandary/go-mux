package main

import (
	"github.com/SrikanthBhandary/go-mux/app"
)

func main() {
	a := app.App{}
	// a.Initialize(
	// 	os.Getenv("APP_DB_USERNAME"),
	// 	os.Getenv("APP_DB_PASSWORD"),
	// 	os.Getenv("APP_DB_NAME"),
	// )

	a.Initialize(
		"postgres",
		"admin",
		"postgres",
	)
	a.Run(":8010")
}
