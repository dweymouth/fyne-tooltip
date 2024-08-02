package widget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
)

// Button widget has a text label and triggers an event func when clicked
type Button struct {
	widget.Button
	ToolTipWidgetExtend
}

// NewButton creates a new button widget with the set label and tap handler
func NewButton(text string, onTapped func()) *Button {
	return NewButtonWithIcon(text, nil, onTapped)
}

// NewButtonWithIcon creates a new button widget with the specified label, themed icon and tap handler
func NewButtonWithIcon(text string, icon fyne.Resource, onTapped func()) *Button {
	b := &Button{
		Button: widget.Button{
			Text:     text,
			Icon:     icon,
			OnTapped: onTapped,
		},
	}
	b.ExtendBaseWidget(b)
	return b
}

func (b *Button) ExtendBaseWidget(wid fyne.Widget) {
	b.ExtendToolTipWidget(wid)
	b.Button.ExtendBaseWidget(wid)
}

func (b *Button) MouseIn(e *desktop.MouseEvent) {
	b.ToolTipWidgetExtend.MouseIn(e)
	b.Button.MouseIn(e)
}

func (b *Button) MouseOut() {
	b.ToolTipWidgetExtend.MouseOut()
	b.Button.MouseOut()
}

func (b *Button) MouseMoved(e *desktop.MouseEvent) {
	b.ToolTipWidgetExtend.MouseMoved(e)
	b.Button.MouseMoved(e)
}
