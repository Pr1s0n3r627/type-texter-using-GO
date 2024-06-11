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
	w := a.NewWindow("Code Typer - Type your code with a click of a button")

	textEdit := widget.NewMultiLineEntry()
	textEdit.SetPlaceHolder("Paste your code here\nmake sure it is correct and complete")

	typeButton := widget.NewButton("Type", func() {
		go simulateTyping(textEdit.Text)
	})

	progressBar := widget.NewProgressBar()
	progressBar.Max = 5 // Set the maximum value of the progress bar to 5 seconds

	// Arrange the widgets using VBox and VBox containers
	content := container.NewVBox(
		textEdit,
		typeButton,
		progressBar,
	)

	w.SetContent(content)
	w.Resize(fyne.NewSize(450, 13))

	w.ShowAndRun()
}

func simulateTyping(code string) {
	lines := strings.Split(code, "\n")
	progressBar.start()
	time.Sleep(5 * time.Second) // Give user time to switch to the desired window
	for _, line := range lines {
		typeLineWithFormatting(line)
	}
}

func typeLineWithFormatting(line string) {
	for _, char := range line {
		time.Sleep(time.Microsecond) // Simulate typing speed
		robotgo.KeyTap(string(char)) // Simulate typing each character
	}
}
