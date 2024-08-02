package widget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

// ProgressBar widget creates a horizontal panel that indicates progress
type ProgressBar struct {
	widget.ProgressBar
	ToolTipWidgetExtend
}

// NewProgressBar creates a new progress bar widget.
// The default Min is 0 and Max is 1, Values set should be between those numbers.
// The display will convert this to a percentage.
func NewProgressBar() *ProgressBar {
	bar := &ProgressBar{ProgressBar: widget.ProgressBar{Min: 0, Max: 1}}
	bar.ExtendBaseWidget(bar)
	return bar
}

// NewProgressBarWithData returns a progress bar connected with the specified data source.
func NewProgressBarWithData(data binding.Float) *ProgressBar {
	p := NewProgressBar()
	p.Bind(data)
	return p
}

func (p *ProgressBar) ExtendBaseWidget(wid fyne.Widget) {
	p.ExtendToolTipWidget(wid)
	p.ProgressBar.ExtendBaseWidget(wid)
}

// ProgressBarInfinite widget creates a horizontal panel that indicates waiting indefinitely
// An infinite progress bar loops 0% -> 100% repeatedly until Stop() is called
type ProgressBarInfinite struct {
	widget.ProgressBarInfinite
	ToolTipWidgetExtend
}

// NewProgressBarInfinite creates a new progress bar widget that loops indefinitely from 0% -> 100%
// SetValue() is not defined for infinite progress bar
// To stop the looping progress and set the progress bar to 100%, call ProgressBarInfinite.Stop()
func NewProgressBarInfinite() *ProgressBarInfinite {
	bar := &ProgressBarInfinite{}
	bar.ExtendBaseWidget(bar)
	return bar
}

func (p *ProgressBarInfinite) ExtendBaseWidget(wid fyne.Widget) {
	p.ExtendToolTipWidget(wid)
	p.ProgressBarInfinite.ExtendBaseWidget(wid)
}
