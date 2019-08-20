package opengl_exercise

import (
	"fmt"
	"time"

	"github.com/nullbus/opengl_exercise/gl"

	"github.com/TheTitanrain/w32"
)

var vbuffer uint32
var vbufferData = [...]float32{
	-1, -1, 0,
	1, -1, 0,
	0, 1, 0,
}

type Graphics struct {
	opengl *OpenGL
	camera *Camera
	model  *Model
	shader *ColorShader
}

func NewGraphics(opengl *OpenGL, hwnd w32.HWND) (*Graphics, error) {
	graphics := &Graphics{
		opengl: opengl,
	}

	return graphics, graphics.initialize()
}

func (g *Graphics) initialize() (err error) {
	// create camera
	g.camera = NewCamera()

	// initial position of the camera
	g.camera.Position = Vector{Z: -10}

	// create model
	if g.model, err = NewModel(); err != nil {
		return err
	}

	// create color shader
	if g.shader, err = NewColorShader(); err != nil {
		return err
	}

	return nil
}

func check() {
	if err := gl.GetError(); err != gl.NO_ERROR {
		panic(fmt.Sprintf("%08x", err))
	}
}

func (g *Graphics) Frame(delta time.Duration) error {
	return g.render()
}

func (g *Graphics) Shutdown() {
	// release the color shader object
	if g.shader != nil {
		g.shader.Shutdown()
		g.shader = nil
	}

	// release the model object
	if g.model != nil {
		g.model.Shutdown()
		g.model = nil
	}

	// release the camera object
	if g.camera != nil {
		g.camera = nil
	}

	// release the pointer to the opengl class object
	g.opengl = nil
}

func (g *Graphics) render() error {
	// clear buffers to begin the scene
	g.opengl.BeginScene(0, 0, 0, 1)

	// generate view matrix baseed on the camera's position
	g.camera.Render()

	// get matrices from opengl and camera objects
	worldMatrix := g.opengl.WorldMatrix()
	viewMatrix := g.camera.ViewMatrix()
	projectionMatrix := g.opengl.ProjectionMatrix()

	// set color shader as the current shader program and set the matrices that will use for rendering
	g.shader.SetShader()
	if err := g.shader.SetShaderParams(worldMatrix, viewMatrix, projectionMatrix); err != nil {
		return err
	}

	// fmt.Println("-----------------------------------------------")
	// worldMatrix.Print(os.Stdout)
	// viewMatrix.Print(os.Stdout)
	// projectionMatrix.Print(os.Stdout)
	// fmt.Printf("%+v\n", Vector{-1, -1, 0}.MultiplyMatrix(&worldMatrix).MultiplyMatrix(&viewMatrix).MultiplyMatrix(&projectionMatrix))
	// fmt.Printf("%+v\n", Vector{0, 1, 0}.MultiplyMatrix(&worldMatrix).MultiplyMatrix(&viewMatrix).MultiplyMatrix(&projectionMatrix))
	// fmt.Printf("%+v\n\n", Vector{1, -1, 0}.MultiplyMatrix(&worldMatrix).MultiplyMatrix(&viewMatrix).MultiplyMatrix(&projectionMatrix))

	// render the model using shader
	g.model.Render()

	// present rendered scene to the screen
	g.opengl.EndScene()

	return nil
}
