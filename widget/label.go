package widget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

type Label struct {
	widget.Label
	ToolTipWidgetExtend
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
	l.Label.ExtendBaseWidget(l)
}
