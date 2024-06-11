package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/go-vgo/robotgo"
)

func main() {
	a := app.NewWithID("text-typer")
	applyDarkTheme(a)

	w := a.NewWindow("Text Typer")

	textEdit := widget.NewMultiLineEntry()
	textEdit.SetPlaceHolder("Enter text to type here")

	progressBar := widget.NewProgressBarInfinite()

	typeButton := widget.NewButton("Type", func() {
		go startTimer(textEdit.Text, progressBar)
	})

	content := container.NewVBox(textEdit, typeButton, progressBar)
	w.SetContent(content)

	w.Resize(fyne.NewSize(400, 166))
	w.ShowAndRun()
}

func startTimer(text string, progressBar *widget.ProgressBarInfinite) {
	progressBar.Start()
	typeText(text)
	progressBar.Stop()
}

func typeText(text string) {
	if text != "" {
		robotgo.TypeStr(text)
	}
}

func applyDarkTheme(a fyne.App) {
	a.Settings().SetTheme(theme.DarkTheme())
}
