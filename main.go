package main

import (
	"github.com/charmbracelet/lipgloss"
	"fmt"
	"os"
    "strconv"

	tea "github.com/charmbracelet/bubbletea"
    "github.com/rdawson46/ttt/game"
)

var (
    even = lipgloss.NewStyle().Background(lipgloss.Color("#ff0000"))
    odd = lipgloss.NewStyle().Background(lipgloss.Color("#0000ff"))
)

type Model struct {
    game game.Game
    msg  string
}

func Initialize() tea.Model {
    return Model{
        game: game.NewGame(0, 0),
        msg: "",
    }
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type){
    case tea.KeyMsg:
        switch msg.String() {
        case "q":
            return m, tea.Quit
        case "r":
            return Initialize(), nil
        default:
            if m.game.Completed {
                break
            }

            key := msg.String()

            x, err := strconv.Atoi(key)

            if err != nil {
                return m, nil
            }

            x -= 1

            g, err := m.game.MakeHumanMove(x)

            if err != nil {
                m.msg = err.Error()
            } else {
                m.msg = ""
            }
            m.game = g

            if g.IsComputerTurn() {
                return m, computerTurn(m.game)
            }

            return m, nil
        }
    case computerPlayerMsg:
        return m, nil
    }
    return m, nil
}

func (m Model) View() string {
    s  := ""
    alternate := true

    for a, x := range m.game.Board.Grid {
        var block_text string
        if x == ' ' {
            block_text = fmt.Sprintf("%d", (a + 1))
        } else {
            block_text = string(x)
        }

        if alternate {
            block := even.Render(block_text)
            s = fmt.Sprintf("%s %s", s, block)
        } else {
            block := odd.Render(block_text)
            s = fmt.Sprintf("%s %s", s, block)
        }

        if a % 3 == 2 {
            s += "\n"
        }

        alternate = !alternate
    }

    s += fmt.Sprintf("%s", m.msg)

    return s
}

func (m Model) Init() tea.Cmd {
    if m.game.IsComputerTurn() {
        return computerTurn(m.game)
    }

    return nil
}

// Cmd for computer to take a turn
func computerTurn(game game.Game) tea.Cmd {
    return func() tea.Msg {
        x := game.MinMax()
        return computerPlayerMsg{ pos: x }
    }
}

type computerPlayerMsg struct {
    pos int
}

func main() {
    p := tea.NewProgram(Initialize())

    if _, err := p.Run(); err != nil {
        fmt.Println("broke:", err)
        os.Exit(1)
    }
}
