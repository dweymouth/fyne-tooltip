package widget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

// Icon widget is a basic image component that load's its resource to match the theme.
type Icon struct {
	widget.Icon
	ToolTipWidgetExtend
}

// NewIcon returns a new icon widget that displays a themed icon resource.
func NewIcon(res fyne.Resource) *Icon {
	w := &Icon{}
	w.ExtendBaseWidget(w)
	w.SetResource(res)
	return w
}

func (w *Icon) ExtendBaseWidget(wid fyne.Widget) {
	w.ExtendToolTipWidget(wid)
	w.Icon.ExtendBaseWidget(wid)
}
