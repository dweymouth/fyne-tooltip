# fyne-tooltip



https://github.com/user-attachments/assets/e9a2b9dd-9d6c-49f2-9c3a-b896fec4609d


**fyne-tooltip** is an add-on package for Fyne adding tooltip support. It provides a tool tip rendering system, along with drop-in tooltip-enabled extensions of several of the Fyne builtin widgets, as well as an easy means to extend existing widgets to add tool tip support, or creating new tooltip-enabled custom widgets. It aims to be easy and low-friction to integrate into your existing Fyne projects, as well as easy to remove and switch back to Fyne core if and when tooltips are supported natively.

Tool tip widgets implement `desktop.Hoverable` to show the tool tips after a short delay on MouseIn, and hide them on MouseOut.

### Enabling tool tips in a window

`fyne-tooltip` requires tool tip layers to be created for the tool tips to be rendered into, and provides APIs to do so for windows and pop-ups. It is called when settting the window's content:

```go
window.SetContent(fynetooltip.AddWindowToolTipLayer(myContent, window.Canvas()))
```

If the window is not your main window, and will be closed while your app continues to run, it is important to call `DestroyWindowToolTipLayer` to release memory resources associated with the tool tip layer.

```go
window.SetCloseIntercept(func() {
    // do any other close intercept app logic
    window.Close()
    fynetooltip.DestroyWindowToolTipLayer(window.Canvas())
})
```

### Using built-in tool tip widgets

Drop-in replacements for several built-in Fyne widgets are provided, which have been extended from the base implementation to add tool tip support.

```go
import ttwidget "github.com/dweymouth/fyne-tooltip/widget"

func main() {
    // app setup

    var btn = ttwidget.NewButton("Hello", func() {
        log.Println("Hello!")
    })
    btn.SetToolTip("world")

    // set up and show app window
}
```

### Enabling tool tips in a PopUp

Simiarly to windows, `fyne-tooltip` requires a tool tip layer to be created for pop ups to enable tool tips to be shown.
It is called after creating the pop up with content, but before showing it.

```go
pop := widget.NewPopUp(content, canvas)
fynetooltip.AddPopUpToolTipLayer(pop) // works for modal popup too
pop.Show()
```

The pop up may be hidden and re-shown. It is only necessary to create the pop up layer once. When the popup is hidden and will not be shown again,
it is important to call `DestroyPopUpToolTipLayer` to release memory resources associated with the tool tip layer.

### Creating new tool tip widgets

To create custom widgets that are tool tip enabled, just extend from the `ToolTipWidget` base struct instead of `BaseWidget`. The widget will automatically
have a SetToolTip API, and the mouse event hooks added to show and hide the tool tip.

```go
import ttwidget "github.com/dweymouth/fyne-tooltip/widget"

type MyWidget struct {
    ttwidget.ToolTipWidget
}

func NewMyWidget() *MyWidget {
    w := &MyWidget{}
    w.ExtendBaseWidget(w)
    return w
}
```


### Extending existing widgets for tool tip support

To extend existing widgets with tool tip support, use the `ToolTipWidgetExtend` struct. You must override `ExtendBaseWidget` to call both the
parent widget's `ExtendBaseWidget`, as well as `ExtendToolTipWidget`

```go
import (
    "fyne.io/fyne/v2/widget"
    ttwidget "github.com/dweymouth/fyne-tooltip/widget"
)

type ToolTipIcon struct {
    widget.Icon
    ttwidget.ToolTipWidget
}

func NewToolTipIcon(resource fyne.Resource) *ToolTipIcon {
    w := &ToolTipIcon{
        Icon: widget.Icon{
            Resource: resource,
        }
    }
    w.ExtendBaseWidget(w)
    return w
}

func (w *ToolTipIcon) ExtendBaseWidget(wid fyne.Widget) {
	w.ExtendToolTipWidget(wid)
	w.Icon.ExtendBaseWidget(wid)
}
```

If the base widget already implements `desktop.Hoverable`, you must override the Hoverable APIs to call both the parent widget's implementation,
as well as the ToolTipWidgetExtend implementation. For example, for the tooltip-enabled button:

```go
func (b *Button) MouseIn(e *desktop.MouseEvent) {
	b.ToolTipWidgetExtend.MouseIn(e)
	b.Button.MouseIn(e)
}

func (b *Button) MouseOut() {
	b.ToolTipWidgetExtend.MouseOut()
	b.Button.MouseOut()
}

func (b *Button) MouseMoved(e *desktop.MouseEvent) {
	b.ToolTipWidgetExtend.MouseMoved(e)
	b.Button.MouseMoved(e)
}
```
