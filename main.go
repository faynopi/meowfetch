package main

import (
	"fmt"
	"log"
	"meowfetch/cmd"
	config "meowfetch/configs"
	"meowfetch/internals/widgets"
	"strconv"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var (
    MEOWLINES = []string {
        "                ",
        "   /| ､         ",
        "  (°､ ｡ 7       ",
        "   |､  ~ヽ      ",
        "   じしf_,)/    ",
        "                ",
    }
)

func main() {
    config.ParseConfig()
    cmd.Run()

    var line, name, MEOW, widget string
    var border lipgloss.Border
    var sep lipgloss.Style
    var w, maxsz int
    var widget_border bool
    widgetlist := strings.Split(*config.Widgets, " ")

    for i := 0; i < 6; i ++ {
        if len(widgetlist) > i {
            switch strings.ToLower(widgetlist[i]) {
                case "de":
                    name = "de"
                    widget = widgets.W_de()
                    w = len(name)
                    if *config.Uppercase_name { name = strings.ToUpper(name) }
                    if *config.Bold_name { name = lipgloss.NewStyle().Bold(true).Render(name) }
                    sep = lipgloss.NewStyle().Width(10 - w)
                    if *config.AlignSeparator { sep.Align(lipgloss.Right) }
                    formated := fmt.Sprintf("%-6s %s %s", name, sep.Render(*config.Separator), widget)
                    if sw, _ := lipgloss.Size(formated); sw > maxsz { maxsz  = sw }
                    line += formated

                case "mem":
                    name = "mem"
                    widget = widgets.W_mem()
                    w = len(name)
                    if *config.Uppercase_name { name = strings.ToUpper(name) }
                    if *config.Bold_name { name = lipgloss.NewStyle().Bold(true).Render(name) }
                    sep = lipgloss.NewStyle().Width(10 - w)
                    if *config.AlignSeparator { sep.Align(lipgloss.Right) }
                    formated := fmt.Sprintf("%-6s %s %s", name, sep.Render(*config.Separator), widget)
                    if sw, _ := lipgloss.Size(formated); sw > maxsz { maxsz  = sw }
                    line += formated

                case "cpu":
                    name = "cpu"
                    widget = widgets.W_cpu()
                    w = len(name)
                    if *config.Uppercase_name { name = strings.ToUpper(name) }
                    if *config.Bold_name { name = lipgloss.NewStyle().Bold(true).Render(name) }
                    sep = lipgloss.NewStyle().Width(10 - w)
                    if *config.AlignSeparator { sep.Align(lipgloss.Right) }
                    formated := fmt.Sprintf("%-6s %s %s", name, sep.Render(*config.Separator), widget)
                    if sw, _ := lipgloss.Size(formated); sw > maxsz { maxsz  = sw }
                    line += formated

                case "uptime":
                    name = "uptime"
                    widget = widgets.W_uptime()
                    w = len(name)
                    if *config.Uppercase_name { name = strings.ToUpper(name) }
                    if *config.Bold_name { name = lipgloss.NewStyle().Bold(true).Render(name) }
                    sep = lipgloss.NewStyle().Width(10 - w)
                    if *config.AlignSeparator { sep.Align(lipgloss.Right) }
                    formated := fmt.Sprintf("%-6s %s %s", name, sep.Render(*config.Separator), widget)
                    if sw, _ := lipgloss.Size(formated); sw > maxsz { maxsz  = sw }
                    line += formated

                case "os":
                    name = "os"
                    widget = widgets.W_os()
                    w = len(name)
                    if *config.Uppercase_name { name = strings.ToUpper(name) }
                    if *config.Bold_name { name = lipgloss.NewStyle().Bold(true).Render(name) }
                    sep = lipgloss.NewStyle().Width(10 - w)
                    if *config.AlignSeparator { sep.Align(lipgloss.Right) }
                    formated := fmt.Sprintf("%-6s %s %s", name, sep.Render(*config.Separator), widget)
                    if sw, _ := lipgloss.Size(formated); sw > maxsz { maxsz  = sw }
                    line += formated

                case "shell":
                    name = "shell"
                    widget = widgets.W_shell()
                    w = len(name)
                    if *config.Uppercase_name { name = strings.ToUpper(name) }
                    if *config.Bold_name { name = lipgloss.NewStyle().Bold(true).Render(name) }
                    sep = lipgloss.NewStyle().Width(10 - w)
                    if *config.AlignSeparator { sep.Align(lipgloss.Right) }
                    formated := fmt.Sprintf("%-6s %s %s", name, sep.Render(*config.Separator), widget)
                    if sw, _ := lipgloss.Size(formated); sw > maxsz { maxsz  = sw }
                    line += formated

                case "colors":
                    widget = widgets.W_colors()
                    line += "\n" + lipgloss.NewStyle().Width(maxsz).Align(lipgloss.Center).Render(widget)
                    i ++ // cuz of the new line

                default:

                    log.Fatal("Unknown widget " + widgetlist[i] + "!")
            }
            if i < 5 {
                line += "\n"
            }
        }
    }

    switch *config.BorderStyle {
        case "round":
            border = lipgloss.RoundedBorder()
        case "single":
            border = lipgloss.NormalBorder()
        case "double":
            border = lipgloss.DoubleBorder()
        case "none":
            border = lipgloss.Border{}
        default:
            log.Fatal("Unknown border_style value!")
    }


    switch *config.BorderMode {
        case "widgets", "widget":
            widget_border = true
        case "all":
            widget_border = false
        default:
            log.Fatal("Unknown border_mode value!")

    }

    switch *config.Mode {
        case "pride":
            for i, M := range MEOWLINES {
                switch i {
                    case 0:
                        MEOW += lipgloss.NewStyle().
                            Foreground(lipgloss.Color("0")).
                            Background(lipgloss.Color("#f38ba8")).
                            Render(M) + "\n"
                    case 1:
                        MEOW += lipgloss.NewStyle().
                            Foreground(lipgloss.Color("0")).
                            Background(lipgloss.Color("#fab387")).
                            Render(M) + "\n"
                    case 2:
                        MEOW += lipgloss.NewStyle().
                            Foreground(lipgloss.Color("0")).
                            Background(lipgloss.Color("#f9e2af")).
                            Render(M) + "\n"
                    case 3:
                        MEOW += lipgloss.NewStyle().
                            Foreground(lipgloss.Color("0")).
                            Background(lipgloss.Color("#a6e3a1	")).
                            Render(M) + "\n"
                    case 4:
                        MEOW += lipgloss.NewStyle().
                            Foreground(lipgloss.Color("0")).
                            Background(lipgloss.Color("#74c7ec")).
                            Render(M) + "\n"
                    case 5:
                        MEOW += lipgloss.NewStyle().
                            Foreground(lipgloss.Color("0")).
                            Background(lipgloss.Color("#cba6f7")).
                            Render(M)
                }
            }
        case "color":
            if *config.Color < 255 {
                for _, M := range MEOWLINES {
                    MEOW += lipgloss.NewStyle().
                        Foreground(lipgloss.Color(strconv.Itoa(*config.Color))).
                        Render(M) + "\n"
                }
            } else {
                log.Fatal("Unknown color value. color must be an integer between 0 and 255.")
            }
        default:
            log.Fatal("Unknown mode " + *config.Mode + "!")
    }

    if *config.Bold_art {
        MEOW = lipgloss.NewStyle().Bold(true).Render(MEOW)
    }

    if widget_border {
        fmt.Println(lipgloss.JoinHorizontal(
            lipgloss.Top,
            lipgloss.NewStyle().
                Margin(1).
                MarginTop(2).
                Render(MEOW),
            lipgloss.NewStyle().
                Height(6).
                MarginTop(1).
                Border(border).
                PaddingLeft(1).
                PaddingRight(3).
                Render(line),
        ))
    } else {
        fmt.Println(lipgloss.NewStyle().
            Height(4).
            Border(border).
            PaddingRight(5).
            Render(
                lipgloss.JoinHorizontal(
                    lipgloss.Center,
                    lipgloss.NewStyle().
                        Margin(1).
                        Render(MEOW),
                    line,
            )))
    }
}

