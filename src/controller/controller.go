package controller

import (
	"context"

	"github.com/dawidd6/checksumo/src/model"
	"github.com/dawidd6/checksumo/src/view"
	"github.com/gotk3/gotk3/glib"
)

type Controller struct {
	v *view.View
	m *model.Model
}

func New(v *view.View, m *model.Model) *Controller {
	return &Controller{
		v: v,
		m: m,
	}
}

func (controller *Controller) Activate() {
	controller.v.FileChooserButton.Connect("file-set", controller.SetFile)
	controller.v.HashValueEntry.Connect("changed", controller.SetHash)
	controller.v.HashValueEntry.Connect("activate", controller.StartHashing)
	controller.v.VerifyButton.Connect("clicked", controller.StartHashing)
	controller.v.CancelButton.Connect("clicked", controller.StopHashing)
}

func (controller *Controller) setFileOrHash() {
	hashType := controller.m.DetectType()
	isReady := controller.m.IsReady()

	if controller.v.MainHeaderBar.GetSubtitle() != hashType {
		controller.v.MainHeaderBar.SetSubtitle(hashType)
	}
	if controller.v.VerifyButton.GetSensitive() != isReady {
		controller.v.VerifyButton.SetSensitive(isReady)
	}
	if controller.v.HashValueEntry.GetProgressFraction() > 0 {
		controller.v.HashValueEntry.SetProgressFraction(0.0)
	}
}

func (controller *Controller) SetFile() {
	filePath := controller.v.FileChooserButton.GetFilename()
	controller.m.SetFile(filePath)
	controller.setFileOrHash()
}

func (controller *Controller) SetHash() {
	hashValue, _ := controller.v.HashValueEntry.GetText()
	controller.m.SetHash(hashValue)
	controller.setFileOrHash()
}

func (controller *Controller) StopHashing() {
	controller.m.StopHashing()
}

func (controller *Controller) StartHashing() {
	if !controller.m.IsReady() {
		return
	}

	controller.m.PrepareHashing()

	controller.v.ButtonStack.SetVisibleChild(controller.v.CancelButton)
	controller.v.FileChooserButton.SetSensitive(false)
	controller.v.HashValueEntry.SetSensitive(false)

	progressSource, _ := glib.TimeoutAdd(10, func() bool {
		controller.v.HashValueEntry.SetProgressFraction(controller.m.GetProgress())
		return true
	})

	controller.m.SetResultFunc(func(ok bool, err error) {
		glib.IdleAdd(func() {
			controller.v.HashValueEntry.SetProgressFraction(controller.m.GetProgress())

			if err == context.Canceled {
				// NOOP
			} else if err != nil {
				controller.v.OnError(err)
			} else if ok {
				controller.v.OnSuccess()
			} else {
				controller.v.OnFailure()
			}

			controller.v.ButtonStack.SetVisibleChild(controller.v.VerifyButton)
			controller.v.FileChooserButton.SetSensitive(true)
			controller.v.HashValueEntry.SetSensitive(true)
		})
		glib.SourceRemove(progressSource)
	})

	go controller.m.StartHashing()
}
