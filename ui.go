package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"time"
)

func ShowUpdateUI(newVersion string, changelog string, downloadFn func()) {
	myApp := app.New()
	window := myApp.NewWindow("Update Available - Free-and-Fast")
	window.Resize(fyne.NewSize(400, 300))

	label := widget.NewLabel("🚀 New Version Found: " + newVersion)
	notes := widget.NewRichTextFromMarkdown(changelog)
	
	// The "Filling Jar" Progress Bar
	progressBar := widget.NewProgressBar()
	progressBar.Hide()

	btn := widget.NewButton("Update Now", func() {
		progressBar.Show()
		// Simulate the "Filling Jar" filling up
		go func() {
			for i := 0.0; i <= 1.0; i += 0.01 {
				time.Sleep(time.Millisecond * 50)
				progressBar.SetValue(i)
			}
			downloadFn() // Triggers the actual GitHub download
		})
	})

	window.SetContent(container.NewVBox(
		label,
		container.NewScroll(notes),
		progressBar,
		btn,
	))

	window.ShowAndRun()
}
