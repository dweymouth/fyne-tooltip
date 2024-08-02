package internal

import (
	"errors"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

var toolTipLayers = make(map[fyne.Canvas]*ToolTipLayer)

type ToolTipHandle struct {
	canvas  fyne.Canvas
	overlay fyne.CanvasObject
}

type ToolTipLayer struct {
	Container fyne.Container
	overlays  map[fyne.CanvasObject]*ToolTipLayer
}

func NewToolTipLayer(canvas fyne.Canvas) *ToolTipLayer {
	t := &ToolTipLayer{}
	toolTipLayers[canvas] = t
	return t
}

func DestroyToolTipLayerForCanvas(canvas fyne.Canvas) {
	delete(toolTipLayers, canvas)
}

func NewPopUpToolTipLayer(popUp *widget.PopUp) *ToolTipLayer {
	ct := toolTipLayers[popUp.Canvas]
	if ct == nil {
		fyne.LogError("", errors.New("no tool tip layer created for parent canvas"))
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
	ct := toolTipLayers[popUp.Canvas]
	if ct != nil {
		delete(ct.overlays, popUp)
	}
}

func ShowToolTipAtMousePosition(canvas fyne.Canvas, pos fyne.Position, text string) *ToolTipHandle {
	overlay := canvas.Overlays().Top()
	handle := &ToolTipHandle{canvas: canvas, overlay: overlay}
	tl := findToolTipLayer(handle)
	if tl == nil {
		return nil
	}

	t := NewToolTip(text)
	tl.Container.Objects = []fyne.CanvasObject{t}
	t.Resize(t.TextMinSize())
	if pop, ok := overlay.(*widget.PopUp); ok && pop != nil {
		pos = pos.Subtract(pop.Content.Position())
	}

	t.Move(pos)
	tl.Container.Refresh()
	return handle
}

func HideToolTip(handle *ToolTipHandle) {
	if handle == nil {
		return
	}
	tl := findToolTipLayer(handle)
	if tl == nil {
		return
	}
	tl.Container.Objects = nil
	tl.Container.Refresh()
}

func findToolTipLayer(handle *ToolTipHandle) *ToolTipLayer {
	tl := toolTipLayers[handle.canvas]
	if tl == nil {
		fyne.LogError("", errors.New("no tool tip layer created for window canvas"))
		return nil
	}
	if handle.overlay != nil {
		tl = tl.overlays[handle.overlay]
		if tl == nil {
			fyne.LogError("", errors.New("no tool tip layer created for current overlay"))
			return nil
		}
	}
	return tl
}
