package widgets

import (
	"log"
	config "meowfetch/configs"
	"strconv"

	"github.com/charmbracelet/lipgloss"
)

func W_colors() string {
    var ret, char string

    switch *config.Colorblock_mode {
        case "circle":
            char = ""
        case "diamond":
            char = ""
        case "square":
            char = "██"
        default:
            log.Fatal("Unknown blockmode " + *config.Colorblock_mode + "!")
    }

    for i := 1; i < 8; i++ {
        ret += lipgloss.NewStyle().Foreground(lipgloss.Color(strconv.Itoa(i))).Render(char)
        if i < 7 { ret += " "}
    }

    return ret
}
