package internal

import (
	"sync"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var (
	ToolTipTextStyleMutex sync.Mutex
	ToolTipTextStyle      = widget.RichTextStyle{SizeName: theme.SizeNameCaptionText}
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

func (t *ToolTip) MinSize() fyne.Size {
	return fyne.NewSize(0, 0)
}

func (t *ToolTip) TextMinSize() fyne.Size {
	return t.BaseWidget.MinSize()
}

func (t *ToolTip) CreateRenderer() fyne.WidgetRenderer {
	variant := fyne.CurrentApp().Settings().ThemeVariant()
	ToolTipTextStyleMutex.Lock()
	style := ToolTipTextStyle
	ToolTipTextStyleMutex.Unlock()
	rt := widget.NewRichTextWithText(t.Text)
	rt.Segments[0].(*widget.TextSegment).Style = style
	return widget.NewSimpleRenderer(
		container.NewStack(
			canvas.NewRectangle(t.Theme().Color(theme.ColorNameOverlayBackground, variant)),
			rt,
		),
	)
}
