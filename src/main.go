// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build darwin linux windows

// An app that draws a green triangle on a red background.
//
// Note: This demo is an early preview of Go 1.5. In order to build this
// program as an Android APK using the gomobile tool.
//
// See http://godoc.org/golang.org/x/mobile/cmd/gomobile to install gomobile.
//
// Get the basic example and use gomobile to build or install it on your device.
//
//   $ go get -d golang.org/x/mobile/example/basic
//   $ gomobile build golang.org/x/mobile/example/basic # will build an APK
//
//   # plug your Android device to your computer or start an Android emulator.
//   # if you have adb installed on your machine, use gomobile install to
//   # build and deploy the APK to an Android target.
//   $ gomobile install golang.org/x/mobile/example/basic
//
// Switch to your device or emulator to start the Basic application from
// the launcher.
// You can also run the application on your desktop by running the command
// below. (Note: It currently doesn't work on Windows.)
//   $ go install golang.org/x/mobile/example/basic && basic
package main

import (
	"encoding/binary"
	"log"

	"golang.org/x/mobile/app"
	"golang.org/x/mobile/event/lifecycle"
	"golang.org/x/mobile/event/paint"
	"golang.org/x/mobile/event/size"
	"golang.org/x/mobile/event/touch"
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

func _appMainLoop(___a app.App) {
	var __glctx gl.Context
	var __sz0 size.Event
	for __e01 := range ___a.Events() {
		switch __e02 := ___a.Filter(__e01).(type) {
		case lifecycle.Event:
			switch __e02.Crosses(lifecycle.StageVisible) {
			case lifecycle.CrossOn:
				__glctx, _ = __e02.DrawContext.(gl.Context)
				onStart(__glctx)
				___a.Send(paint.Event{})
			case lifecycle.CrossOff:
				onStop(__glctx)
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

			onPaint(__glctx, __sz0)
			___a.Publish()
			// Drive the animation by preparing to paint the next frame
			// after this one is shown.
			___a.Send(paint.Event{})
		case touch.Event:
			_touchX = __e02.X
			_touchY = __e02.Y
		} // switch __e02 := ___a.Filter(__e01).(type) {
	} // for __e01 := range ___a.Events() {
} // _appMainLoop

func main() {
	app.Main(_appMainLoop) // app.Main( // func(___a app.App) {
} // main

func onStart(___glctx1 gl.Context) {
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
} // onStart

func onStop(___glctx2 gl.Context) {
	___glctx2.DeleteProgram(_program)
	___glctx2.DeleteBuffer(_buf)
	_fps.Release()
	_images.Release()
} // onStop

func onPaint(___glctx3 gl.Context, __sz3 size.Event) {
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
} // onPaint

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
