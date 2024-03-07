package engine

type Color int

const (
	Black Color = iota
	White
)

type Board struct {
	Celss [][]Cell
}

type Cell struct {
	Piece *Piece
	Color Color
}

type Piece struct {
	Name  string
	Color Color
	Moves []Move
}

func (c Color) String() string {
	if c == Black {
		return "b"
	} else if c == White {
		return "w"
	}

	return ""
}

type Move struct {
	x         int
	y         int
	canRepeat bool
}

func NewBoard(x, y int) Board {
	var b Board

	var isWhite bool

	b.Celss = make([][]Cell, x)
	for i := 0; i < len(b.Celss); i++ {
		b.Celss[i] = make([]Cell, y)
		for j := 0; j < len(b.Celss[i]); j++ {
			if isWhite {
				b.Celss[i][j] = Cell{
					Color: White,
				}
			} else {
				b.Celss[i][j] = Cell{
					Color: Black,
				}
			}

			isWhite = !isWhite

		}

		if y%2 == 0 { // to create bord when y is odd number and normal colors
			isWhite = !isWhite
		}
	}

	return b
}

// x from left, y from top
func (b *Board) InitializePiece(x, y int, p Piece) {
	b.Celss[y][x].Piece = &p
}
