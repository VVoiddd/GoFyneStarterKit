
package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/menu"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Go GUI Example")

	mainMenu := menu.NewMainMenu(
		menu.NewMenu("File",
			menu.NewMenuItem("Quit", func() { myApp.Quit() }),
		),
	)

	myWindow.SetMainMenu(mainMenu)

	content := container.NewVBox(
		widget.NewLabel("This is a label."),
		widget.NewButton("Click me!", func() {
			dialog.ShowInformation("Clicked", "You clicked the button.", myWindow)
		}),
		widget.NewCheck("Check me!", func(value bool) {
			if value {
				dialog.ShowInformation("Checked", "You checked the checkbox.", myWindow)
			}
		}),
		widget.NewEntry(),
		widget.NewSlider(0, 100),
	)

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
