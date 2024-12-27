package main

import (
	"math"
	"strconv"
	"strings"
	"unicode"

	"github.com/hajimehoshi/ebiten/v2"
)

func BoardImage() *ebiten.Image {
  img := ebiten.NewImage(gridSize * 8, gridSize * 8)

	for x := 0; x < 8; x++ {
		for y := 0; y < 8; y++ {
			rect := ebiten.NewImage(gridSize, gridSize)
			drawOptions := ebiten.DrawImageOptions{}
			drawOptions.GeoM.Translate(float64(x*gridSize), float64(y*gridSize))
			if (x+y)%2 == 0 {
				rect.Fill(brightColor)
			} else {
				rect.Fill(darkColor)
			}
			img.DrawImage(rect, &drawOptions)
		}
	}
  return img
}

func BoardImageOptions() ebiten.DrawImageOptions {
  options := ebiten.DrawImageOptions{}
  options.GeoM.Translate(margins, margins)
  return options
}

func SetupBoard() {
	whiteMove = true
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
}

func GetPieceAt(pos [2]int) Piece {
  return Piece{
    currentState[pos[1]][pos[0]],
    pos,
  }
}

func SetPieceAtTo(p Piece) {
  pos := p.pos
	currentState[pos[1]][pos[0]] = p.pieceType
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
