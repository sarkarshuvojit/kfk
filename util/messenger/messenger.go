package messenger

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func printWithColor(text string, color lipgloss.Color) {
    var style = lipgloss.NewStyle().
        Bold(true).
        Foreground(color)

    fmt.Println(style.Render(text))
}

func Success(message string, args ...string) {
    var fmtMsg string
    if len(args) == 0 {
        fmtMsg = message
    } else {
        fmtMsg = fmt.Sprintf(message, args)
    }
    printWithColor("✓ "+fmtMsg, "#27ae60")
}

func Error(message string, args ...string) {
    var fmtMsg string
    if len(args) == 0 {
        fmtMsg = message
    } else {
        fmtMsg = fmt.Sprintf(message, args)
    }
    printWithColor("x "+fmtMsg, "#c0392b")
}

func Warning(message string, args ...string) {
    var fmtMsg string
    if len(args) == 0 {
        fmtMsg = message
    } else {
        fmtMsg = fmt.Sprintf(message, args)
    }
    printWithColor("! "+fmtMsg, "#f39c12")
}

