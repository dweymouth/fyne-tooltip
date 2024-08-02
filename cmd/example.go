package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	fynetooltip "github.com/dweymouth/fyne-tooltip"
	ttwidget "github.com/dweymouth/fyne-tooltip/widget"
)

func main() {
	app := app.New()
	win := app.NewWindow("fyne-tooltip demo")

	var btnA *ttwidget.Button
	btnA = ttwidget.NewButton("Show PopUp", func() {
		driver := fyne.CurrentApp().Driver()
		canvas := driver.CanvasForObject(btnA)
		showPopUp(canvas, driver.AbsolutePositionForObject(btnA).Add(btnA.Size()))
	})
	btnA.SetToolTip("Show a tooltip-enabled PopUp")

	btnB := ttwidget.NewButton("Show Modal PopUp", nil)
	btnB.SetToolTip("Show a tooltip-enabled modal PopUp")

	content := container.NewCenter(
		container.NewHBox(btnA, btnB),
	)

	win.SetContent(fynetooltip.AddWindowToolTipLayer(content, win.Canvas()))
	win.Resize(fyne.NewSize(400, 300))
	win.ShowAndRun()
}

var reusablePopUp *widget.PopUp

func showPopUp(canvas fyne.Canvas, pos fyne.Position) {
	if reusablePopUp == nil {
		btn := ttwidget.NewButton("hello", nil)
		btn.SetToolTip("world yeah")
		reusablePopUp = widget.NewPopUp(container.NewPadded(btn), canvas)

		// a pop up that will be reused only needs a call to
		// AddPopUpToolTipLayer during setup
		fynetooltip.AddPopUpToolTipLayer(reusablePopUp)
	}

	reusablePopUp.ShowAtPosition(pos)
}
