package main

import (
	"github.com/hello-go-gl/pkg/render"
)

func main() {
	options := &render.Options{
		Width:  512,
		Height: 512,
		Title:  "go!",
	}

	render.Loop(options)
}
