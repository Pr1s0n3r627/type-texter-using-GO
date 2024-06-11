package main

import (
	"time"

	"github.com/fyne-io/fyne/v2/app"
	"github.com/fyne-io/fyne/v2/container"
	"github.com/fyne-io/fyne/v2/theme"
	"github.com/fyne-io/fyne/v2/widget"
	"github.com/go-vgo/robotgo"
)

func main() {
	a := app.NewWithID("com.example.texttyper")
	applyDarkTheme(a)

	w := a.NewWindow("Text Typer")

	textEdit := widget.NewMultiLineEntry()
	progressBar := widget.NewProgressBar()
	progressBar.Max = 100

	typeButton := widget.NewButton("Type Text", func() {
		startTimer(textEdit.Text, progressBar)
	})

	content := container.NewVBox(textEdit, typeButton, progressBar)
	w.SetContent(content)

	w.Resize(fyne.NewSize(400, 300))
	w.ShowAndRun()
}

func startTimer(text string, progressBar *widget.ProgressBar) {
	countdown := 100
	go func() {
		for countdown >= 0 {
			time.Sleep(50 * time.Millisecond)
			countdown--
			progressBar.SetValue(float64(countdown))
		}
		typeText(text)
	}()
}

func typeText(text string) {
	if text != "" {
		robotgo.TypeStr(text)
	}
}

func applyDarkTheme(a fyne.App) {
	a.Settings().SetTheme(theme.DarkTheme())
}
