package opengl_exercise

import (
	"unsafe"

	"github.com/nullbus/opengl_exercise/gl"
)

type Vertex struct {
	X, Y, Z float32
	R, G, B float32
}

type Model struct {
	vertices []Vertex
	indices  []uint32

	vertexArray  uint32
	vertexBuffer uint32
	indexBuffer  uint32
}

func NewModel() (*Model, error) {
	model := &Model{}

	// initialize the vetex and index buffer and that hold the geometry fot the triangle
	return model, model.initialize()
}

func (m *Model) initialize() error {
	// load vertex array with data
	m.vertices = []Vertex{
		{
			-1, -1, 0,
			0, 1, 0,
		},
		{
			0, 1, 0,
			0, 1, 0,
		},
		{
			1, -1, 0,
			0, 1, 0,
		},
	}

	// load index array with data
	m.indices = []uint32{
		0, // bottom left
		1, // top middle
		2, // bottom right
	}

	// allocate opengl vertex array object
	gl.GenVertexArrays(1, &m.vertexArray)

	// bind vertex array object to store all the buffers and vertex attributes we create here
	gl.BindVertexArray(m.vertexArray)

	// gerenete an id for the vertex buffer
	gl.GenBuffers(1, &m.vertexBuffer)

	// bind vertex buffer and load the vertex data into the vertex buffer
	gl.BindBuffer(gl.ARRAY_BUFFER, m.vertexBuffer)
	gl.BufferData(gl.ARRAY_BUFFER, len(m.vertices)*int(unsafe.Sizeof(Vertex{})), unsafe.Pointer(&m.vertices[0]), gl.STATIC_DRAW)

	// enable two vertex array attributes
	gl.EnableVertexAttribArray(0) // vertex position
	gl.EnableVertexAttribArray(1) // vertex color

	// specify the location and format of the position portion of the vertex buffer
	stride := int32(unsafe.Sizeof(Vertex{}))
	gl.BindBuffer(gl.ARRAY_BUFFER, m.vertexBuffer)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, stride, nil)

	// specify the location and format of the color portion of the vertex buffer
	gl.BindBuffer(gl.ARRAY_BUFFER, m.vertexBuffer)
	gl.VertexAttribPointer(1, 3, gl.FLOAT, false, stride, unsafe.Pointer(uintptr(3*4 /*sizeof(float32)*/)))

	// generate an id for the index buffer
	gl.GenBuffers(1, &m.indexBuffer)

	// bind index buffer and load the index data to it
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, m.indexBuffer)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(m.indices)*4 /*sizeof(uint32)*/, unsafe.Pointer(&m.indices[0]), gl.STATIC_DRAW)

	return nil
}

func (m *Model) Shutdown() {
	// disable two vertex array attributes
	gl.DisableVertexAttribArray(0)
	gl.DisableVertexAttribArray(1)

	// release vertex buffer
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
	gl.DeleteBuffers(1, &m.vertexBuffer)

	// release index buffer
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, 0)
	gl.DeleteBuffers(1, &m.indexBuffer)

	// release vertex array object
	gl.BindVertexArray(0)
	gl.DeleteVertexArrays(1, &m.vertexArray)
}

func (m *Model) Render() {
	m.renderBuffers()
}

func (m *Model) renderBuffers() {
	// bind the vertex array object that stored all the information about the vetex and index buffers
	gl.BindVertexArray(m.vertexArray)

	// render vertex buffer using the index buffer
	gl.DrawElements(gl.TRIANGLES, int32(len(m.indices)), gl.UNSIGNED_INT, nil)
}
