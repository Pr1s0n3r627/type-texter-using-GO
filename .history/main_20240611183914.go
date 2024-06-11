package main

import (
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/go-vgo/robotgo"
)

func main() {
	a := app.New()
	w := a.NewWindow("Code Typer")

	textEdit := widget.NewMultiLineEntry()
	textEdit.SetPlaceHolder("Paste your code here")

	typeButton := widget.NewButton("Type", func() {
		go simulateTyping(textEdit.Text)
	})

	// Arrange the widgets using VBox and VBox containers
	content := container.NewVBox(
		textEdit,
		typeButton,
	)

	w.SetContent(content)
	w.Resize(fyne.NewSize(400, 13))

	w.ShowAndRun()
}

func simulateTyping(code string) {
	lines := strings.Split(code, "\n")
	for _, line := range lines {
		typeLineWithFormatting(line)
		//pressEnter() // Press Enter after typing each line
	}
}

func typeLineWithFormatting(line string) {
	for _, char := range line {
		time.Sleep(time.Microsecond) // Simulate typing speed
		robotgo.KeyTap(string(char)) // Simulate typing each character
	}
}
