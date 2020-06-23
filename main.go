package main

import (
	"image/png"
	"os"
	"strconv"
	"strings"

	"github.com/nfnt/resize"
)

func main() {
	// image size publish_service

	if len(os.Args) < 3 {
		panic("Use template: <input image> <size (WxH)> [output file name]")
	}
	i := os.Args[1]
	s := os.Args[2]
	o := os.Args[3]
	size := strings.Split(s, "x")
	width, _ := strconv.Atoi(size[0])
	height, _ := strconv.Atoi(size[1])

	file, _ := os.Open(i)
	img, _ := png.Decode(file)

	m := resize.Resize(uint(width), uint(height), img, resize.Lanczos3)

	out, _ := os.Create(o + ".png")
	defer out.Close()
	png.Encode(out, m)
}
