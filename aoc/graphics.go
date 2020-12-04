package aoc

import (
	"fmt"
	"image"
	"image/color"
	"image/color/palette"
	"image/draw"
	"image/gif"
	"image/png"
	"log"
	"os"
)

// SavePNG from an image
func SavePNG(filename string, i image.Image) {
	f, err := os.Create(filename)
	if err != nil {
		log.Fatalf("os.Create() failed with %s\n", err)
	}
	defer f.Close()

	err = png.Encode(f, i)
	if err != nil {
		log.Fatalf("png.Encode() failed with %s\n", err)
	}
}

// PaletteImageFromImage return a paletted image copied from a regular image
func PaletteImageFromImage(img image.Image) (palleted *image.Paletted) {
	palleted = image.NewPaletted(img.Bounds(), palette.Plan9)
	draw.FloydSteinberg.Draw(palleted, img.Bounds(), img, image.ZP)
	return
}

// SaveGIF from a single palette image
func SaveGIF(filename string, img *image.Paletted) {
	SaveGIFs("text.gif", []*image.Paletted{img}, 0)
}

// SaveGIFs from a series of paletted images
func SaveGIFs(filename string, images []*image.Paletted, delay int) {
	f, err := os.Create(filename)
	if err != nil {
		log.Fatalf("os.Create() failed with %s\n", err)
	}
	defer f.Close()

	anim := gif.GIF{Image: images}

	for range images {
		anim.Delay = append(anim.Delay, delay)
	}

	err = gif.EncodeAll(f, &anim)
	if err != nil {
		log.Fatalf("gif.Encode() failed with %s\n", err)
	}
}

// ParseHexColor parses a webcolor string
// from https://stackoverflow.com/questions/54197913/parse-hex-string-to-image-color
func ParseHexColor(s string) (c color.RGBA, err error) {
	c.A = 0xff
	switch len(s) {
	case 7:
		_, err = fmt.Sscanf(s, "#%02x%02x%02x", &c.R, &c.G, &c.B)
	case 4:
		_, err = fmt.Sscanf(s, "#%1x%1x%1x", &c.R, &c.G, &c.B)
		// Double the hex digits:
		c.R *= 17
		c.G *= 17
		c.B *= 17
	default:
		err = fmt.Errorf("invalid length, must be 7 or 4")
	}
	return
}

// future: text version of graphic image
// https://golang.org/pkg/image/draw/#Quantizer
