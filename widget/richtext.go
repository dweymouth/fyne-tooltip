package widget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// RichText represents the base element for a rich text-based widget.
type RichText struct {
	widget.RichText
	ToolTipWidgetExtend
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
