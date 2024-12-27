package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

var defaultFont text.Face

func init() {
	tt, err := opentype.Parse(fonts.PressStart2P_ttf)
	if err != nil {
		log.Fatal(err)
	}
	const dpi = 72
	fontFace, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    24,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}

	defaultFont = text.NewGoXFace(fontFace)
}

func UIImage() *ebiten.Image {
  img := ebiten.NewImage(screenWidth - (margins * 3) - (gridSize * 8), gridSize)
  textOptions := text.DrawOptions{}
  textOptions.PrimaryAlign = text.AlignCenter
  textOptions.SecondaryAlign = text.AlignCenter
  textOptions.GeoM.Translate(float64(img.Bounds().Dx()) / 2, float64(img.Bounds().Dy()) / 2)
  if whiteMove {
    img.Fill(color.RGBA{255, 255, 255, 255})
    textOptions.ColorScale.SetR(0)
    textOptions.ColorScale.SetG(0)
    textOptions.ColorScale.SetB(0)
    text.Draw(img, "White's Turn", defaultFont, &textOptions)
  } else {
    img.Fill(color.RGBA{0, 0, 0, 255})
    textOptions.ColorScale.SetR(255)
    textOptions.ColorScale.SetG(255)
    textOptions.ColorScale.SetB(255)
    text.Draw(img, "Black's Turn", defaultFont, &textOptions)
  }
  return img
}

func UIImageOptions() ebiten.DrawImageOptions {
  options := ebiten.DrawImageOptions{}
  options.GeoM.Translate(
    float64((margins * 2) + (gridSize * 8)),
    float64(margins + (gridSize * 7)),
  )
  return options
}
