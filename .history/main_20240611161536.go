package main

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/go-vgo/robotgo"
)

func main() {
	a := app.NewWithID("com.example.texttyper")
	applyDarkTheme(a)

	w := a.NewWindow("Text Typer")

	textEdit := widget.NewMultiLineEntry()
	progressBar := widget.NewProgressBar()

	typeButton
