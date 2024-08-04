package widget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
)

// RichText represents the base element for a rich text-based widget.
//
// NOTE: since the tool tip RichText implements desktop.Hoverable while the
// standard RichText does not, this widget may result in hover events not
// reaching the parent Hoverable widget. It provides a callback API to allow
// parent widgets to be notified of hover events received on this widget.
type RichText struct {
	widget.RichText
	ToolTipWidgetExtend

	// Sets a callback that will be invoked for MouseIn events received
	OnMouseIn func(*desktop.MouseEvent)
	// Sets a callback that will be invoked for MouseMoved events received
	OnMouseMoved func(*desktop.MouseEvent)
	// Sets a callback that will be invoked for MouseOut events received
	OnMouseOut func()
}

// NewRichText returns a new RichText widget that renders the given text and segments.
// If no segments are specified it will be converted to a single segment using the default text settings.
func NewRichText(segments ...widget.RichTextSegment) *RichText {
	t := &RichText{RichText: widget.RichText{Segments: segments}}
	t.Scroll = container.ScrollNone
	t.ExtendBaseWidget(t)
	return t
}

// NewRichTextWithText returns a new RichText widget that renders the given text.
// The string will be converted to a single text segment using the default text settings.
func NewRichTextWithText(text string) *RichText {
	return NewRichText(&widget.TextSegment{
		Style: widget.RichTextStyleInline,
		Text:  text,
	})
}

func (r *RichText) ExtendBaseWidget(wid fyne.Widget) {
	r.ExtendToolTipWidget(wid)
	r.RichText.ExtendBaseWidget(wid)
}

func (r *RichText) MouseIn(e *desktop.MouseEvent) {
	r.ToolTipWidgetExtend.MouseIn(e)
	if f := r.OnMouseIn; f != nil {
		f(e)
	}
}

func (r *RichText) MouseMoved(e *desktop.MouseEvent) {
	r.ToolTipWidgetExtend.MouseMoved(e)
	if f := r.OnMouseMoved; f != nil {
		f(e)
	}
}

func (r *RichText) MouseOut() {
	r.ToolTipWidgetExtend.MouseOut()
	if f := r.OnMouseOut; f != nil {
		f()
	}
}
