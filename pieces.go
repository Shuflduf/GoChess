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
}

//go:embed assets/pieces_atlas_big.png
var textureData []byte
var nullPiece = Piece{
	0,
	[2]int{-1, -1},
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

func (p *Piece) ValidPositions() (valid [][2]int) {
	switch int(math.Abs(float64(p.pieceType))) {

	//King
	case 1:
		for i := -1; i < 2; i++ {
			for j := -1; j < 2; j++ {
				if i == 0 && j == 0 {
					continue
				}
				if p.pos[0]+i >= 0 && p.pos[0]+i < 8 && p.pos[1]+j >= 0 && p.pos[1]+j < 8 {
					valid = append(valid, [2]int{p.pos[0] + i, p.pos[1] + j})
				}
			}
		}

	//Queen
	case 2:
		for i := -1; i < 2; i++ {
			for j := -1; j < 2; j++ {
				if i == 0 && j == 0 {
					continue
				}
				for k := 1; k < 8; k++ {
					if p.pos[0]+i*k >= 0 && p.pos[0]+i*k < 8 && p.pos[1]+j*k >= 0 && p.pos[1]+j*k < 8 {
						valid = append(valid, [2]int{p.pos[0] + i*k, p.pos[1] + j*k})
						if GetPieceAt([2]int{p.pos[0] + i*k, p.pos[1] + j*k}).pieceType != 0 {
							break
						}
					} else {
						break
					}
				}
			}
		}

	//Bishop
	case 3:
		for i := -1; i < 2; i++ {
			for j := -1; j < 2; j++ {
				if i == 0 && j == 0 {
					continue
				}
				for k := 1; k < 8; k++ {
					if p.pos[0]+i*k >= 0 && p.pos[0]+i*k < 8 && p.pos[1]+j*k >= 0 && p.pos[1]+j*k < 8 {
						valid = append(valid, [2]int{p.pos[0] + i*k, p.pos[1] + j*k})
						if GetPieceAt([2]int{p.pos[0] + i*k, p.pos[1] + j*k}).pieceType != 0 {
							break
						}
					} else {
						break
					}
				}
			}
		}

	//Knight
	case 4:
		for i := -2; i < 3; i++ {
			for j := -2; j < 3; j++ {
				if i == 0 || j == 0 || i == j || i == -j {
					continue
				}
				if p.pos[0]+i >= 0 && p.pos[0]+i < 8 && p.pos[1]+j >= 0 && p.pos[1]+j < 8 {
					valid = append(valid, [2]int{p.pos[0] + i, p.pos[1] + j})
				}
			}
		}

	//Rook
	case 5:
		for i := -1; i < 2; i++ {
			for j := -1; j < 2; j++ {
				if i == 0 && j == 0 {
					continue
				}
				for k := 1; k < 8; k++ {
					if i != 0 && j != 0 {
						break
					}
					if p.pos[0]+i*k >= 0 && p.pos[0]+i*k < 8 && p.pos[1]+j*k >= 0 && p.pos[1]+j*k < 8 {
						valid = append(valid, [2]int{p.pos[0] + i*k, p.pos[1] + j*k})
						if GetPieceAt([2]int{p.pos[0] + i*k, p.pos[1] + j*k}).pieceType != 0 {
							break
						}
					} else {
						break
					}
				}
			}
		}

	//Pawn
	case 6:
		if p.pieceType < 0 {
			if p.pos[0] > 0 {
				if GetPieceAt([2]int{p.pos[0] - 1, p.pos[1] + 1}).pieceType > 0 {
					valid = append(valid, [2]int{p.pos[0] - 1, p.pos[1] + 1})
				}
			}
			if p.pos[0] < 7 {
				if GetPieceAt([2]int{p.pos[0] + 1, p.pos[1] + 1}).pieceType > 0 {
					valid = append(valid, [2]int{p.pos[0] + 1, p.pos[1] + 1})
				}
			}
			if p.pos[1] == 1 {
				valid = append(valid, [2]int{p.pos[0], p.pos[1] + 2})
			}
			if GetPieceAt([2]int{p.pos[0], p.pos[1] + 1}).pieceType == 0 {
				valid = append(valid, [2]int{p.pos[0], p.pos[1] + 1})
			}
		} else {
			if p.pos[0] > 0 {
				if GetPieceAt([2]int{p.pos[0] - 1, p.pos[1] - 1}).pieceType < 0 {
					valid = append(valid, [2]int{p.pos[0] - 1, p.pos[1] - 1})
				}
			}
			if p.pos[0] < 7 {
				if GetPieceAt([2]int{p.pos[0] + 1, p.pos[1] - 1}).pieceType < 0 {
					valid = append(valid, [2]int{p.pos[0] + 1, p.pos[1] - 1})
				}
			}
			if p.pos[1] == 6 {
				valid = append(valid, [2]int{p.pos[0], p.pos[1] - 2})
			}
			if GetPieceAt([2]int{p.pos[0], p.pos[1] - 1}).pieceType == 0 {
				valid = append(valid, [2]int{p.pos[0], p.pos[1] - 1})
			}
		}
	}
	return
}

func (p *Piece) IsTurn() bool {
	return (p.pieceType > 0 && whiteMove) || (p.pieceType < 0 && !whiteMove)
}

func (p *Piece) MovedTo(pos [2]int) Piece {
	new := *p
	new.pos = pos
	return new
}
