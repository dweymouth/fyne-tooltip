package widget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
)

// Check widget has a text label and a checked (or unchecked) icon and triggers an event func when toggled
type Check struct {
	widget.Check
	ToolTipWidgetExtend
}

// NewCheck creates a new check widget with the set label and change handler
func NewCheck(label string, changed func(bool)) *Check {
	c := &Check{
		Check: widget.Check{
			Text:      label,
			OnChanged: changed,
		},
	}
	c.ExtendBaseWidget(c)
	return c
}

// NewCheckWithData returns a check widget connected with the specified data source.
func NewCheckWithData(label string, data binding.Bool) *Check {
	check := NewCheck(label, nil)
	check.Bind(data)

	return check
}

func (c *Check) ExtendBaseWidget(wid fyne.Widget) {
	c.ExtendToolTipWidget(wid)
	c.Check.ExtendBaseWidget(wid)
}

func (c *Check) MouseIn(e *desktop.MouseEvent) {
	c.ToolTipWidgetExtend.MouseIn(e)
	c.Check.MouseIn(e)
}

func (c *Check) MouseMoved(e *desktop.MouseEvent) {
	c.ToolTipWidgetExtend.MouseMoved(e)
	c.Check.MouseMoved(e)
}

func (c *Check) MouseOut() {
	c.ToolTipWidgetExtend.MouseOut()
	c.Check.MouseOut()
}
