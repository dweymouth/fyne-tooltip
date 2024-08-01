package internal

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type ToolTip struct {
	widget.BaseWidget

	Text string
}

func NewToolTip(text string) *ToolTip {
	t := &ToolTip{Text: text}
	t.ExtendBaseWidget(t)
	return t
}

func (t *ToolTip) CreateRenderer() fyne.WidgetRenderer {
	// TODO
	return nil
}
