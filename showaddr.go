// main.go
package main

import (
	"path/filepath"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

var imgNum = 0

func showImg(parent widgets.QWidget_ITF, fn, title string) {
	dlg := widgets.NewQDialog(parent, core.Qt__Dialog)
	img := gui.NewQPixmap5(fn, "", core.Qt__NoFormatConversion)
	label := widgets.NewQLabel(dlg, core.Qt__Widget)
	label.SetPixmap(img)
	dlg.Layout().AddWidget(label)

	if title == "" {
		dlg.SetWindowTitle(filepath.Base(fn))
	} else {
		dlg.SetWindowTitle(title)
	}
	dlg.SetFixedWidth(img.Width())
	dlg.SetFixedHeight(img.Height())
	pos := window.Pos()
	pos.SetX(pos.X() + imgNum*img.Width())
	pos.SetY(pos.Y() + window.Height())
	dlg.Move(pos)
	dlg.Show()
	imgNum++
}
