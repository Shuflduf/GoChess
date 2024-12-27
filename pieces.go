package main

import (
	"bytes"
	_ "embed"
	"image"
	_ "image/png"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Piece struct {
	pieceType int
	pos       [2]int
	symbol    rune
}

//go:embed assets/pieces_atlas_big.png
var textureData []byte
var nullPiece = Piece{
  0,
  [2]int{-1, -1},
  ' ',
}

var texture *ebiten.Image
var pieceSize int
var pieceMap = map[rune]int{
	'k': 1,
	'q': 2,
	'b': 3,
	'n': 4,
	'r': 5,
	'p': 6,
}

func GetPieceFromIndex(i int) *ebiten.Image {
	absI := int(math.Abs(float64(i)))
	cropRect := image.Rect((absI-1)*pieceSize, 0, absI*pieceSize, pieceSize)
	if i < 0 {
		cropRect = cropRect.Add(image.Point{0, pieceSize})
	}
	return texture.SubImage(cropRect).(*ebiten.Image)
}

func init() {
	img, _, err := ebitenutil.NewImageFromReader(bytes.NewReader(textureData))
	if err != nil {
		log.Fatal(err)
	}

	texture = img
	pieceSize = texture.Bounds().Dx() / 6
}

func ValidPositions(piece int, from [2]int) (valid [][2]int) {
	switch int(math.Abs(float64(piece))) {

	//Pawn
	case 6:
		if piece > 0 {
			t := from
			t[1]--
			valid = append(valid, t)
			if from[1] == 6 {
				t[1]--
				valid = append(valid, t)
			}
		} else {
			t := from
			t[1]++
			valid = append(valid, t)
			if from[1] == 1 {
				t[1]++
				valid = append(valid, t)
			}
		}
	}
	return
}
