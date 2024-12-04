package main

import "task-9/cmd/internal/app"

func main() {
	application := app.App{}

	if err := application.Run(); err != nil {
		panic(err)
	}
}
