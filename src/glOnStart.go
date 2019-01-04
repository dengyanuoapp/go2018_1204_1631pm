package main
import (
	//"encoding/binary"
	"log"

	//"golang.org/x/mobile/app"
	//"golang.org/x/mobile/event/lifecycle"
	//"golang.org/x/mobile/event/paint"
	//"golang.org/x/mobile/event/size"
	//"golang.org/x/mobile/event/touch"
	"golang.org/x/mobile/exp/app/debug"
	//"golang.org/x/mobile/exp/f32"
	"golang.org/x/mobile/exp/gl/glutil"
	"golang.org/x/mobile/gl"
)

func glOnStart(___glctx1 gl.Context) {
	var __err1 error
	_program, __err1 = glutil.CreateProgram(___glctx1, vertexShader, fragmentShader)
	if __err1 != nil {
		log.Printf("error creating GL program: %v", __err1)
		return
	}

	_buf = ___glctx1.CreateBuffer()
	___glctx1.BindBuffer(gl.ARRAY_BUFFER, _buf)
	___glctx1.BufferData(gl.ARRAY_BUFFER, triangleData, gl.STATIC_DRAW)

	_position = ___glctx1.GetAttribLocation(_program, "position")
	_color = ___glctx1.GetUniformLocation(_program, "color")
	_offset = ___glctx1.GetUniformLocation(_program, "offset")

	_images = glutil.NewImages(___glctx1)
	_fps = debug.NewFPS(_images)
} // glOnStart
