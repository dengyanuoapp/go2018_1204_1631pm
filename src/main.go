package main

import (
	//"encoding/binary"
	//"log"

	"golang.org/x/mobile/app"
	//"golang.org/x/mobile/event/lifecycle"
	//"golang.org/x/mobile/event/paint"
	//"golang.org/x/mobile/event/size"
	//"golang.org/x/mobile/event/touch"
	"golang.org/x/mobile/exp/app/debug"
	//"golang.org/x/mobile/exp/f32"
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

func main() {
	app.Main(glMainLoop) // app.Main( // func(___a app.App) {
} // main
