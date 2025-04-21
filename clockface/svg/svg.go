// Package svg produces an SVG clockface representation of a time.
package svg

import (
	"fmt"
	"io"
	"time"

	cf "github.com/suic86/learn-go-with-tests/clockface"
)

const (
	secondHandLength = 90
	minuteHandLength = 80
	hourHandLenght   = 50
	clockCenterX     = 150
	clockCenterY     = 150

	svgStart = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">
<svg xmlns="http://www.w3.org/2000/svg"
     width="100%"
     height="100%"
     viewBox="0 0 300 300"
     version="2.0">`

	bezel  = `<circle cx="150" cy="150" r="100" style="fill:#fff;stroke:#000;stroke-width:5px;"/>`
	svgEnd = `</svg>`
)

// Writer writes an SVG representation of an analogue clock, showing the time t, to the writer w
func Writer(w io.Writer, t time.Time) {
	io.WriteString(w, svgStart)
	io.WriteString(w, bezel)
	secondHand(w, t)
	minuteHand(w, t)
	hourHand(w, t)
	io.WriteString(w, svgEnd)
}

func makeHand(p cf.Point, length float64) cf.Point {
	// scale it to the length of the hand
	// flip it over the X axis to account for the SVG having an origin in the top left corner
	// translate it to the right position (so that it's coming from an origin of (clockCenterX, clockCenterY))
	p = cf.Point{X: p.X * length, Y: p.Y * length}                // scale
	p = cf.Point{X: p.X, Y: -p.Y}                                 // flip
	return cf.Point{X: p.X + clockCenterX, Y: p.Y + clockCenterY} // translate
}

func secondHand(w io.Writer, t time.Time) {
	p := makeHand(cf.SecondHandPoint(t), secondHandLength)
	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#f00;stroke-width:3px;"/>`, p.X, p.Y)
}

func minuteHand(w io.Writer, t time.Time) {
	p := makeHand(cf.MinuteHandPoint(t), minuteHandLength)
	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#000;stroke-width:3px;"/>`, p.X, p.Y)
}

func hourHand(w io.Writer, t time.Time) {
	p := makeHand(cf.HourHandPoint(t), hourHandLenght)
	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#000;stroke-width:3px;"/>`, p.X, p.Y)
}
