package opengl_exercise

import (
	"errors"
	"fmt"
	"log"
	"syscall"
	"time"
	"unsafe"

	"github.com/TheTitanrain/w32"
)

type System struct {
	Fullscreen  bool
	Vsync       bool
	ScreenDepth float32
	ScreenNear  float32

	opengl    *OpenGL
	input     *Input
	graphics  *Graphics
	hwnd      w32.HWND
	hinstance w32.HINSTANCE
}

func NewSystem(fullscreen, vsync bool) *System {
	return &System{
		Fullscreen:  fullscreen,
		Vsync:       vsync,
		ScreenDepth: 1000,
		ScreenNear:  0.1,
	}
}

func (s *System) Initialize() error {
	s.opengl = NewOpenGL()

	_, _, err := s.initializeWindows()
	if err != nil {
		return err
	}

	s.input, err = NewInput()
	if err != nil {
		return errors.New("creating input: " + err.Error())
	}

	s.graphics, err = NewGraphics(s.opengl, s.hwnd)
	if err != nil {
		return errors.New("creating graphics: " + err.Error())
	}

	return nil
}

func (s *System) Run() error {
	var message w32.MSG
	done := false
	lastFrame := time.Now()
	for !done {
		// Handle the windows messages
		if w32.PeekMessage(&message, 0, 0, 0, w32.PM_REMOVE) {
			w32.TranslateMessage(&message)
			w32.DispatchMessage(&message)
		}

		if message.Message == w32.WM_QUIT {
			// if windows signals to end the application then exit out.
			done = true
			continue

		}

		now := time.Now()

		// do the frame processing
		if err := s.frame(now.Sub(lastFrame)); err != nil {
			log.Println(err)
			done = true
		}

		lastFrame = now
	}

	return nil
}

func (s *System) Shutdown() {
	if s.graphics != nil {
		s.graphics.Shutdown()
		s.graphics = nil
	}

	if s.opengl != nil {
		s.opengl.Shutdown(s.hwnd)
		s.opengl = nil
	}

	s.shutdownWindow()
}

func (s *System) initializeWindows() (width, height int, err error) {
	// instance of this app
	s.hinstance = w32.GetModuleHandle("")
	if s.hinstance == 0 {
		return 0, 0, fmt.Errorf("failed to get module handle: %d", w32.GetLastError())
	}

	// name of this app
	applicationName := syscall.StringToUTF16Ptr("engine")
	windowName := syscall.StringToUTF16Ptr("Engine")

	wc := w32.WNDCLASSEX{
		Style:      w32.CS_HREDRAW | w32.CS_VREDRAW | w32.CS_OWNDC,
		WndProc:    syscall.NewCallback(s.windowProc),
		Instance:   s.hinstance,
		Icon:       w32.LoadIcon(0, w32.MakeIntResource(w32.IDI_WINLOGO)),
		IconSm:     w32.LoadIcon(0, w32.MakeIntResource(w32.IDI_WINLOGO)),
		Cursor:     w32.LoadCursor(0, w32.MakeIntResource(w32.IDC_ARROW)),
		Background: w32.HBRUSH(w32.GetStockObject(w32.BLACK_BRUSH)),
		ClassName:  applicationName,
		Size:       uint32(unsafe.Sizeof(w32.WNDCLASSEX{})),
	}

	// register the window class
	if ret := w32.RegisterClassEx(&wc); ret == 0 {
		return 0, 0, fmt.Errorf("Failed to register class: %d", w32.GetLastError())
	}

	// temporary window for the opengl extension setup
	s.hwnd = w32.CreateWindowEx(w32.WS_EX_APPWINDOW, applicationName, windowName, w32.WS_POPUP,
		0, 0, 640, 480, 0, 0, s.hinstance, nil)

	if s.hwnd == 0 {
		return 0, 0, fmt.Errorf("Failed to create window(1): %d", w32.GetLastError())
	}

	// don't show the window
	w32.ShowWindow(s.hwnd, w32.SW_HIDE)

	// initialize a temporary opengl window and load the opengl extensions
	if err := s.opengl.InitializeExtensions(s.hwnd); err != nil {
		return 0, 0, err
	}

	// release the temporary window now that the extensions have been initialized
	w32.DestroyWindow(s.hwnd)
	s.hwnd = 0

	// determine the resolution of the clients desktop screen
	width, height = w32.GetSystemMetrics(w32.SM_CXSCREEN), w32.GetSystemMetrics(w32.SM_CYSCREEN)

	// setup the screen settings depending on whether it is running in full screen or in windowed mode
	var posX, posY int
	if s.Fullscreen {
		dmScreenSettings := w32.DEVMODE{
			DmPelsWidth:  uint32(width),
			DmPelsHeight: uint32(height),
			DmBitsPerPel: 32,
			DmFields:     w32.DM_BITSPERPEL | w32.DM_PELSWIDTH | w32.DM_PELSHEIGHT,
			DmSize:       uint16(unsafe.Sizeof(w32.DEVMODE{})),
		}

		w32.ChangeDisplaySettingsEx(nil, &dmScreenSettings, s.hwnd, w32.CDS_FULLSCREEN, 0)
	} else {
		width, height = 800, 600

		// center
		posX = (w32.GetSystemMetrics(w32.SM_CXSCREEN) - width) / 2
		posY = (w32.GetSystemMetrics(w32.SM_CYSCREEN) - height) / 2
	}

	// create the window with the screen settings and get the handle to it
	s.hwnd = w32.CreateWindowEx(w32.WS_EX_APPWINDOW, applicationName, applicationName, w32.WS_POPUP,
		posX, posY, width, height, 0, 0, s.hinstance, nil)

	if s.hwnd == 0 {
		return 0, 0, fmt.Errorf("Failed to create window(2): %d", w32.GetLastError())
	}

	if err := s.opengl.InitializeOpenGL(s.hwnd, width, height, s.ScreenDepth, s.ScreenNear, s.Vsync); err != nil {
		return 0, 0, errors.New("Could not initialize opengl, check if video card supports opengl 4.0: " + err.Error())
	}

	// bring the window up on the screen and set it as main focus
	w32.ShowWindow(s.hwnd, w32.SW_SHOW)
	w32.SetForegroundWindow(s.hwnd)
	w32.SetFocus(s.hwnd)

	// hide the mouse cursor
	// w32.ShowCursor(false)
	return width, height, nil
}

func (s *System) shutdownWindow() {
	// show mouse cursor
	// w32.ShowCursor(true)

	if s.Fullscreen {
		// fix display settings if leaving full screen mode
		w32.ChangeDisplaySettingsEx(nil, nil, s.hwnd, 0, 0)
	}

	if s.hwnd != 0 {
		// remove the window
		w32.DestroyWindow(s.hwnd)
		s.hwnd = 0
	}
}

func (s *System) windowProc(hwnd syscall.Handle, msg uint32, wparam, lparam uintptr) uintptr {
	switch msg {
	case w32.WM_CLOSE:
		w32.PostQuitMessage(0)
		return 0

	case w32.WM_KEYDOWN:
		s.input.KeyDown(uint32(wparam))
		return 0

	case w32.WM_KEYUP:
		s.input.KeyUp(uint32(wparam))
		return 0

	default:
		return w32.DefWindowProc(w32.HWND(hwnd), msg, wparam, lparam)
	}
}

func (s *System) frame(delta time.Duration) error {
	// check if the user pressed escape and wants to exit the application
	if s.input.IsKeyDown(w32.VK_ESCAPE) {
		return errors.New("exit requested")
	}

	// do frame processing for the graphics object
	return s.graphics.Frame(delta)
}
