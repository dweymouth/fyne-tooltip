package widget

import (
	"context"
	"sync"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
	"github.com/dweymouth/fyne-tooltip/internal"
)

type toolTipContext struct {
	lock                 sync.Mutex
	toolTipHandle        *internal.ToolTipHandle
	absoluteMousePos     fyne.Position
	pendingToolTipCtx    context.Context
	pendingToolTipCancel context.CancelFunc
}

// ToolTipWidget is a base struct for building new tool tip supporting widgets.
// Inherit from this struct instead of from `widget.BaseWidget` to automatically
// add tool tip support to your widget.
type ToolTipWidget struct {
	widget.BaseWidget
	toolTipContext

	toolTip string
}

func (t *ToolTipWidget) SetToolTip(toolTip string) {
	t.toolTip = toolTip
}

func (t *ToolTipWidget) ToolTip() string {
	return t.toolTip
}

func (t *ToolTipWidget) MouseIn(e *desktop.MouseEvent) {
	t.lock.Lock()
	defer t.lock.Unlock()
	t.absoluteMousePos = e.AbsolutePosition
	t.setPendingToolTip(t, t.toolTip)
}

func (t *ToolTipWidget) MouseOut() {
	t.lock.Lock()
	defer t.lock.Unlock()
	t.cancelToolTip()
}

func (t *ToolTipWidget) MouseMoved(e *desktop.MouseEvent) {
	t.lock.Lock()
	defer t.lock.Unlock()
	t.absoluteMousePos = e.AbsolutePosition
}

// ToolTipWidgetExtend is a struct for extending existing widgets for tool tip support.
// Use this to extend existing widgets for tool tip support. When creating an extended
// widget with ToolTipWidgetExtend you must override ExtendBaseWidget to call both the
// ExtendBaseWidget implementation of the parent widget, and ExtendToolTipWidget.
type ToolTipWidgetExtend struct {
	toolTipContext

	// Obj is the widget this ToolTipWidgetExtend is embedded in; set by ExtendToolTipWidget
	Obj fyne.CanvasObject

	toolTip string
}

func (t *ToolTipWidgetExtend) SetToolTip(toolTip string) {
	t.lock.Lock()
	defer t.lock.Unlock()
	t.toolTip = toolTip
}

func (t *ToolTipWidgetExtend) ToolTip() string {
	t.lock.Lock()
	defer t.lock.Unlock()
	return t.toolTip
}

// ExtendToolTipWidget sets up a tool tip extended widget.
func (t *ToolTipWidgetExtend) ExtendToolTipWidget(wid fyne.Widget) {
	t.Obj = wid
}

func (t *ToolTipWidgetExtend) MouseIn(e *desktop.MouseEvent) {
	t.lock.Lock()
	defer t.lock.Unlock()
	if t.toolTip != "" {
		t.absoluteMousePos = e.AbsolutePosition
		t.setPendingToolTip(t.Obj, t.toolTip)
	}
}

func (t *ToolTipWidgetExtend) MouseOut() {
	t.lock.Lock()
	defer t.lock.Unlock()
	t.cancelToolTip()
}

func (t *ToolTipWidgetExtend) MouseMoved(e *desktop.MouseEvent) {
	t.lock.Lock()
	defer t.lock.Unlock()
	t.absoluteMousePos = e.AbsolutePosition
}

func (t *toolTipContext) setPendingToolTip(wid fyne.CanvasObject, toolTipText string) {
	ctx, cancel := context.WithCancel(context.Background())
	t.pendingToolTipCtx, t.pendingToolTipCancel = ctx, cancel

	go func() {
		<-time.After(internal.NextToolTipDelayTime())
		select {
		case <-ctx.Done():
			return
		default:
			t.lock.Lock()
			defer t.lock.Unlock()
			t.cancelToolTip() // don't leak ctx resources
			pos := t.absoluteMousePos
			canvas := fyne.CurrentApp().Driver().CanvasForObject(wid)
			t.toolTipHandle = internal.ShowToolTipAtMousePosition(canvas, pos, toolTipText)
		}
	}()
}

func (t *toolTipContext) cancelToolTip() {
	if t.pendingToolTipCancel != nil {
		t.pendingToolTipCancel()
		t.pendingToolTipCancel = nil
		t.pendingToolTipCtx = nil
	}
	if t.toolTipHandle != nil {
		internal.HideToolTip(t.toolTipHandle)
		t.toolTipHandle = nil
	}
}
