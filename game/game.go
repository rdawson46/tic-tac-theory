package game

import (
	"errors"
	"fmt"
	"os"
    // "github.com/rdawson46/ttt/utils"
)

type Game struct {
    player1   player
    player2   player
    Board     *Board
    Completed bool
    turn      bool
}

const (
    Human = iota
    Computer
)

func playerType(t int, team rune) player {
    switch t {
    case Human:
        return humanPlayer{team}
    case Computer:
        return computerPlayer{team}
    }

    fmt.Println("Invalid Player type")
    os.Exit(1)
    return nil
}


func NewGame(type1, type2 int) Game {
    player1 := playerType(type1, 'X')
    player2 := playerType(type2, 'Y')

    b := newBoard()

    return Game {
        player1: player1,
        player2: player2,
        Board: b,
        turn: true,
    }
}

// check for a winner
func (g Game) IsWinner() (bool, rune) {
    return false, ' '
}

func (g Game) IsComputerTurn() bool {
    if g.turn {
        switch g.player1.(type) {
        case computerPlayer:
            return true
        default:
            return false
        }
    } else {
        switch g.player2.(type) {
        case computerPlayer:
            return true
        default:
            return false
        }
    }
}

func (g Game) MakeHumanMove(pos int) (Game, error) {
    // check if is a human player making the move
    if g.turn {
        switch g.player1.(type) {
        case humanPlayer:
            return g.MakeMove(pos)
        default:
            return g, errors.New("Not player turn")
        }
    } else {
        switch g.player2.(type) {
        case humanPlayer:
            return g.MakeMove(pos)
        default:
            return g, errors.New("Not player turn")
        }
    }
}

func (g Game) MinMax() int {
    return 0
}

func (g Game) MakeComputerMove(pos int) (Game, error) {
    // check if is a computer player making the move
    if g.turn {
        switch g.player1.(type) {
        case computerPlayer:
            return g.MakeMove(pos)
        default:
            return g, errors.New("Not player turn")
        }
    } else {
        switch g.player2.(type) {
        case computerPlayer:
            return g.MakeMove(pos)
        default:
            return g, errors.New("Not player turn")
        }
    }
}


func (g Game) MakeMove(pos int) (Game, error) {
    if g.Completed || !g.SpotsRemaining() {
        return g, errors.New("Game Over")
    }

    if pos > 9 {
        return g, errors.New("Invalid location")
    }

    var res bool

    if g.turn {
        res = g.player1.makeMove(g.Board, pos)
    } else {
        res = g.player2.makeMove(g.Board, pos)
    }

    if !res {
        return g, errors.New("Ivalid location")
    }

    // check for winner
    var winner rune
    if g.Completed, winner = g.IsWinner(); g.Completed {
        return g, errors.New(fmt.Sprintf("Winner is %c\n", winner))
    }

    g.turn = !g.turn

    // switch current player
    return g, nil
}

func (g Game) SpotsRemaining() bool {
    for _, x := range g.Board.Grid {
        if x == ' ' {
            return true
        }
    }

    return false
}
