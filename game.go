package main

import (
	"image"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type Game struct {
	raster          *canvas.Raster
	field           *Field
	container       *fyne.Container
	generationLabel *widget.Label
}

func (game *Game) draw(w, h int) image.Image {
	rasterImg := image.NewRGBA(image.Rect(0, 0, FIELD_LENGTH, FIELD_HEIGHT))

	for i := 0; i < FIELD_LENGTH; i++ {
		for j := 0; j < FIELD_HEIGHT; j++ {
			if game.field.cells[i][j] {
				rasterImg.Set(i, j, CELL_COLOR)
			}
		}
	}

	return rasterImg
}

func (game *Game) runGameLoop() {
	for range time.NewTicker(time.Second / time.Duration(SPEED)).C {
		game.raster.Refresh()
		game.generationLabel.SetText(strconv.Itoa(game.field.generation))

		game.field.calculateNextState()
	}
}

func createGame() Game {
	game := Game{}
	game.raster = canvas.NewRaster(game.draw)
	game.raster.ScaleMode = canvas.ImageScalePixels
	game.raster.SetMinSize(fyne.NewSize(float32(SCREEN_WIDTH_PX), float32(SCREEN_HEIGHT_PX)))

	game.field = createNewField()
	game.field.setRandomFieldState()

	game.generationLabel = widget.NewLabel("0")
	labelBox := container.New(
		layout.NewHBoxLayout(),
		layout.NewSpacer(),
		widget.NewLabel("Generation: "),
		game.generationLabel,
	)
	game.container = container.New(layout.NewVBoxLayout(), game.raster, labelBox)

	return game
}
