package fynetooltip

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/dweymouth/fyne-tooltip/internal"
)

func AddWindowToolTipLayer(windowContent fyne.CanvasObject, canvas fyne.Canvas) fyne.CanvasObject {
	return container.NewStack(windowContent, &internal.NewToolTipLayer(canvas).Container)
}

func DestroyWindowToolTipLayer(canvas fyne.Canvas) {
	internal.DestroyToolTipLayerForCanvas(canvas)
}

func AddPopUpToolTipLayer(p *widget.PopUp) {
	l := internal.NewPopUpToolTipLayer(p)
	p.Content = container.NewStack(p.Content, &l.Container)
}

func DestroyPopUpToolTipLayer(p *widget.PopUp) {
	internal.DestroyToolTipLayerForPopup(p)
}
