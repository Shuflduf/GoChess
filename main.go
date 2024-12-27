package main

import (
	"image/color"
	"log"
	"math"
	"strconv"
	"strings"
	"unicode"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct{}

const margins = 45
const startingFEN = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR"

var darkColor = color.RGBA{149, 68, 35, 255}
var brightColor = color.RGBA{220, 192, 144, 255}
var gridSize = (720 - (margins * 2)) / 8
var currentState [8][8]int     // currentState[y][x]
var heldPiece = [2]int{-1, -1} // heldPiece{x, y}

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
	if inpututil.IsKeyJustPressed(ebiten.KeyF) {
		ebiten.SetFullscreen(!ebiten.IsFullscreen())
	}
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		clickGridPos := GetMouseGridPos()
		if clickGridPos != [2]int{-1, -1} {
			if GetPieceAt(clickGridPos) != 0 {
				heldPiece = clickGridPos
			}
			log.Println(clickGridPos)
		}
	}

	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		if heldPiece != [2]int{-1, -1} {
			clickGridPos := GetMouseGridPos()

			if clickGridPos != [2]int{-1, -1} {
				// TODO make capturing and stuff
				targetPiece := GetPieceAt(clickGridPos)
				movingPiece := GetPieceAt(heldPiece)
				var canCapture = targetPiece == 0
				if movingPiece > 0 {
					if targetPiece < 0 {
						canCapture = true
					}
				} else {
					if targetPiece > 0 {
						canCapture = true
					}
				}
				if canCapture {
					SetPieceAtTo(clickGridPos, GetPieceAt(heldPiece))
					SetPieceAtTo(heldPiece, 0)
				}
			}
			heldPiece = [2]int{-1, -1}
		}
	}
	return nil
}

func GetPieceAt(pos [2]int) int {
	return currentState[pos[1]][pos[0]]
}

func SetPieceAtTo(pos [2]int, to int) {
	currentState[pos[1]][pos[0]] = to
}

func GetMouseGridPos() [2]int {
	mousePosX, mousePosY := ebiten.CursorPosition()
	mousePos := [2]int{
		int(math.Floor(float64((mousePosX - margins) / gridSize))),
		int(math.Floor(float64((mousePosY - margins) / gridSize))),
	}
	if mousePos[0] >= 0 && mousePos[0] < 8 && mousePos[1] >= 0 && mousePos[1] < 8 {
		return mousePos
	}
	return [2]int{-1, -1}
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
			if [2]int{x, y} == heldPiece {
				continue
				// mousePosX, MousePosY := ebiten.CursorPosition()
				// updatedOptions.GeoM.Translate(
				// 	float64(mousePosX) - (float64(gridSize) / 2.0),
				// 	float64(MousePosY) - (float64(gridSize) / 2.0),
				// )
			} else {
				updatedOptions.GeoM.Translate(
					margins+float64(gridSize*x),
					margins+float64(gridSize*y),
				)
			}
			screen.DrawImage(GetPieceFromIndex(val), &updatedOptions)
		}
	}
	if heldPiece != [2]int{-1, -1} {

		updatedOptions := drawOptions
		mousePosX, MousePosY := ebiten.CursorPosition()
		updatedOptions.GeoM.Translate(
			float64(mousePosX)-(float64(gridSize)/2.0),
			float64(MousePosY)-(float64(gridSize)/2.0),
		)
		screen.DrawImage(GetPieceFromIndex(GetPieceAt(heldPiece)), &updatedOptions)
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
