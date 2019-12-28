package main

import (
	"log"

	exercise "github.com/nullbus/opengl_exercise"
	"github.com/nullbus/opengl_exercise/gl"
)

func c() {
	if err := gl.GetError(); err != 0 {
		log.Fatalln("fatal", err)
	}
}

func main() {
	system := exercise.System{
		Fullscreen:  false,
		ScreenNear:  0.1,
		ScreenDepth: 1000,
	}
	defer system.Shutdown()

	if err := system.Initialize(); err != nil {
		log.Println("error", err)
	} else {
		system.Run()
	}
}
