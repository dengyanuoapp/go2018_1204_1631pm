package main

import (
	"encoding/binary"
	//"log"

	"golang.org/x/mobile/app"
	//"golang.org/x/mobile/event/lifecycle"
	//"golang.org/x/mobile/event/paint"
	//"golang.org/x/mobile/event/size"
	//"golang.org/x/mobile/event/touch"
	"golang.org/x/mobile/exp/app/debug"
	"golang.org/x/mobile/exp/f32"
	"golang.org/x/mobile/exp/gl/glutil"
	"golang.org/x/mobile/gl"
)

var (
	_images   *glutil.Images
	_fps      *debug.FPS
	_program  gl.Program
	_position gl.Attrib
	_offset   gl.Uniform
	_color    gl.Uniform
	_buf      gl.Buffer

	_green  float32
	_touchX float32
	_touchY float32
)

var triangleData = f32.Bytes(binary.LittleEndian,
	0.0, 0.4, 0.0, // top left
	0.0, 0.0, 0.0, // bottom left
	0.4, 0.0, 0.0, // bottom right
) // triangleData

const (
	_coordsPerVertex = 3
	_vertexCount     = 3
)

const vertexShader = `#version 100
uniform vec2 offset;

attribute vec4 position;
void main() {
	// offset comes in with x/y values between 0 and 1.
	// position bounds are -1 to 1.
	vec4 offset4 = vec4(2.0*offset.x-1.0, 1.0-2.0*offset.y, 0, 0);
	gl_Position = position + offset4;
}`

const fragmentShader = `#version 100
precision mediump float;
uniform vec4 color;
void main() {
	gl_FragColor = color;
}`

func main() {
	app.Main(glMainLoop) // app.Main( // func(___a app.App) {
} // main
