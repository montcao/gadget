package gadget 

import (
	"github.com/charmbracelet/lipgloss"
    "github.com/charmbracelet/lipgloss/table"
)
var style = lipgloss.NewStyle().
    Foreground(lipgloss.Color("#2f824dff")).
    Background(lipgloss.Color("rgba(56, 128, 59, 1)"))

var loading_style = lipgloss.NewStyle().Bold(true)


var (
    gray      = lipgloss.Color("245")
    lightGray = lipgloss.Color("241")
	green     = lipgloss.Color("#2f824dff")

    headerStyle  = lipgloss.NewStyle().Foreground(green).Bold(true).Align(lipgloss.Center)
    cellStyle    = lipgloss.NewStyle().Padding(0, 1).Width(11)
    oddRowStyle  = cellStyle.Foreground(gray)
    evenRowStyle = cellStyle.Foreground(lightGray)
)
func CreateBanner() string {
	banner := style.Render(`
    ▄▄▄▄      ▄▄     ▄▄▄▄▄        ▄▄▄▄   ▄▄▄▄▄▄▄▄  ▄▄▄▄▄▄▄▄ 
  ██▀▀▀▀█    ████    ██▀▀▀██    ██▀▀▀▀█  ██▀▀▀▀▀▀  ▀▀▀██▀▀▀ 
 ██          ████    ██    ██  ██        ██           ██    
 ██  ▄▄▄▄   ██  ██   ██    ██  ██  ▄▄▄▄  ███████      ██    
 ██  ▀▀██   ██████   ██    ██  ██  ▀▀██  ██           ██    
  ██▄▄▄██  ▄██  ██▄  ██▄▄▄██    ██▄▄▄██  ██▄▄▄▄▄▄     ██    
    ▀▀▀▀   ▀▀    ▀▀  ▀▀▀▀▀        ▀▀▀▀   ▀▀▀▀▀▀▀▀     ▀▀

	Provided by Montcao
`)
	return banner
}

func CreateLoadingBar() string { 
	loading_bar := loading_style.Render("Inspecting your image...\n")
	return loading_bar
}

func CreateTable() *table.Table {
	return table.New().
    Border(lipgloss.NormalBorder()).
    BorderStyle(lipgloss.NewStyle().Foreground(green)).
    StyleFunc(func(row, col int) lipgloss.Style {
        switch {
		case col == 0:
			return cellStyle.Width(40)
        case row == table.HeaderRow:
            return headerStyle
        case row%2 == 0:
            return evenRowStyle
        default:
            return oddRowStyle
        }
    })
}