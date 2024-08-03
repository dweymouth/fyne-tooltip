package internal

import (
	"errors"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

const (
	initialToolTipDelay             = 750 * time.Millisecond
	subsequentToolTipDelay          = 300 * time.Millisecond
	subsequentToolTipDelayValidTime = 1500 * time.Millisecond
)

var (
	toolTipLayers             = make(map[fyne.Canvas]*ToolTipLayer)
	lastToolTipShownUnixMilli int64
)

func NextToolTipDelayTime() time.Duration {
	if time.Now().UnixMilli()-lastToolTipShownUnixMilli < subsequentToolTipDelayValidTime.Milliseconds() {
		return subsequentToolTipDelay
	}
	return initialToolTipDelay
}

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
	lastToolTipShownUnixMilli = time.Now().UnixMilli()
	overlay := canvas.Overlays().Top()
	handle := &ToolTipHandle{canvas: canvas, overlay: overlay}
	tl := findToolTipLayer(handle)
	if tl == nil {
		return nil
	}

	t := NewToolTip(text)
	tl.Container.Objects = []fyne.CanvasObject{t}

	var zeroPos fyne.Position
	if pop, ok := overlay.(*widget.PopUp); ok && pop != nil {
		zeroPos = pop.Content.Position()
	} else {
		zeroPos = fyne.CurrentApp().Driver().AbsolutePositionForObject(&tl.Container)
	}

	sizeAndPositionToolTip(zeroPos, pos.Subtract(zeroPos), t, canvas)
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

const (
	maxToolTipWidth = 600
	belowMouseDist  = 16
	aboveMouseDist  = 8
)

func sizeAndPositionToolTip(zeroPos, relPos fyne.Position, t *ToolTip, canvas fyne.Canvas) {
	canvasSize := canvas.Size()
	canvasPad := theme.Padding()

	// calculate width of tooltip
	w := fyne.Min(t.NonWrappingTextWidth(), fyne.Min(canvasSize.Width-canvasPad*2, maxToolTipWidth))
	t.Resize(fyne.NewSize(w, 1)) // set up to get min height with wrapping at width w
	t.Resize(fyne.NewSize(w, t.TextMinSize().Height))

	// if would overflow the right edge of the window, move back to the left
	if rightEdge := relPos.X + zeroPos.X + w; rightEdge > canvasSize.Width-canvasPad {
		relPos.X -= rightEdge - canvasSize.Width + canvasPad
	}

	// if would overflow the bottom of the window, move above mouse
	if bottomEdge := relPos.Y + zeroPos.Y + t.Size().Height + belowMouseDist; bottomEdge > canvasSize.Height-canvasPad {
		relPos.Y -= t.Size().Height + aboveMouseDist
	} else {
		relPos.Y += belowMouseDist
	}

	t.Move(relPos)
}
