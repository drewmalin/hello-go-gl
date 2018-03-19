package render

import (
	"runtime"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

// Options !
type Options struct {
	Width  int
	Height int
	Title  string
}

// Loop !
func Loop(options *Options) {
	runtime.LockOSThread()

	window := newGlfwWindow(options)
	defer glfw.Terminate()

	program := newOpenGLProgram()

	for !window.ShouldClose() {
		render(window, program)
	}
}

func render(window *glfw.Window, program uint32) {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.UseProgram(program)

	glfw.PollEvents()
	window.SwapBuffers()
}

func newGlfwWindow(options *Options) *glfw.Window {
	if err := glfw.Init(); err != nil {
		panic(err)
	}

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	window, err := glfw.CreateWindow(options.Width, options.Height, options.Title, nil, nil)
	if err != nil {
		panic(err)
	}

	window.MakeContextCurrent()
	return window
}

func newOpenGLProgram() uint32 {
	if err := gl.Init(); err != nil {
		panic(err)
	}

	program := gl.CreateProgram()
	gl.LinkProgram(program)
	return program
}
