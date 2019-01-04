package main

import (
    "fmt"
    "image"
    "os"

	//"golang.org/x/mobile/app"
    "golang.org/x/exp/shiny/text"
    "golang.org/x/image/font"
    "golang.org/x/image/math/fixed"

	/*
	"bytes"
	"errors"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
	"io"
	"strings"
	"unicode/utf8"
	*/

	//"time"
)

// _toyFace implements the font.Face interface by measuring every rune's width
// as 1 pixel.
type _toyFace struct{}

func (_toyFace) Close() error {
    return nil
}

func (_toyFace) Glyph(dot fixed.Point26_6, r rune) (image.Rectangle, image.Image, image.Point, fixed.Int26_6, bool) {
    panic("unimplemented 183811")
}

func (_toyFace) GlyphBounds(r rune) (fixed.Rectangle26_6, fixed.Int26_6, bool) {
    panic("unimplemented 183812")
}

func (_toyFace) GlyphAdvance(r rune) (fixed.Int26_6, bool) {
    return fixed.I(1), true
}

func (_toyFace) Kern(r0, r1 rune) fixed.Int26_6 {
    return 0
}

func (_toyFace) Metrics() font.Metrics {
    return font.Metrics{}
}

func _printFrame(f *text.Frame, softReturnsOnly bool) {
    for p := f.FirstParagraph(); p != nil; p = p.Next(f) {
        for l := p.FirstLine(f); l != nil; l = l.Next(f) {
            for b := l.FirstBox(f); b != nil; b = b.Next(f) {
                if softReturnsOnly {
                    os.Stdout.Write(b.TrimmedText(f))
                } else {
                    os.Stdout.Write(b.Text(f))
                }
            }
            if softReturnsOnly {
                fmt.Println()
            }
        }
    }
}

func textBoxMain() {
    var f text.Frame
    f.SetFace(_toyFace{})
    f.SetMaxWidth(fixed.I(60))

    c := f.NewCaret()
    c.WriteString(_mobyDick)
    c.Close()


    fmt.Println("====")
    _printFrame(&f, false)

    fmt.Println("====")
    fmt.Println("123456789_123456789_123456789_123456789_123456789_123456789_")
    _printFrame(&f, true)


	//time.Sleep(150 * time.Second)

    fmt.Println("====")

}

const _mobyDick = "CHAPTER 1. Loomings.\nCall me Ishmael. Some years ago—never mind how long precisely—having little or no money in my purse, and nothing particular to interest me on shore, I thought I would sail about a little and see the watery part of the world...\n"
