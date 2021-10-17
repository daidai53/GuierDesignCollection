// Copyright@daidai53 2021
package imagereader

import (
	"image"
	"image/color"
)

type JpegReader interface {
	ImageReader
	image.Image
	String() string
}

type JpegPic struct {
	name string
}

func (j JpegPic) ReadByFile(in interface{}) image.Image {
	return nil
}

func (j JpegPic) ColorModel() color.Model {
	return nil
}

func (j JpegPic) Bounds() image.Rectangle {
	return image.Rectangle{
		Min: image.Point{},
		Max: image.Point{},
	}
}
func (j JpegPic) At(x int, y int) color.Color {
	return nil
}

func (j JpegPic) String() string {
	return "[image](type: jpeg, name: " + j.name + ",size: default)"
}

func (j *JpegPic) SetName(n string) error {
	j.name = n
	return nil
}
