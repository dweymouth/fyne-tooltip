package widget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
)

// Slider is a widget that can slide between two fixed values.
type Slider struct {
	widget.Slider
	ToolTipWidgetExtend
}

// NewSlider returns a basic slider.
func NewSlider(min, max float64) *Slider {
	slider := &Slider{
		Slider: widget.Slider{
			Value:       0,
			Min:         min,
			Max:         max,
			Step:        1,
			Orientation: widget.Horizontal,
		},
	}
	slider.ExtendBaseWidget(slider)
	return slider
}

// NewSliderWithData returns a slider connected with the specified data source.
func NewSliderWithData(min, max float64, data binding.Float) *Slider {
	slider := NewSlider(min, max)
	slider.Bind(data)

	return slider
}

func (s *Slider) ExtendBaseWidget(wid fyne.Widget) {
	s.ExtendToolTipWidget(wid)
	s.Slider.ExtendBaseWidget(wid)
}

func (s *Slider) MouseIn(e *desktop.MouseEvent) {
	s.ToolTipWidgetExtend.MouseIn(e)
	s.Slider.MouseIn(e)
}

func (s *Slider) MouseMoved(e *desktop.MouseEvent) {
	s.ToolTipWidgetExtend.MouseMoved(e)
	s.Slider.MouseMoved(e)
}

func (s *Slider) MouseOut() {
	s.ToolTipWidgetExtend.MouseOut()
	s.Slider.MouseOut()
}
