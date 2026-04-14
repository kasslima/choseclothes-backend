package main

import "choseclothes/internal/app"

func main() {
	r := app.BuildRouter()
	r.Run()
}
