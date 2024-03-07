package main

import (
	"fmt"

	"github.com/zumosik/chess_engine_go/engine"
)

func main() {
	const x, y = 8, 8
	b := engine.NewBoard(x, y)

	rowsPieces := []engine.Piece{
		engine.TestPiece1, engine.TestPiece2,
	}

	for yi := 0; yi < 2; yi++ {
		for i := 0; i < x; i++ {
			b.InitializePiece(i, yi, rowsPieces[yi])
		}
	}

	for yi := y - 1; yi > y-3; yi-- {
		for i := 0; i < x; i++ {
			b.InitializePiece(i, yi, rowsPieces[y-yi-1])
		}
	}

	print(b)

}

func print(b engine.Board) {
	for i := 0; i < len(b.Celss); i++ {
		for j := 0; j < len(b.Celss[i]); j++ {

			cell := b.Celss[i][j]
			color := "w"
			if cell.Color == engine.Black {
				color = "b"
			}

			if cell.Piece == nil {
				fmt.Printf("%s ", color)
			} else {
				fmt.Printf("%s,%s,%s ", color, b.Celss[i][j].Piece.Name, b.Celss[i][j].Piece.Color.String())
			}

		}
		fmt.Println()

	}
}
