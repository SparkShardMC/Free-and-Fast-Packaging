package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func ShowUpdateUI(version string, notes string, onAction func()) {
	myApp := app.New()
	myWindow := myApp.NewWindow("Free & Fast Packager v" + version)

	title := widget.NewLabelWithStyle("FREE & FAST PACKAGING", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
	status := widget.NewLabel("Ready to package your Minecraft files.")
	
	actionBtn := widget.NewButton("START PACKAGING", func() {
		status.SetText("Processing... check console.")
		onAction()
	})

	myWindow.SetContent(container.NewVBox(
		title,
		widget.NewLabel("Version: "+version),
		widget.NewSeparator(),
		widget.NewLabel(notes),
		actionBtn,
		status,
	))

	myWindow.Resize(fyne.NewSize(400, 300))
	myWindow.ShowAndRun()
}
