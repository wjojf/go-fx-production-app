package main

import "github.com/wjojf/go-uber-fx/internal/pkg/app"

func main() {
	a, err := app.NewConsumer()
	if err != nil {
		panic(err)
	}

	a.Run()
}
