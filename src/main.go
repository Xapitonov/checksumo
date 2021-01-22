package main

//// This header is generated by glib-compile-resources in Makefile
// #cgo pkg-config: gio-2.0
// #include "resources.h"
import "C"

import (
	"os"
	"unsafe"

	"github.com/dawidd6/checksumo/src/constants"

	"github.com/dawidd6/checksumo/src/views"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

func main() {
	// Initialize i18n
	glib.InitI18n(constants.LocaleDomain, constants.LocaleDir)

	// Set program name to app ID,
	// so DEs could recognize the app and desktop file
	//
	// This should be upstreamed to gotk3
	appIDc := (*C.gchar)(C.CString(constants.AppID))
	C.g_set_prgname(appIDc)
	C.free(unsafe.Pointer(appIDc))

	// Create app
	app, _ := gtk.ApplicationNew(constants.AppID, glib.APPLICATION_FLAGS_NONE)
	app.Connect("activate", views.NewMainView().Activate)

	// Run app
	os.Exit(app.Run(os.Args))
}
