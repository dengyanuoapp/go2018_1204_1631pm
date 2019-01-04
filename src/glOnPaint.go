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

func glOnPaint(___glctx3 gl.Context, __sz3 size.Event) {
	___glctx3.ClearColor(1, 0, 0, 1)
	___glctx3.Clear(gl.COLOR_BUFFER_BIT)

	___glctx3.UseProgram(_program)

	_green += 0.01
	if _green > 1 {
		_green = 0
	}
	___glctx3.Uniform4f(_color, 0, _green, 0, 1)

	___glctx3.Uniform2f(_offset, _touchX/float32(__sz3.WidthPx), _touchY/float32(__sz3.HeightPx))

	___glctx3.BindBuffer(gl.ARRAY_BUFFER, _buf)
	___glctx3.EnableVertexAttribArray(_position)
	___glctx3.VertexAttribPointer(_position, _coordsPerVertex, gl.FLOAT, false, 0, 0)
	___glctx3.DrawArrays(gl.TRIANGLES, 0, _vertexCount)
	___glctx3.DisableVertexAttribArray(_position)

	_fps.Draw(__sz3)
} // glOnPaint
