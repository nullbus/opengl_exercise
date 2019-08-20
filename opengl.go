package opengl_exercise

import (
	"errors"
	"fmt"
	"io"
	"log"
	"math"
	"unsafe"

	"github.com/TheTitanrain/w32"
	"github.com/nullbus/opengl_exercise/gl"
	"github.com/nullbus/opengl_exercise/wgl"
)

type Matrix [16]float32

func (m1 *Matrix) Multiply(m2 *Matrix) Matrix {
	return Matrix{
		(*m1)[0]*(*m2)[0] + (*m1)[1]*(*m2)[4] + (*m1)[2]*(*m2)[8] + (*m1)[3]*(*m2)[12],
		(*m1)[0]*(*m2)[1] + (*m1)[1]*(*m2)[5] + (*m1)[2]*(*m2)[9] + (*m1)[3]*(*m2)[13],
		(*m1)[0]*(*m2)[2] + (*m1)[1]*(*m2)[6] + (*m1)[2]*(*m2)[10] + (*m1)[3]*(*m2)[14],
		(*m1)[0]*(*m2)[3] + (*m1)[1]*(*m2)[7] + (*m1)[2]*(*m2)[11] + (*m1)[3]*(*m2)[15],

		(*m1)[4]*(*m2)[0] + (*m1)[5]*(*m2)[4] + (*m1)[6]*(*m2)[8] + (*m1)[7]*(*m2)[12],
		(*m1)[4]*(*m2)[1] + (*m1)[5]*(*m2)[5] + (*m1)[6]*(*m2)[9] + (*m1)[7]*(*m2)[13],
		(*m1)[4]*(*m2)[2] + (*m1)[5]*(*m2)[6] + (*m1)[6]*(*m2)[10] + (*m1)[7]*(*m2)[14],
		(*m1)[4]*(*m2)[3] + (*m1)[5]*(*m2)[7] + (*m1)[6]*(*m2)[11] + (*m1)[7]*(*m2)[15],

		(*m1)[8]*(*m2)[0] + (*m1)[9]*(*m2)[4] + (*m1)[10]*(*m2)[8] + (*m1)[11]*(*m2)[12],
		(*m1)[8]*(*m2)[1] + (*m1)[9]*(*m2)[5] + (*m1)[10]*(*m2)[9] + (*m1)[11]*(*m2)[13],
		(*m1)[8]*(*m2)[2] + (*m1)[9]*(*m2)[6] + (*m1)[10]*(*m2)[10] + (*m1)[11]*(*m2)[14],
		(*m1)[8]*(*m2)[3] + (*m1)[9]*(*m2)[7] + (*m1)[10]*(*m2)[11] + (*m1)[11]*(*m2)[15],

		(*m1)[12]*(*m2)[0] + (*m1)[13]*(*m2)[4] + (*m1)[14]*(*m2)[8] + (*m1)[15]*(*m2)[12],
		(*m1)[12]*(*m2)[1] + (*m1)[13]*(*m2)[5] + (*m1)[14]*(*m2)[9] + (*m1)[15]*(*m2)[13],
		(*m1)[12]*(*m2)[2] + (*m1)[13]*(*m2)[6] + (*m1)[14]*(*m2)[10] + (*m1)[15]*(*m2)[14],
		(*m1)[12]*(*m2)[3] + (*m1)[13]*(*m2)[7] + (*m1)[14]*(*m2)[11] + (*m1)[15]*(*m2)[15],
	}
}

func (m *Matrix) Print(f io.Writer) {
	fmt.Fprintf(f, "%.2f %.2f %.2f %.2f\n", m[0], m[1], m[2], m[3])
	fmt.Fprintf(f, "%.2f %.2f %.2f %.2f\n", m[4], m[5], m[6], m[7])
	fmt.Fprintf(f, "%.2f %.2f %.2f %.2f\n", m[8], m[9], m[10], m[11])
	fmt.Fprintf(f, "%.2f %.2f %.2f %.2f\n\n", m[12], m[13], m[14], m[15])
}

type OpenGL struct {
	renderingContext w32.HGLRC
	deviceContext    w32.HDC
	worldMatrix      Matrix
	projectionMatrix Matrix

	videoCardDesc string
}

func NewOpenGL() *OpenGL {
	return &OpenGL{}
}

// InitializeExtensions is what allows us to get to get pointers to all the OpenGL functions that are available from the video card driver.
// In this function we will setup a temporary device context, pixel format, and rendering context so that we can call the LoadExtensionList function which then gets us all the function pointers we need.
func (o *OpenGL) InitializeExtensions(hwnd w32.HWND) error {
	// ge the device context for this window
	dc := w32.GetDC(hwnd)
	if dc == 0 {
		return errors.New("failed to retrieve device context")
	}

	// set a temporary default pixel format
	var pixelFormat w32.PIXELFORMATDESCRIPTOR
	if !w32.SetPixelFormat(dc, 1, &pixelFormat) {
		return errors.New("SetPixelFormat failed")
	}

	// create a temporary rendering context
	rc := w32.WglCreateContext(dc)
	if rc == 0 {
		return errors.New("failed to create rendering context")
	}

	// set the temporary rendering context as the current rendering context for this window
	if !w32.WglMakeCurrent(dc, rc) {
		return errors.New("MakeCurrent failed")
	}

	// initialize the opengl extensions needed for this application. Note that a temporary rendering context was needed to do so.
	if err := o.loadExtensionList(); err != nil {
		return err
	}

	// release the temporary rendering context now that the extensions have been loaded
	w32.WglMakeCurrent(0, 0)
	w32.WglDeleteContext(rc)

	// release the device context for this window
	w32.ReleaseDC(hwnd, dc)

	return nil
}

func (o *OpenGL) InitializeOpenGL(hwnd w32.HWND, screenWidth, screenHeight int, screenDepth, screenNear float32, vsync bool) error {

	// get device context for this window
	o.deviceContext = w32.GetDC(hwnd)
	if o.deviceContext == 0 {
		return errors.New("failed to get device context")
	}

	// support for opengl rendering
	attributeListInt := [...]int32{
		wgl.SUPPORT_OPENGL_ARB, gl.TRUE, // support for opengl rendering
		wgl.DRAW_TO_WINDOW_ARB, gl.TRUE, // support for rendering to window
		wgl.ACCELERATION_ARB, wgl.FULL_ACCELERATION_ARB, // support for hardware acceleration
		wgl.COLOR_BITS_ARB, 32, // support 32bit color
		wgl.ALPHA_BITS_ARB, 8, // support 8bit alpha
		wgl.DEPTH_BITS_ARB, 24, // support for 24bit depth buffer
		wgl.STENCIL_BITS_ARB, 8, // support for a 8bit stencil buffer
		wgl.DOUBLE_BUFFER_ARB, gl.TRUE, // support for double buffer
		wgl.SWAP_METHOD_ARB, wgl.SWAP_EXCHANGE_ARB, // support for swapping front and back buffer
		wgl.PIXEL_TYPE_ARB, wgl.TYPE_RGBA_ARB, // support for the RGBA pixel type
		0, // null terminate the attribute list
	}

	var pixelFormat [16]int32
	var formatCount uint32

	// query for a pixel format that fits the attributes we want
	result := wgl.ChoosePixelFormatARB(unsafe.Pointer(o.deviceContext), unsafe.Pointer(&attributeListInt[0]), nil, unsafe.Pointer(uintptr(len(pixelFormat))), unsafe.Pointer(&pixelFormat[0]), unsafe.Pointer(&formatCount))
	if uintptr(result) != gl.TRUE || formatCount == 0 {
		log.Println(result, pixelFormat, formatCount)
		return errors.New("wglChoosePixelFormatARB failed")
	}

	log.Println("formats:", formatCount)

	// if the video card/display can handle our desired pixel format then we set it as the current one
	var pixelFormatDescriptor w32.PIXELFORMATDESCRIPTOR
	result = wgl.SetPixelFormat(unsafe.Pointer(o.deviceContext), unsafe.Pointer(uintptr(pixelFormat[0])), unsafe.Pointer(&pixelFormatDescriptor))
	if uintptr(result) != gl.TRUE {
		log.Println(result, pixelFormat, formatCount)
		return errors.New("wglSetPixelFormat failed")
	}

	// set the 4.0 version of opengl in the attribute list
	attributeList := [...]int32{
		wgl.CONTEXT_MAJOR_VERSION_ARB, 4,
		wgl.CONTEXT_MINOR_VERSION_ARB, 0,
		0,
	}

	// create a opengl 4.0 rendering context
	o.renderingContext = w32.HGLRC(wgl.CreateContextAttribsARB(unsafe.Pointer(o.deviceContext), nil, unsafe.Pointer(&attributeList[0])))
	if o.renderingContext == 0 {
		return errors.New("Failed to create rendering context")
	}

	// set the rendering context to active
	result = wgl.MakeCurrent(unsafe.Pointer(o.deviceContext), unsafe.Pointer(o.renderingContext))
	if uintptr(result) != gl.TRUE {
		return errors.New("Failed wglMakeCurrent")
	}

	// set the depth buffer to be entirely cleared to 1.0 values
	gl.ClearDepth(1.0)

	// enable depth testing
	gl.Enable(gl.DEPTH_TEST)

	// set the polygon winding to front facing for the left-handed system
	gl.FrontFace(gl.CW)

	// enable back face culling
	gl.Enable(gl.CULL_FACE)
	gl.CullFace(gl.BACK)

	// initialize the world/model matrix to the identity matrix
	o.worldMatrix = o.identityMatrix()

	// set the field of view and screen aspect ratio
	fov := float32(math.Pi / 4.0)
	screenAspect := float32(screenWidth) / float32(screenHeight)

	// build the perspective projection matrix
	o.projectionMatrix = o.perspectiveFovLHMatrix(fov, screenAspect, screenNear, screenDepth)

	// get the name of the video card
	vendorString := gl.GoStr(gl.GetString(gl.VENDOR))
	rendererString := gl.GoStr(gl.GetString(gl.RENDERER))
	log.Println("vendor:", vendorString, "renderer:", rendererString)

	// store the video card anme in a class member variable so it can be retrieved later
	o.videoCardDesc = fmt.Sprintf("%s - %s", vendorString, rendererString)

	// turn on or off the vertical sync depending on the input bool variable
	if vsync {
		result = wgl.SwapIntervalEXT(unsafe.Pointer(uintptr(1)))
	} else {
		result = wgl.SwapIntervalEXT(unsafe.Pointer(uintptr(0)))
	}

	// check of vsync was set correctly
	if uintptr(result) != gl.TRUE {
		return errors.New("failed to set vsync")
	}

	return nil
}

func (o *OpenGL) WorldMatrix() Matrix {
	return o.worldMatrix
}

func (o *OpenGL) ProjectionMatrix() Matrix {
	return o.projectionMatrix
}

func (o *OpenGL) BeginScene(r, g, b, a float32) {
	// set the color to clear the scene to
	gl.ClearColor(r, g, b, a)
	check()

	// clear the screen and depth buffer
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	check()
}

func (o *OpenGL) EndScene() {
	// present the back buffer to the screen since rendering is complete
	w32.SwapBuffers(o.deviceContext)
	check()
}

func (o *OpenGL) Shutdown(hwnd w32.HWND) {
	// release rendering context
	if o.renderingContext != 0 {
		w32.WglMakeCurrent(0, 0)
		w32.WglDeleteContext(o.renderingContext)
		o.renderingContext = 0
	}

	// release the device context
	if o.deviceContext != 0 {
		w32.ReleaseDC(hwnd, o.deviceContext)
		o.deviceContext = 0
	}
}

func (o *OpenGL) loadExtensionList() error {
	if err := gl.Init(); err != nil {
		return errors.New("gl.Init error: " + err.Error())
	}

	if err := wgl.Init(); err != nil {
		return errors.New("wgl.Init error: " + err.Error())
	}

	return nil
}

func (o *OpenGL) identityMatrix() Matrix {
	return Matrix{
		1.0, 0, 0, 0,
		0, 1.0, 0, 0,
		0, 0, 1.0, 0,
		0, 0, 0, 1.0,
	}
}

func (o *OpenGL) perspectiveFovLHMatrix(fov, screenAspect, screenNear, screenDepth float32) Matrix {
	return Matrix{
		1.0 / (screenAspect * float32(math.Tan(float64(fov)*0.5))), 0, 0, 0,
		0, 1 / float32(math.Tan(float64(fov)*0.5)), 0, 0,
		0, 0, screenDepth / (screenDepth - screenNear), 1,
		0, 0, (-screenNear * screenDepth) / (screenDepth - screenNear), 0,
	}
}

func (o *OpenGL) rotationMatrix(angle float32) Matrix {
	return Matrix{
		float32(math.Cos(float64(angle))), 0, float32(-math.Sin(float64(angle))), 0,
		0, 1, 0, 0,
		float32(math.Sin(float64(angle))), 0, float32(math.Cos(float64(angle))), 0,
		0, 0, 0, 1,
	}
}

func (o *OpenGL) translationMatrix(x, y, z float32) Matrix {
	return Matrix{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		x, y, z, 1,
	}
}
