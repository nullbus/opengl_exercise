package main

import (
	"fmt"
	"log"

	exercise "github.com/nullbus/opengl_exercise"
	"golang.org/x/sys/windows"
)

func main() {
	namea := "ChoosePixelFormat"
	opengl32 := windows.NewLazySystemDLL("gdi32")
	p := opengl32.NewProc(namea)
	if err := p.Find(); err != nil {
		// The proc is not found.
		fmt.Println(err)
	}

	//
	system := exercise.System{
		Fullscreen: false,
	}
	defer system.Shutdown()

	if err := system.Initialize(); err != nil {
		log.Println("error", err)
	} else {
		system.Run()
	}
}
