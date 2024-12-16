package main

import "task-10/cmd/internal/app"

func main() {
	application := app.App{}

	if err := application.Run(); err != nil {
		panic(err)
	}
}
