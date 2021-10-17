// Copyright@daidai53 2021
package proc

import (
	"image"
	"image/color"
	"image/jpeg"
	"io"
	"log"
	"os"
)

var ImageProcIns ImageProc

func init() {
	ImageProcIns = ImageProc{}
}

type ImageProc struct {
	imageFile io.Reader
	imageData image.Image
}

func (i *ImageProc) Close() {
}

func (i *ImageProc) LoadImage(fileName string) error {
	var err error
	i.imageFile, err = os.OpenFile(fileName, os.O_RDWR, 0777)
	if err != nil {
		return err
	}
	return nil
}

func (i *ImageProc) DecodeImage() error {
	var err error
	i.imageData, err = jpeg.Decode(i.imageFile)
	if err != nil {
		return err
	}
	return nil
}

func (i *ImageProc) ReadJpegAndConvertToGrayExample() {
	err := i.LoadImage("input.jpg")
	if err != nil {
		log.Default().Println("Error:open image file failed, exit. Cause: %v.", err)
		return
	}
	err = i.DecodeImage()
	if err != nil {
		log.Default().Println("Error:decode image file failed, exit. Cause: %v.", err)
		return
	}
	imageFileOut, err := os.OpenFile("output.jpg", os.O_RDWR, 0777)
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
	grayImage := image.NewGray(i.imageData.Bounds())
	for index := 0; index < i.imageData.Bounds().Dx(); index++ {
		for j := 0; j < i.imageData.Bounds().Dy(); j++ {
			colorPix := i.imageData.At(index, j)
			ycbcr, _ := colorPix.(color.YCbCr)
			r, g, b := color.YCbCrToRGB(ycbcr.Y, ycbcr.Cb, ycbcr.Cr)
			grayNum := float64(r)*0.299 + float64(g)*0.587 + float64(b)*0.114
			gray := color.Gray{
				Y: uint8(grayNum),
			}
			grayImage.SetGray(index, j, gray)
		}
	}
	err = jpeg.Encode(imageFileOut, grayImage, &option)
	if err != nil {
		log.Default().Printf("Error: encode failed, cause: %v", err)
		return
	}
}
