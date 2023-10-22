package main

import (
	"fmt"
	"image"
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

type Game struct {
	raster *canvas.Raster
	field  *Field
}

func (game *Game) draw(w, h int) image.Image {
	rasterImg := image.NewRGBA(image.Rect(0, 0, FIELD_LENGTH, FIELD_HEIGHT))

	for i := 0; i < FIELD_LENGTH; i++ {
		for j := 0; j < FIELD_HEIGHT; j++ {
			if game.field.cells[i][j] {
				rasterImg.Set(i, j, color.Black)
			}
		}
	}

	return rasterImg
}

func (game *Game) runGameLoop() {
	for range time.NewTicker(time.Second / time.Duration(SPEED)).C {
		game.raster.Refresh()
		fmt.Println("Generation: ", game.field.generation)

		game.field.calculateNextState()
	}
}

func runGame(gameWindow fyne.Window) {
	game := Game{}
	game.raster = canvas.NewRaster(game.draw)
	game.raster.ScaleMode = canvas.ImageScalePixels

	game.field = createNewField()
	game.field.setRandomFieldState()

	gameWindow.SetContent(game.raster)

	go game.runGameLoop()
}
