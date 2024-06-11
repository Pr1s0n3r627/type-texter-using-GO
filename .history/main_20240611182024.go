package main

import (
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/go-vgo/robotgo"
)

func main() {
	a := app.New()
	w := a.NewWindow("Code Typer")

	textEdit := widget.NewMultiLineEntry()
	textEdit.SetPlaceHolder("Paste your code here")

	// Set the preferred size of the textEdit widget
	textEdit.Resize(fyne.NewSize(200, 300))

	typeButton := widget.NewButton("Type", func() {
		go simulateTyping(textEdit.Text)
	})

	// Use a VBox for the textEdit and typeButton, with a spacer to push the button to the bottom
	content := container.NewVBox(
		textEdit,
		widget.NewSeparator(),
		container.NewHBox(
			layout.NewSpacer(),
			typeButton,
		),
	)

	w.SetContent(content)
	w.Resize(fyne.NewSize(200, 400))

	w.ShowAndRun()
}

func simulateTyping(code string) {
	lines := strings.Split(code, "\n")
	for _, line := range lines {
		typeLineWithFormatting(line)
		pressEnter() // Press Enter after typing each line
	}
}

func typeLineWithFormatting(line string) {
	for _, char := range line {
		time.Sleep(50 * time.Millisecond) // Simulate typing speed
		robotgo.KeyTap(string(char))      // Simulate typing each character
	}
}

func pressEnter() {
	time.Sleep(100 * time.Millisecond) // Add slight delay before pressing Enter
	robotgo.KeyTap("enter")            // Press Enter
}
