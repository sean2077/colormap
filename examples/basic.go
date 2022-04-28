package main

import (
	"fmt"

	"github.com/zhangxianbing/colormap"
)

func main() {
	fmt.Println(colormap.RGB("red"))
	fmt.Println(colormap.RGBA("red", 80))
}
