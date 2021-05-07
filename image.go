package main

import (
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
	"image"
	"image/color"
)

var (
	x int = 0
	y int = 9
	datax int = 37
)


func addLabel(img *image.RGBA, x, y int, label string) {
	col := color.RGBA{255, 255, 255, 255}
	point := fixed.Point26_6{fixed.Int26_6(x * 64), fixed.Int26_6(y * 64)}

	d := &font.Drawer{
			Dst:  img,
			Src:  image.NewUniform(col),
			Face: basicfont.Face7x13,
			Dot:  point,
	}
	d.DrawString(label)
}

func buildImage(data *Data) (*image.RGBA) {
	img := image.NewRGBA(image.Rect(0, 0, 128, 64))

	// Write text to image
	
	// IP
	addLabel(img, x, y, "IP:")
	addLabel(img, datax, y, data.IP)
  
	// CPU
	addLabel(img, x, y+18, "CPU:")
	addLabel(img, datax, y+18, data.CPU)
	
	// MEM
	addLabel(img, x, y+30, "Mem:")
	addLabel(img, datax, y+30, data.Mem)
  
	// TEMP
	addLabel(img, x, y+42, "Temp:")
	addLabel(img, datax, y+42, data.Temp)
  
	// SD
	addLabel(img, x, y+54, "SD:")
	addLabel(img, datax, y+54, data.SD)

	if data.SSD != "" {
		addLabel(img, x+68, y+54, "SSD:")
		addLabel(img, datax+60, y+54, data.SSD)
	}
	
	return img
}