package widget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type ToolbarAction struct {
	widget.ToolbarAction
	tooltip string
}

// NewToolbarActionWithIcon creates a new ToolbarAction widget with the specified label, themed icon and tap handler
func NewToolbarAction(icon fyne.Resource, onActivated func()) *ToolbarAction {
	return &ToolbarAction{
		ToolbarAction: widget.ToolbarAction{
			Icon:        icon,
			OnActivated: onActivated,
		},
	}
}

// ToolbarObject gets a button to render this ToolbarAction
func (t *ToolbarAction) ToolbarObject() fyne.CanvasObject {
	button := NewButtonWithIcon("", t.Icon, t.OnActivated)
	button.Importance = widget.LowImportance
	button.SetToolTip(t.tooltip)

	return button
}

func (t *ToolbarAction) SetToolTip(tooltip string) {
	t.tooltip = tooltip
}
