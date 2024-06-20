package game

type Board struct {
    Grid [9]rune
}

func newBoard() *Board {
    x := [9]rune {
        ' ', ' ', ' ',
        ' ', ' ', ' ',
        ' ', ' ', ' ',
    }

    return &Board{Grid: x}
}

func (b *Board) makeMove(pos int, team rune) bool {
    if b.Grid[pos] != ' ' {
        return false
    }

    b.Grid[pos] = team
    return true
}
