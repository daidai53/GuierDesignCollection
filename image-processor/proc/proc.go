// Copyright@daidai53 2021
package proc

import (
	"image"
	"image/color"
	"image/jpeg"
	"log"
	"os"
)

var ImageProcIns ImageProc

func init() {
	ImageProcIns = ImageProc{}
}

type ImageProc struct {
}

func (i ImageProc) ReadJpegAndConvertToGrayExample() {
	imageFile, err := os.Open("input.jpg")
	if err != nil {
		log.Default().Println("Error:open image file failed, exit.")
		return
	}
	imageData, err := jpeg.Decode(imageFile)
	if err != nil {
		log.Default().Println("Error:decode image file failed, exit.")
		return
	}
	imageFileOut, err := os.Open("output.jpg")
	if err != nil {
		imageFileOut, err = os.Create("output.jpg")
		if err != nil {
			log.Default().Println("Error:create output.jpg failed, exit.")
			return
		}
	}
	option := jpeg.Options{
		Quality: 100,
	}
	grayImage := image.NewGray(imageData.Bounds())
	for i := 0; i < imageData.Bounds().Dx(); i++ {
		for j := 0; j < imageData.Bounds().Dy(); j++ {
			colorPix := imageData.At(i, j)
			ycbcr, _ := colorPix.(color.YCbCr)
			r, g, b := color.YCbCrToRGB(ycbcr.Y, ycbcr.Cb, ycbcr.Cr)
			grayNum := float64(r)*0.299 + float64(g)*0.587 + float64(b)*0.114
			gray := color.Gray{
				Y: uint8(grayNum),
			}
			grayImage.SetGray(i, j, gray)
		}
	}
	jpeg.Encode(imageFileOut, grayImage, &option)
}
