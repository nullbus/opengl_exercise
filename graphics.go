package opengl_exercise

import (
	"time"

	"github.com/TheTitanrain/w32"
)

type Graphics struct {
	opengl *OpenGL
}

func NewGraphics(opengl *OpenGL, hwnd w32.HWND) (*Graphics, error) {
	return &Graphics{
		opengl: opengl,
	}, nil
}

func (g *Graphics) Frame(delta time.Duration) error {
	return g.render()
}

func (g *Graphics) Shutdown() {
	g.opengl = nil
}

func (g *Graphics) render() error {
	g.opengl.BeginScene(0.5, 0.5, 0.5, 1)
	g.opengl.EndScene()

	return nil
}
