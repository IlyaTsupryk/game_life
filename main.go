package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

var (
	FIELD_LENGTH     = 130
	FIELD_HEIGHT     = 100
	SPEED            = 8
	SCREEN_WIDTH_PX  = 1000
	SCREEN_HEIGHT_PX = 700
	CELL_COLOR       = color.RGBA{R: 55, G: 102, B: 68, A: 255}
)

func main() {
	gameApp := app.New()
	gameWindow := gameApp.NewWindow("Game Life")

	game := createGame()
	gameWindow.SetContent(game.container)
	gameWindow.Resize(fyne.NewSize(float32(SCREEN_WIDTH_PX), float32(SCREEN_HEIGHT_PX)))
	go game.runGameLoop()

	gameWindow.ShowAndRun()
}
