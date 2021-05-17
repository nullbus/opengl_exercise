package opengl_exercise

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/nullbus/opengl_exercise/gl"
)

type ColorShader struct {
	vertexShaderPath string
	pixelShaderPath  string

	vertexShader uint32
	pixelShader  uint32

	shaderProgram uint32
}

func NewColorShader() (*ColorShader, error) {
	shader := &ColorShader{
		vertexShaderPath: "../shaders/color.vs",
		pixelShaderPath:  "../shaders/color.ps",
	}

	return shader, shader.initializeShader()
}

func (s *ColorShader) Shutdown() {
	// detach the vertex and fragment shaders from the program
	gl.DetachShader(s.shaderProgram, s.vertexShader)
	gl.DetachShader(s.shaderProgram, s.pixelShader)

	// delete the vetex and fragment shaders
	gl.DeleteShader(s.vertexShader)
	gl.DeleteShader(s.pixelShader)

	// delete the shader program
	gl.DeleteProgram(s.shaderProgram)
}

func (s *ColorShader) SetShader() {
	// install the shader program as part of the current rendering state
	gl.UseProgram(s.shaderProgram)
}

func (s *ColorShader) SetShaderParams(worldMatrix, viewMatrix, projectionMatrix Matrix) error {
	// set the world matrix in the vertex shader
	if location, err := s.findUniform("worldMatrix"); err == nil {
		gl.UniformMatrix4fv(location, 1, false, worldMatrix.Ptr())
	} else {
		return err
	}

	// set the view matrix in the vetex shader
	if location, err := s.findUniform("viewMatrix"); err == nil {
		gl.UniformMatrix4fv(location, 1, false, viewMatrix.Ptr())
	} else {
		return err
	}

	// set the projection matrix in the vetex shader
	if location, err := s.findUniform("projectionMatrix"); err == nil {
		gl.UniformMatrix4fv(location, 1, false, projectionMatrix.Ptr())
	} else {
		return err
	}

	return nil
}

func (s *ColorShader) findUniform(name string) (int32, error) {
	if location := gl.GetUniformLocation(s.shaderProgram, &[]byte(name)[0]); location >= 0 {
		return location, nil
	}

	return -1, fmt.Errorf("failed to find uniform '%s'", name)
}

func (s *ColorShader) initializeShader() error {
	vertexShaderBuffer, err := ioutil.ReadFile(s.vertexShaderPath)
	if err != nil {
		return err
	}
	vertexShaderBufferPtr := &vertexShaderBuffer[0]
	vertexShaderBufferLen := int32(len(vertexShaderBuffer))

	fragmentShaderBuffer, err := ioutil.ReadFile(s.pixelShaderPath)
	if err != nil {
		return err
	}
	fragmentShaderBufferPtr := &fragmentShaderBuffer[0]
	fragmentShaderBufferLen := int32(len(fragmentShaderBuffer))

	// create a vertex and fragment shader object
	s.vertexShader = gl.CreateShader(gl.VERTEX_SHADER)
	s.pixelShader = gl.CreateShader(gl.FRAGMENT_SHADER)

	// copy the shader source code strings into the vertex and fragment shader objects
	gl.ShaderSource(s.vertexShader, 1, &vertexShaderBufferPtr, &vertexShaderBufferLen)
	gl.ShaderSource(s.pixelShader, 1, &fragmentShaderBufferPtr, &fragmentShaderBufferLen)

	// compile the shaders
	gl.CompileShader(s.vertexShader)
	gl.CompileShader(s.pixelShader)

	var compileStatus int32

	// check to see if the vertex shader compiled successfully
	gl.GetShaderiv(s.vertexShader, gl.COMPILE_STATUS, &compileStatus)
	if compileStatus != gl.TRUE {
		// if it did not compile then write the syntax error message out to a text file for review
		shaderErrorMessage(s.vertexShader)
		return errors.New("Shader build failed")
	}

	gl.GetShaderiv(s.pixelShader, gl.COMPILE_STATUS, &compileStatus)
	if compileStatus != gl.TRUE {
		// if it did not compile then write the syntax error message out to a text file for review
		shaderErrorMessage(s.pixelShader)
		return errors.New("Shader build failed")
	}

	// create a shader program object
	s.shaderProgram = gl.CreateProgram()

	// attach the vertex and fragment shader to the program object
	gl.AttachShader(s.shaderProgram, s.vertexShader)
	gl.AttachShader(s.shaderProgram, s.pixelShader)

	// bind the shader input variables
	s.bindAttrib(0, "inputPosition")
	s.bindAttrib(1, "inputColor")

	// link the shader program
	gl.LinkProgram(s.shaderProgram)

	// check the status of the link
	gl.GetProgramiv(s.shaderProgram, gl.LINK_STATUS, &compileStatus)
	if compileStatus != gl.TRUE {
		linkErrorMessage(s.shaderProgram)
		return errors.New("Shader linkage failed")
	}

	return nil
}

func (s *ColorShader) bindAttrib(idx uint32, name string) {
	gl.BindAttribLocation(s.shaderProgram, idx, &[]byte(name)[0])
}

func shaderErrorMessage(shaderID uint32) {
	// if it did not compile then write the syntax error message out to a text file for review
	var logSize int32
	gl.GetShaderiv(shaderID, gl.INFO_LOG_LENGTH, &logSize)

	// create char buffer to hold the info log
	infoLog := make([]byte, logSize)

	// retrieve the info log
	gl.GetShaderInfoLog(shaderID, logSize, nil, &infoLog[0])

	log.Println("shader build error:")
	log.Println(string(infoLog))
}

func linkErrorMessage(programID uint32) {
	// if it did not compile then write the syntax error message out to a text file for review
	var logSize int32
	gl.GetProgramiv(programID, gl.INFO_LOG_LENGTH, &logSize)

	// create char buffer to hold the info log
	infoLog := make([]byte, logSize)

	// retrieve the info log
	gl.GetProgramInfoLog(programID, logSize, nil, &infoLog[0])

	log.Println("shader link error:")
	log.Println(string(infoLog))
}
