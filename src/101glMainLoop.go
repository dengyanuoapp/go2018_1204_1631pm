package main

import (
	//"encoding/binary"
	//"log"

	"golang.org/x/mobile/app"
	"golang.org/x/mobile/event/lifecycle"
	"golang.org/x/mobile/event/paint"
	"golang.org/x/mobile/event/size"
	"golang.org/x/mobile/event/touch"
	//"golang.org/x/mobile/exp/app/debug"
	//"golang.org/x/mobile/exp/f32"
	//"golang.org/x/mobile/exp/gl/glutil"
	"golang.org/x/mobile/gl"
)

var (
	_touchX float32
	_touchY float32
)

func glMainLoop(___a app.App) {
	var __glctx gl.Context
	var __sz0 size.Event
	for __e01 := range ___a.Events() {
		switch __e02 := ___a.Filter(__e01).(type) {
		case lifecycle.Event:
			switch __e02.Crosses(lifecycle.StageVisible) {
			case lifecycle.CrossOn:
				__glctx, _ = __e02.DrawContext.(gl.Context)
				glOnStart(__glctx)
				___a.Send(paint.Event{})
			case lifecycle.CrossOff:
				glOnStop(__glctx)
				__glctx = nil
			}
		case size.Event:
			__sz0 = __e02
			_touchX = float32(__sz0.WidthPx / 2)
			_touchY = float32(__sz0.HeightPx / 2)
		case paint.Event:
			if __glctx == nil || __e02.External {
				// As we are actively painting as fast as
				// we can (usually 60 FPS), skip any paint
				// events sent by the system.
				continue
			}

			glOnPaint(__glctx, __sz0)
			___a.Publish()
			// Drive the animation by preparing to paint the next frame
			// after this one is shown.
			___a.Send(paint.Event{})
		case touch.Event:
			_touchX = __e02.X
			_touchY = __e02.Y
		} // switch __e02 := ___a.Filter(__e01).(type) {
	} // for __e01 := range ___a.Events() {
} // glMainLoop
