package widget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
)

type Button struct {
	widget.Button
	ToolTipWidgetExtend
}

func NewButton(text string, onTapped func()) *Button {
	return NewButtonWithIcon(text, nil, onTapped)
}

func NewButtonWithIcon(text string, icon fyne.Resource, onTapped func()) *Button {
	b := &Button{
		Button: widget.Button{
			Text:     text,
			Icon:     icon,
			OnTapped: onTapped,
		},
	}
	b.ExtendToolTipWidget(b)
	return b
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
