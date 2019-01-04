package main

import (
	"encoding/binary"
	"log"

	//"golang.org/x/mobile/app"
	//"golang.org/x/mobile/event/lifecycle"
	//"golang.org/x/mobile/event/paint"
	//"golang.org/x/mobile/event/size"
	//"golang.org/x/mobile/event/touch"
	"golang.org/x/mobile/exp/app/debug"
	"golang.org/x/mobile/exp/f32"
	"golang.org/x/mobile/exp/gl/glutil"
	"golang.org/x/mobile/gl"

	//"time"
)

var _triangleDataByteARR = f32.Bytes(binary.LittleEndian,
0.0, 0.05, 0.0, // top left
0.0, 0.0, 0.0, // bottom left
0.05, 0.0, 0.0, // bottom right
) // _triangleDataByteARR

const _vertexShaderSTR = `#version 100
uniform vec2 offset;

attribute vec4 position;
void main() {
	// offset comes in with x/y values between 0 and 1.
	// position bounds are -1 to 1.
	vec4 offset4 = vec4(2.0*offset.x-1.0, 1.0-2.0*offset.y, 0, 0);
	gl_Position = position + offset4;
}`

const _fragmentShaderSTR = `#version 100
precision mediump float;
uniform vec4 color;
void main() {
	gl_FragColor = color;
}`

var (
	_dbFps      *debug.FPS
	_glImages   *glutil.Images
	_glProgram  gl.Program
	_glPosition gl.Attrib
	_glOffset   gl.Uniform
	_glColor    gl.Uniform
	_glBuf      gl.Buffer
)

func glOnStart(___glctx1 gl.Context) {
	var __err1 error
	_glProgram, __err1 = glutil.CreateProgram(___glctx1, _vertexShaderSTR, _fragmentShaderSTR)
	if __err1 != nil {
		log.Printf("error creating GL program: %v", __err1)
		return
	}
	test01() 

	_glBuf = ___glctx1.CreateBuffer() // CreateBuffer() Buffer  // CreateFramebuffer creates a framebuffer object.
	___glctx1.BindBuffer(gl.ARRAY_BUFFER, _glBuf) //  BindBuffer(target Enum, b Buffer)
	___glctx1.BufferData(gl.ARRAY_BUFFER, _triangleDataByteARR, gl.STATIC_DRAW) // BufferData(target Enum, src []byte, usage Enum)

	_glPosition = ___glctx1.GetAttribLocation(_glProgram, "position") // GetAttribLocation(p Program, name string) Attrib
	_glColor = ___glctx1.GetUniformLocation(_glProgram, "color")      // GetUniformLocation(p Program, name string) Uniform
	_glOffset = ___glctx1.GetUniformLocation(_glProgram, "offset")    // ...

	_glImages = glutil.NewImages(___glctx1)  // func (p *Images) NewImage(w, h int) *Image
	_dbFps = debug.NewFPS(_glImages)         // func NewFPS(images *glutil.Images) *FPS
} // glOnStart

func test01() {
	ggMain()
	//time.Sleep(150 * time.Second)
	textBoxMain()
	//time.Sleep(150 * time.Second)
}
