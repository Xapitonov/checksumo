package views

import (
	"github.com/dawidd6/checksumo/src/settings"

	"github.com/dawidd6/checksumo/src/constants"
	"github.com/dawidd6/checksumo/src/utils"

	"github.com/gotk3/gotk3/gtk"
)

type settingsView struct {
	SettingsWindow         *gtk.Window
	SettingsHeaderBar      *gtk.HeaderBar
	SaveButton             *gtk.Button
	ShowNotificationsCheck *gtk.CheckButton
}

func NewSettingsView() *settingsView {
	return &settingsView{}
}

func (view *settingsView) Activate() {
	// Bind widgets
	utils.BindWidgets(view, constants.UIResourcePath)

	// Display current settings state
	view.ShowNotificationsCheck.SetActive(settings.ShowNotifications())

	// Connect handlers to events
	view.SaveButton.Connect("clicked", view.SettingsWindow.Hide)
	view.ShowNotificationsCheck.Connect("toggled", view.onShowNotificationsToggle)

	// Show settings window
	view.SettingsWindow.HideOnDelete()
	view.SettingsWindow.Present()
}

func (view *settingsView) onShowNotificationsToggle() {
	enabled := view.ShowNotificationsCheck.GetActive()
	settings.ShowNotifications(enabled)
}