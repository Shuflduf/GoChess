package main

import (
	"image/color"
	"log"
	"math"
	"slices"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct{}

const margins = 45
const startingFEN = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR"
var darkColor = color.RGBA{149, 68, 35, 255}
var brightColor = color.RGBA{220, 192, 144, 255}

var screenHeight int
var screenWidth int
var gridSize int
var currentState [8][8]int     // currentState[y][x]
var heldPiecePos = [2]int{-1, -1} // heldPiece{x, y}
var whiteMove = true

func (g *Game) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyF) {
		ebiten.SetFullscreen(!ebiten.IsFullscreen())
	}
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		clickGridPos := GetMouseGridPos()
		if clickGridPos != [2]int{-1, -1} {
			targetPiece := GetPieceAt(clickGridPos)
			if targetPiece.pieceType != 0 {
				if targetPiece.IsTurn() {
					heldPiecePos = clickGridPos
				}
			}
		}
	}

	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		if heldPiecePos != [2]int{-1, -1} {
			clickGridPos := GetMouseGridPos()

			if clickGridPos != [2]int{-1, -1} {
				targetPiece := GetPieceAt(clickGridPos)
				movingPiece := GetPieceAt(heldPiecePos)
				var canCapture = targetPiece.pieceType == 0
        if movingPiece == targetPiece {
          canCapture = false
        } else if movingPiece.pieceType > 0 {
					if targetPiece.pieceType < 0 {
						canCapture = true
					}
				} else {
					if targetPiece.pieceType > 0 {
						canCapture = true
					}
				}
        validPos := movingPiece.ValidPositions()
				if canCapture && slices.Contains(validPos, clickGridPos) {
          whiteMove = !whiteMove
          heldPiece := GetPieceAt(heldPiecePos)
					SetPieceAtTo(heldPiece.MovedTo(clickGridPos))
          SetPieceAtTo(Piece{ 0, heldPiecePos })
          // if king captured, reset
					if int(math.Abs(float64(targetPiece.pieceType))) == 1 {
						SetupBoard()
					}
				}
			}
			heldPiecePos = [2]int{-1, -1}
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{70, 70, 70, 255})
  opt := UIImageOptions()
  screen.DrawImage(UIImage(), &opt)

  opt = BoardImageOptions()
  screen.DrawImage(BoardImage(), &opt)

	// pieces
	drawOptions := ebiten.DrawImageOptions{}
	newScale := float64(gridSize) / float64(pieceSize)
	drawOptions.GeoM.Scale(newScale, newScale)
	for y, row := range currentState {
		for x, val := range row {
			updatedOptions := drawOptions
			if [2]int{x, y} == heldPiecePos {
				continue
			} else {
				updatedOptions.GeoM.Translate(
					margins+float64(gridSize*x),
					margins+float64(gridSize*y),
				)
			}
			screen.DrawImage(GetPieceFromIndex(val), &updatedOptions)
		}
	}
	if heldPiecePos != [2]int{-1, -1} {

		updatedOptions := drawOptions
		mousePosX, MousePosY := ebiten.CursorPosition()
		updatedOptions.GeoM.Translate(
			float64(mousePosX)-(float64(gridSize)/2.0),
			float64(MousePosY)-(float64(gridSize)/2.0),
		)
		screen.DrawImage(GetPieceFromIndex(GetPieceAt(heldPiecePos).pieceType), &updatedOptions)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	screenWidth = 1920
	screenHeight = 1080
	gridSize = (screenHeight - (margins * 2)) / 8
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowTitle("GoChess")
	SetupBoard()

	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
