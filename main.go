package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

var (
	FIELD_LENGTH = 130
	FIELD_HEIGHT = 100
	SPEED        = 8
)

func main() {
	// TODO:
	// 1. Add a label with generation #
	// 2. Change cell colors
	// 3*. Add start button
	gameApp := app.New()
	gameWindow := gameApp.NewWindow("Game Life")

	runGame(gameWindow)

	gameWindow.Resize(fyne.NewSize(1000, 700))
	gameWindow.ShowAndRun()
}
