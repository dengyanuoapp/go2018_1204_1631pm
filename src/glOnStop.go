package main

import (
	//"encoding/binary"
	//"log"

	//"golang.org/x/mobile/app"
	//"golang.org/x/mobile/event/lifecycle"
	//"golang.org/x/mobile/event/paint"
	//"golang.org/x/mobile/event/size"
	//"golang.org/x/mobile/event/touch"
	//"golang.org/x/mobile/exp/app/debug"
	//"golang.org/x/mobile/exp/f32"
	//"golang.org/x/mobile/exp/gl/glutil"
	"golang.org/x/mobile/gl"
)

func glOnStop(___glctx2 gl.Context) {
	___glctx2.DeleteProgram(_program)
	___glctx2.DeleteBuffer(_buf)
	_fps.Release()
	_images.Release()
} // glOnStop
