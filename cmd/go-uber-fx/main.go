package main

import (
	"github.com/wjojf/go-uber-fx/internal/pkg/app"
)

func main() {
	a, err := app.New()
	if err != nil {
		panic(err)
	}

	a.Run()
}
