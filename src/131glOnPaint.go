package main

import (
	//"encoding/binary"
	//"log"

	//"golang.org/x/mobile/app"
	//"golang.org/x/mobile/event/lifecycle"
	//"golang.org/x/mobile/event/paint"
	"golang.org/x/mobile/event/size"
	//"golang.org/x/mobile/event/touch"
	//"golang.org/x/mobile/exp/app/debug"
	//"golang.org/x/mobile/exp/f32"
	//"golang.org/x/mobile/exp/gl/glutil"
	"golang.org/x/mobile/gl"
)

const (
	_coordsPerVertex = 3
	_vertexCount     = 3
)

var (
	_green float32
)

func glOnPaint(___glctx3 gl.Context, __sz3 size.Event) {
	___glctx3.ClearColor(1, 0, 0, 1) // ClearColor(red, green, blue, alpha float32) // ClearColor specifies the RGBA values used to clear color buffers.
	___glctx3.Clear(gl.COLOR_BUFFER_BIT) // Clear(mask Enum) // // Clear clears the window.

	___glctx3.UseProgram(_glProgram) // UseProgram(p Program) // UseProgram sets the active program

	_green += 0.01
	if _green > 1 {
		_green = 0
	}
	___glctx3.Uniform4f(_glColor, 0, _green, 0, 1) // Uniform4f(dst Uniform, v0, v1, v2, v3 float32) // writes a vec4 uniform variable.

	___glctx3.Uniform2f(_glOffset, _touchX/float32(__sz3.WidthPx), _touchY/float32(__sz3.HeightPx)) // Uniform2f(dst Uniform, v0, v1 float32)

	___glctx3.BindBuffer(gl.ARRAY_BUFFER, _glBuf)
	___glctx3.EnableVertexAttribArray(_glPosition)
	___glctx3.VertexAttribPointer(_glPosition, _coordsPerVertex, gl.FLOAT, false, 0, 0)
	___glctx3.DrawArrays(gl.TRIANGLES, 0, _vertexCount)
	___glctx3.DisableVertexAttribArray(_glPosition)

	_dbFps.Draw(__sz3)
} // glOnPaint
