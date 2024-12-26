package main

import (
	"image/color"
	"log"
	"strconv"
	"strings"
	"unicode"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct{}

const margins = 45
const startingFEN = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR"

var darkColor = color.RGBA{149, 68, 35, 255}
var brightColor = color.RGBA{220, 192, 144, 255}
var gridSize = (720 - (margins * 2)) / 8
var currentState [8][8]int // currentState[y][x]

func SetupBoard() {
	parts := strings.Split(startingFEN, "/")
	for i, row := range parts {
		var j = 0
		for _, char := range row {
			if char >= '1' && char <= '8' {
				empty, _ := strconv.Atoi(string(char))
				for k := 0; k < empty; k++ {
					currentState[i][j] = 0
					j++
				}
			} else {
				sign := !unicode.IsUpper(char)
				currentState[i][j] = pieceMap[rune(unicode.ToLower(char))]
				if sign {
					currentState[i][j] *= -1
				}
				j++
			}
		}
	}
	log.Println(currentState)
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{70, 70, 70, 255})

	// board
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
	for y, row := range currentState {
		for x, val := range row {
			updatedOptions := drawOptions
			updatedOptions.GeoM.Translate(
				margins+float64(gridSize*x),
				margins+float64(gridSize*y),
			)
			screen.DrawImage(GetPieceFromIndex(val), &updatedOptions)
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
  screenWidth = 1920
  screenHeight = 1080
  gridSize = (screenHeight - (margins * 2)) / 8
	return
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowTitle("Hello, World!")
	LoadTexture()
	SetupBoard()

	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
