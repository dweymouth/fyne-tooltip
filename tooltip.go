package fynetooltip

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/dweymouth/fyne-tooltip/internal"
)

// AddWindowToolTipLayer adds a layer to the given window content for tool tips to be drawn into.
// This call is required for each new window you create to enable tool tips to be shown in it.
// It is typically invoked with `window.SetContent` as
//
//	`window.SetContent(fynetooltip.AddWindowToolTipLayer(myContent, window.Canvas()))``
//
// If the window not your main window and is closed before the app exits, it is important to call
// `DestroyWindowToolTipLayer` to release memory resources associated with the tool tip layer.
func AddWindowToolTipLayer(windowContent fyne.CanvasObject, canvas fyne.Canvas) fyne.CanvasObject {
	return container.NewStack(windowContent, &internal.NewToolTipLayer(canvas).Container)
}

// DestroyWindowToolTipLayer destroys the tool tip layer for a given window canvas.
// It should be called after a window is closed to free associated memory resources.
func DestroyWindowToolTipLayer(canvas fyne.Canvas) {
	internal.DestroyToolTipLayerForCanvas(canvas)
}

// AddPopUpToolTipLayer adds a layer to the given `*widget.PopUp` for tool tips to be drawn into.
// This call is required for each new PopUp you create to enable tool tips to be shown in it.
// It is invoked after the PopUp has been created with content, but before it is shown.
// Once the pop up is hidden and will not be shown anymore, it is important to call
// `DestroyPopUpToolTipLayer` to release memory resources associated with the tool tip layer.
// A pop up that will be shown again should not have DestroyPopUpToolTipLayer called.
func AddPopUpToolTipLayer(p *widget.PopUp) {
	l := internal.NewPopUpToolTipLayer(p)
	p.Content = container.NewStack(p.Content, &l.Container)
}

// DestroyPopUpToolTipLayer destroys the tool tip layer for a given pop up.
// It should be called after the pop up is hidden and will no longer be shown
// to free associated memory resources.
func DestroyPopUpToolTipLayer(p *widget.PopUp) {
	internal.DestroyToolTipLayerForPopup(p)
}

// SetToolTipTextStyle sets the TextStyle that will be used to render tool tip text.
func SetToolTipTextStyle(style fyne.TextStyle) {
	internal.ToolTipTextStyleMutex.Lock()
	defer internal.ToolTipTextStyleMutex.Unlock()
	internal.ToolTipTextStyle.TextStyle = style
}

// SetToolTipTextSizeName sets the theme size name that will control the size
// of tool tip text. By default, tool tips use theme.SizeNameCaptionText.
func SetToolTipTextSizeName(sizeName fyne.ThemeSizeName) {
	internal.ToolTipTextStyleMutex.Lock()
	defer internal.ToolTipTextStyleMutex.Unlock()
	internal.ToolTipTextStyle.SizeName = sizeName
}
