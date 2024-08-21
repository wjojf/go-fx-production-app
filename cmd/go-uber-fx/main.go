package main

import (
	"github.com/wjojf/go-uber-fx/internal/pkg/app"
)

func main() {
	app, err := app.New()
	if err != nil {
		panic(err)
	}

	app.Run()
}
