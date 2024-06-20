package game

type player interface {
    makeMove(board *Board, pos int) bool
}

type computerPlayer struct {
    symbol rune
}

func (p computerPlayer) makeMove(board *Board, pos int) bool {
    return board.makeMove(pos, p.symbol)
}

type humanPlayer struct {
    symbol rune
}

func (p humanPlayer) makeMove(board *Board, pos int) bool {
    return board.makeMove(pos, p.symbol)
}
