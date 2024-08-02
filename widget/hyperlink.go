package widget

import (
	"net/url"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
)

// Hyperlink widget is a text component with appropriate padding and layout.
// When clicked, the default web browser should open with a URL
type Hyperlink struct {
	widget.Hyperlink
	ToolTipWidgetExtend
}

// NewHyperlink creates a new hyperlink widget with the set text content
func NewHyperlink(text string, url *url.URL) *Hyperlink {
	return NewHyperlinkWithStyle(text, url, fyne.TextAlignLeading, fyne.TextStyle{})
}

// NewHyperlinkWithStyle creates a new hyperlink widget with the set text content
func NewHyperlinkWithStyle(text string, url *url.URL, alignment fyne.TextAlign, style fyne.TextStyle) *Hyperlink {
	l := &Hyperlink{
		Hyperlink: widget.Hyperlink{
			Text:      text,
			URL:       url,
			Alignment: alignment,
			TextStyle: style,
		},
	}

	l.ExtendBaseWidget(l)
	return l
}

func (l *Hyperlink) ExtendBaseWidget(wid fyne.Widget) {
	l.ExtendToolTipWidget(wid)
	l.Hyperlink.ExtendBaseWidget(wid)
}

func (l *Hyperlink) MouseIn(e *desktop.MouseEvent) {
	l.ToolTipWidgetExtend.MouseIn(e)
	l.Hyperlink.MouseIn(e)
}

func (l *Hyperlink) MouseOut() {
	l.ToolTipWidgetExtend.MouseOut()
	l.Hyperlink.MouseOut()
}

func (l *Hyperlink) MouseMoved(e *desktop.MouseEvent) {
	l.ToolTipWidgetExtend.MouseMoved(e)
	l.Hyperlink.MouseMoved(e)
}
