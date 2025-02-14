// filepath: /c:/Users/magnu/Documents/GitHub/uit-misc/div-scripts/div-projects/int-grade-conversion.go
package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	// Create a new application
	myApp := app.New()

	// Create a new window
	myWindow := myApp.NewWindow("Hello Fyne")

	// Create a label widget
	label := widget.NewLabel("Hello, World!")

	// Set the content of the window
	myWindow.SetContent(container.NewVBox(
		label,
		widget.NewButton("Quit", func() {
			myApp.Quit()
		}),
	))

	// Show and run the application
	myWindow.ShowAndRun()
}
