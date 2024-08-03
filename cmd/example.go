package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
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

	btnB := ttwidget.NewButton("Show Modal PopUp", func() {
		showModalPopUp(fyne.CurrentApp().Driver().CanvasForObject(btnA))
	})
	btnB.SetToolTip("Show a tooltip-enabled modal PopUp. This tool tip text is so very very very long that it must wrap to fit on the screen.")

	lbl := ttwidget.NewLabel("a tooltip-enabled label near bottom")
	lbl.SetToolTip("Hello, world! Tooltips are great!")
	lbl.Alignment = fyne.TextAlignCenter

	content := container.NewStack(
		container.NewCenter(
			container.NewHBox(btnA, btnB),
		),
		container.NewVBox(
			layout.NewSpacer(),
			lbl,
		),
	)

	win.SetMainMenu(fyne.NewMainMenu(fyne.NewMenu("File")))
	win.SetContent(fynetooltip.AddWindowToolTipLayer(content, win.Canvas()))
	win.Resize(fyne.NewSize(400, 300))

	win.ShowAndRun()
}

var reusablePopUp *widget.PopUp

func showPopUp(canvas fyne.Canvas, pos fyne.Position) {
	hide := func() {
		reusablePopUp.Hide()
	}
	if reusablePopUp == nil {
		btnA := ttwidget.NewButton("hello", hide)
		btnA.SetToolTip("world")
		btnB := ttwidget.NewButton("world", hide)
		btnB.SetToolTip("hello - this is also a bit longer text")
		title := widget.NewLabel("My popup")
		title.Alignment = fyne.TextAlignCenter
		content := container.NewVBox(
			title,
			container.NewHBox(btnA, btnB),
		)
		reusablePopUp = widget.NewPopUp(container.NewPadded(content), canvas)

		// a pop up that will be reused only needs a call to
		// AddPopUpToolTipLayer during setup
		fynetooltip.AddPopUpToolTipLayer(reusablePopUp)
	}

	reusablePopUp.ShowAtPosition(pos)
}

func showModalPopUp(canvas fyne.Canvas) {
	var pop *widget.PopUp
	hide := func() {
		pop.Hide()
		// release memory resources when the pop up will no longer be shown
		fynetooltip.DestroyPopUpToolTipLayer(pop)
	}
	btnA := ttwidget.NewButton("hello", hide)
	btnA.SetToolTip("world")
	btnB := ttwidget.NewButton("world", hide)
	btnB.SetToolTip("hello - this is also a bit longer text")
	title := widget.NewLabel("My popup")
	title.Alignment = fyne.TextAlignCenter
	content := container.NewVBox(
		title,
		container.NewHBox(btnA, btnB),
	)
	pop = widget.NewModalPopUp(content, canvas)
	fynetooltip.AddPopUpToolTipLayer(pop)
	pop.Show()
}
