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

	var progressBar *widget.ProgressBar // Declare progressBar variable
	content := container.NewVBox()
	typeButton := widget.NewButton("Type", func() {
		progressBar = widget.NewProgressBar()
		progressBar.Max = 5 // Set the maximum value of the progress bar to 5 seconds

		go simulateTyping(textEdit.Text, progressBar)

		// Add progress bar to the content
		content.Add(progressBar)
	})

	// Create a VBox to hold the text field, button, and progress bar
	content := container.NewVBox(
		textEdit,
		typeButton,
	)

	w.SetContent(content)
	w.Resize(fyne.NewSize(450, 200)) // Adjusted window size

	w.ShowAndRun()
}

func simulateTyping(code string, progressBar *widget.ProgressBar) {
	lines := strings.Split(code, "\n")

	// Update progress bar while waiting
	for i := 0; i <= 100; i++ {
		progressBar.SetValue(float64(i) / 100)
		time.Sleep(50 * time.Millisecond)
	}

	// Simulate typing after 5 seconds
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
