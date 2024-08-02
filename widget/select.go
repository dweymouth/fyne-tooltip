package widget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
)

// Select widget has a list of options, with the current one shown, and triggers an event func when clicked
type Select struct {
	widget.Select
	ToolTipWidgetExtend
}

// NewSelect creates a new select widget with the set list of options and changes handler
func NewSelect(options []string, changed func(string)) *Select {
	s := &Select{
		Select: widget.Select{
			OnChanged:   changed,
			Options:     options,
			PlaceHolder: "(Select one)",
		},
	}
	s.ExtendBaseWidget(s)
	return s
}

func (s *Select) ExtendBaseWidget(wid fyne.Widget) {
	s.ExtendToolTipWidget(wid)
	s.Select.ExtendBaseWidget(wid)
}

func (s *Select) MouseIn(e *desktop.MouseEvent) {
	s.ToolTipWidgetExtend.MouseIn(e)
	s.Select.MouseIn(e)
}

func (s *Select) MouseMoved(e *desktop.MouseEvent) {
	s.ToolTipWidgetExtend.MouseMoved(e)
	s.Select.MouseMoved(e)
}

func (s *Select) MouseOut() {
	s.ToolTipWidgetExtend.MouseOut()
	s.Select.MouseOut()
}
