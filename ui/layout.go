package ui

import "fyne.io/fyne/v2/container"

// Setup all of our UI elements.
func Setup(app *AppInit) {
	SetupMenus(app)
	swatchesContainer := BuildSwatches(app)
	colorPicker := SetupColorPicker(app)

	appLayout := container.NewBorder(nil, swatchesContainer, nil, colorPicker, app.PixlCanvas)

	app.PixlWindow.SetContent(appLayout)
}
