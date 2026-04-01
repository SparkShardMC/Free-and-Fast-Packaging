package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"log"
)

// ShowUpdateUI now accepts a custom application handle to set the icon
func ShowUpdateUI(newVersion string, changelog string, downloadFn func()) {
	// 1. Initialize the Fyne App
	myApp := app.New()

	// 2. Load the Logo from assets/ (Crucial Step)
	// This tells Go, "Go find that PNG and read its data"
	iconResource, err := fyne.LoadResourceFromPath("assets/Free&Fast_logo.png")
	if err != nil {
		// If the file is missing, we log it, but the app will still run (no crash)
		log.Println("⚠️ Could not load app logo from assets/Free&Fast_logo.png:", err)
	} else {
		// If found, set the icon for the WHOLE application (App Button)
		myApp.SetIcon(iconResource)
	}

	// 3. Create the window as before
	window := myApp.NewWindow("Update Available - Free-and-Fast")
	window.Resize(fyne.NewSize(450, 350))

	// Create the logo image for INSIDE the window, too
	var logoImage *canvas.Image
	if iconResource != nil {
		logoImage = canvas.NewImageFromResource(iconResource)
		logoImage.FillMode = canvas.ImageFillContain
		logoImage.SetMinSize(fyne.NewSize(64, 64)) // Standard Icon Size
	}

	// Rest of the UI elements
	label := widget.NewLabel("🚀 New Version Found: " + newVersion)
	label.Alignment = fyne.TextAlignCenter // Centered look

	notes := widget.NewRichTextFromMarkdown(changelog)
	progressBar := widget.NewProgressBar()
	progressBar.Hide()

	// ... (Rest of the button code is the same as image_12)
}
