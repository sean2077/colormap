//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/goregular"
	"golang.org/x/image/math/fixed"

	"github.com/zhangxianbing/colormap"
)

func main() {
	fmt.Println(colormap.RGB("red"))
	fmt.Println(colormap.RGBA("red", 80))
	m := colormap.Map()
	tot := len(m)
	dh := 80
	h := dh * tot
	w := 1000
	img := image.NewRGBA(image.Rect(0, 0, w, h))
	i := 0
	for n, c := range m {
		for y := i * dh; y < (i+1)*dh; y++ {
			for x := 0; x < 100; x++ {
				img.Set(x, y, c)
			}
			addLabel(img, 150, i*dh+dh/2, n)
		}
		i++
	}
	file, err := os.Create("examples/output.png")
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()
	_ = png.Encode(file, img)
}

func addLabel(img draw.Image, x, y int, label string) {
	col := color.RGBA{R: 200, G: 100, A: 255}
	point := fixed.Point26_6{X: fixed.I(x), Y: fixed.I(y)}
	f, err := truetype.Parse(goregular.TTF)
	if err != nil {
		log.Fatal(err)
	}
	face := truetype.NewFace(f, &truetype.Options{Size: 48})

	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(col),
		Face: face,
		Dot:  point,
	}
	d.DrawString(label)
}
