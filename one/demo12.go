// 方法
package main

import (
	"bytes"
	"fmt"
	"image/color"
	"io"
	"os"
)

type Point struct {
	X, Y float64
}

type ColoredPoint struct {
	Point
	Color color.RGBA
}

const debug = true

func main() {
	twelveAA()
	twelveBB()
	twelveCC()
	twelveDD()

}

func twelveEE() {
	fmt.Println("twelveEE")

}

func twelveDD() {
	fmt.Println("twelveDD")
	var w io.Writer
	fmt.Printf("%T\n", w)
	w = os.Stdout
	fmt.Printf("%T\n", w)

	w.Write([]byte("heoo\n"))

	w = new(bytes.Buffer)
	fmt.Printf("%T\n", w)

	w.Write([]byte("hello")) // writes "hello" to the bytes.Buffers
	fmt.Println(w)
	w = nil

}

func twelveCC() {
	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}
	var p = ColoredPoint{Point{1, 2}, red}
	var q = ColoredPoint{Point{3, 4}, blue}
	p.ScaleBy(2)
	q.ScaleBy(2)
	fmt.Println("twelveCC")
	fmt.Println(p)
	fmt.Println(q)
}

func twelveBB() {
	var cp ColoredPoint
	cp.X = 1
	fmt.Println(cp.Point.X) // "1"
	fmt.Println(cp.X)       // "1"
	cp.Point.Y = 2
	fmt.Println(cp.Y) // "2"
}

func twelveAA() {
	r := Point{1, 2}
	r.ScaleBy(2)
	fmt.Printf("X = %v\n", r.X)
	fmt.Printf("Y = %v\n", r.Y)
}

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}
