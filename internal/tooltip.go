package internal

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/dweymouth/fyne-tooltip/internal/shadow"
)

var ToolTipTextStyle = widget.RichTextStyle{SizeName: theme.SizeNameCaptionText}

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
	// zero so that ToolTip won't force a PopUp or other overlay to a larger size
	// TextMinSize returns the actual minimum size for rendering
	return fyne.NewSize(0, 0)
}

func (t *ToolTip) Resize(size fyne.Size) {
	t.updateRichText()
	t.richtext.Resize(size)
	t.BaseWidget.Resize(size)
}

func (t *ToolTip) TextMinSize() fyne.Size {
	t.updateRichText()
	return t.richtext.MinSize().Subtract(
		fyne.NewSquareSize(2 * t.Theme().Size(theme.SizeNameInnerPadding))).
		Add(fyne.NewSize(2, 8))
}

func (t *ToolTip) NonWrappingTextWidth() float32 {
	style := ToolTipTextStyle
	th := t.Theme()
	return fyne.MeasureText(t.Text, th.Size(style.SizeName), style.TextStyle).Width + th.Size(theme.SizeNameInnerPadding)*2
}

func (t *ToolTip) updateRichText() {
	if t.richtext == nil {
		t.richtext = widget.NewRichTextWithText(t.Text)
		t.richtext.Wrapping = fyne.TextWrapWord
	}
	style := ToolTipTextStyle
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
	r.toolTip.richtext.Move(fyne.NewPos(0, -innerPad+3))
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
