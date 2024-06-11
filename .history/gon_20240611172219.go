package main

import (
	"strings"
	"time"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/go-vgo/robotgo"
)

func gon() {
	a := app.New()
	w := a.NewWindow("Code Typer")

	textEdit := widget.NewMultiLineEntry()
	textEdit.SetPlaceHolder("Paste your code here")

	typeButton := widget.NewButton("Type", func() {
		go simulateTyping(textEdit.Text)
	})

	content := container.NewVBox(textEdit, typeButton)
	w.SetContent(content)

	w.ShowAndRun()
}

func simulateTyping(code string) {
	lines := strings.Split(code, "\n")
	for _, line := range lines {
		typeLineWithFormatting(line)
	}
}

func typeLineWithFormatting(line string) {
	time.Sleep(100 * time.Millisecond) // Simulate typing speed
	robotgo.TypeString(line)
	robotgo.KeyTap("enter") // Press Enter after typing each line
}
