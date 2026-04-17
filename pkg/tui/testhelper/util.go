package testhelper

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/jybp/render-cli/v2/pkg/tui"
)

func Stackify(m tea.Model) tea.Model {
	stack := tui.NewStack()
	stack.Push(tui.ModelWithCmd{Model: m, Cmd: ""})
	return stack
}
