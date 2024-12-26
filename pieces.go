package main

import (
	"bytes"
	_ "embed"
	"image"
	_ "image/png"
	"log"
	"unicode"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

//go:embed assets/pieces_atlas_big.png
var textureData []byte

var texture *ebiten.Image
var pieceSize int
var pieceMap = map[rune]int{
	'k': 0,
	'q': 1,
	'b': 2,
	'n': 3,
	'r': 4,
	'p': 5,
}

func GetPieceFromFen(r rune) *ebiten.Image {
	index := pieceMap[unicode.ToLower(r)]
	cropRect := image.Rect(index*pieceSize, 0, (index+1)*pieceSize, pieceSize)
	if unicode.IsLower(r) {
		cropRect = cropRect.Add(image.Point{0, pieceSize})
	}
	return texture.SubImage(cropRect).(*ebiten.Image)
}

func load_texture() {
	img, _, err := ebitenutil.NewImageFromReader(bytes.NewReader(textureData))
	if err != nil {
		log.Fatal(err)
	}

	texture = img
	pieceSize = texture.Bounds().Dx() / 6

}
