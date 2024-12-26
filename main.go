package main

import (
	"image/color"
	"log"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
)


type Game struct{}

const margins = 45
const startingFEN = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR"

var darkColor = color.RGBA{149, 68, 35, 255}
var brightColor = color.RGBA{220, 192, 144, 255}
var currentState [8][8]int8 // currentState[y][x]


func SetupBoard() {
  parts := strings.Split(startingFEN, "/")
  for row, i := range parts {
    
  }
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{70, 70, 70, 255})
	// board
	gridSize := (720 - (margins * 2)) / 8
	for x := 0; x < 8; x++ {
		for y := 0; y < 8; y++ {
			rect := ebiten.NewImage(gridSize, gridSize)
			drawOptions := ebiten.DrawImageOptions{}
			drawOptions.GeoM.Translate(float64(x*gridSize+margins), float64(y*gridSize+margins))
			if (x+y)%2 == 0 {
				rect.Fill(brightColor)
			} else {
				rect.Fill(darkColor)
			}
			screen.DrawImage(rect, &drawOptions)
		}
	}

  // pieces
  drawOptions := ebiten.DrawImageOptions{}
  newScale := float64(gridSize) / float64(pieceSize)
  drawOptions.GeoM.Scale(newScale, newScale)
  drawOptions.GeoM.Translate(margins, margins)
  screen.DrawImage(GetPieceFromFen('Q'), &drawOptions)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1280, 720
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowTitle("Hello, World!")
  load_texture()

	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
