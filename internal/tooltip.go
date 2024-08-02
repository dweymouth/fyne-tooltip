package internal

import (
	"sync"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/dweymouth/fyne-tooltip/internal/shadow"
)

var (
	ToolTipTextStyleMutex sync.Mutex
	ToolTipTextStyle      = widget.RichTextStyle{SizeName: theme.SizeNameCaptionText}
)

type ToolTip struct {
	widget.BaseWidget

	Text string

	richtext *widget.RichText
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
	t.updateRichText()
	return t.richtext.MinSize().Subtract(
		fyne.NewSquareSize(2 * t.Theme().Size(theme.SizeNameInnerPadding))).
		Add(fyne.NewSquareSize(2))
}

func (t *ToolTip) updateRichText() {
	if t.richtext == nil {
		t.richtext = widget.NewRichTextWithText(t.Text)
	}
	ToolTipTextStyleMutex.Lock()
	style := ToolTipTextStyle
	ToolTipTextStyleMutex.Unlock()
	t.richtext.Segments[0].(*widget.TextSegment).Text = t.Text
	t.richtext.Segments[0].(*widget.TextSegment).Style = style
}

type toolTipRenderer struct {
	*shadow.ShadowingRenderer
	toolTip        *ToolTip
	backgroundRect canvas.Rectangle
}

func (r *toolTipRenderer) Layout(s fyne.Size) {
	r.LayoutShadow(s, fyne.NewPos(0, 0))
	r.backgroundRect.Resize(s)
	r.backgroundRect.Move(fyne.NewPos(0, 0))
	innerPad := r.toolTip.Theme().Size(theme.SizeNameInnerPadding)
	r.toolTip.richtext.Resize(s)
	r.toolTip.richtext.Move(fyne.NewPos(1-innerPad, -innerPad))
}

func (r *toolTipRenderer) MinSize() fyne.Size {
	return r.toolTip.TextMinSize()
}

func (r *toolTipRenderer) Refresh() {
	r.ShadowingRenderer.RefreshShadow()
	th := r.toolTip.Theme()
	variant := fyne.CurrentApp().Settings().ThemeVariant()
	r.backgroundRect.FillColor = th.Color(theme.ColorNameOverlayBackground, variant)
	r.backgroundRect.StrokeColor = th.Color(theme.ColorNameInputBorder, variant)
	r.backgroundRect.StrokeWidth = th.Size(theme.SizeNameInputBorder)
	r.backgroundRect.Refresh()
	r.toolTip.updateRichText()
	r.toolTip.richtext.Refresh()

	canvas.Refresh(r.toolTip)
}

func (t *ToolTip) CreateRenderer() fyne.WidgetRenderer {
	t.updateRichText()
	r := &toolTipRenderer{toolTip: t}
	r.ShadowingRenderer = shadow.NewShadowingRenderer([]fyne.CanvasObject{&r.backgroundRect, t.richtext}, shadow.ToolTipLevel)
	return r
}
