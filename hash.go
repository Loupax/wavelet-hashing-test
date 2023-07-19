package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path"
	"strings"

	"github.com/Loupax/imagehash"
)

func main() {
	file, _ := os.Open(os.Args[1])
	defer file.Close()

	var img image.Image
	ext := strings.ToLower(path.Ext(os.Args[1]))
	switch ext {
	case ".png":
		img, _ = png.Decode(file)
	case ".jpg", ".jpeg":
		img, _ = jpeg.Decode(file)
	default:
		panic(fmt.Sprintf("unsupported extension %s", ext))
	}
	hash := imagehash.Imagehash{}
	hash.Whash(img, 8)
	fmt.Println(hash.String())
}
