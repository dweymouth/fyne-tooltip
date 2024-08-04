package widget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
)

// Label widget is a label component with appropriate padding and layout.
//
// NOTE: since the tool tip label implements desktop.Hoverable while the
// standard Label does not, this widget may result in hover events not
// reaching the parent Hoverable widget. It provides a callback API to allow
// parent widgets to be notified of hover events received on this widget.
type Label struct {
	widget.Label
	ToolTipWidgetExtend

	// Sets a callback that will be invoked for MouseIn events received
	OnMouseIn func(*desktop.MouseEvent)
	// Sets a callback that will be invoked for MouseMoved events received
	OnMouseMoved func(*desktop.MouseEvent)
	// Sets a callback that will be invoked for MouseOut events received
	OnMouseOut func()
}

// NewLabel creates a new label widget with the set text content
func NewLabel(text string) *Label {
	return NewLabelWithStyle(text, fyne.TextAlignLeading, fyne.TextStyle{})
}

// NewLabelWithData returns an Label widget connected to the specified data source.
func NewLabelWithData(data binding.String) *Label {
	label := NewLabel("")
	label.Bind(data)

	return label
}

// NewLabelWithStyle creates a new label widget with the set text content
func NewLabelWithStyle(text string, alignment fyne.TextAlign, style fyne.TextStyle) *Label {
	l := &Label{
		Label: widget.Label{
			Text:      text,
			Alignment: alignment,
			TextStyle: style,
		},
	}

	l.ExtendBaseWidget(l)
	return l
}

func (l *Label) ExtendBaseWidget(wid fyne.Widget) {
	l.ExtendToolTipWidget(wid)
	l.Label.ExtendBaseWidget(wid)
}

func (l *Label) MouseIn(e *desktop.MouseEvent) {
	l.ToolTipWidgetExtend.MouseIn(e)
	if f := l.OnMouseIn; f != nil {
		f(e)
	}
}

func (l *Label) MouseMoved(e *desktop.MouseEvent) {
	l.ToolTipWidgetExtend.MouseMoved(e)
	if f := l.OnMouseMoved; f != nil {
		f(e)
	}
}

func (l *Label) MouseOut() {
	l.ToolTipWidgetExtend.MouseOut()
	if f := l.OnMouseOut; f != nil {
		f()
	}
}
