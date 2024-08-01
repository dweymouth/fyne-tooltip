package internal

import (
	"errors"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

var toolTipLayers = make(map[fyne.Canvas]*ToolTipLayer)

type ToolTipLayer struct {
	Container fyne.Container
	overlays  map[fyne.CanvasObject]*ToolTipLayer
}

func NewToolTipLayer(canvas fyne.Canvas) *ToolTipLayer {
	t := &ToolTipLayer{}
	toolTipLayers[canvas] = t
	return t
}

func ToolTipLayerForCanvas(canvas fyne.Canvas) *ToolTipLayer {
	return toolTipLayers[canvas]
}

func DestroyToolTipLayerForCanvas(canvas fyne.Canvas) {
	delete(toolTipLayers, canvas)
}

func NewPopUpToolTipLayer(popUp *widget.PopUp) *ToolTipLayer {
	ct := ToolTipLayerForCanvas(popUp.Canvas)
	if ct == nil {
		fyne.LogError("", errors.New("no pop up layer created for parent canvas"))
		return nil
	}
	t := &ToolTipLayer{}
	if ct.overlays == nil {
		ct.overlays = make(map[fyne.CanvasObject]*ToolTipLayer)
	}
	ct.overlays[popUp] = t
	return t
}

func DestroyToolTipLayerForPopup(popUp *widget.PopUp) {
	ct := ToolTipLayerForCanvas(popUp.Canvas)
	if ct != nil {
		delete(ct.overlays, popUp)
	}
}

func ShowToolTipAtMousePosition(canvas fyne.Canvas, pos fyne.Position, text string) {
	tl := ToolTipLayerForCanvas(canvas)
	if tl == nil {
		fyne.LogError("", errors.New("no pop up layer created for parent canvas"))
		return
	}
	overlay := canvas.Overlays().Top()
	if overlay != nil {
		tl = tl.overlays[overlay]
		if tl == nil {
			fyne.LogError("", errors.New("no pop up layer created for overlay"))
			return
		}
	}

	t := NewToolTip(text)
	tl.Container.Objects = []fyne.CanvasObject{t}
	tl.Container.Refresh()
}
