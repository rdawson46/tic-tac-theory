package game

type Board struct {
    Grid [3][3]rune
}

func newBoard() *Board {
    x := [3][3]rune {
        {' ', ' ', ' '},
        {' ', ' ', ' '},
        {' ', ' ', ' '},
    }

    return &Board{Grid: x}
}

func (b *Board) makeMove(pos int, team rune) bool {
    x := pos / 3
    y := pos % 3

    if b.Grid[x][y] != ' ' {
        return false
    }

    b.Grid[x][y] = team
    return true
}
