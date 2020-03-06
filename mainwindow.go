// main.go
package main

import (
	"flag"
	"net/http"
	"os"

	"github.com/therecipe/qt/gui"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

var window *widgets.QMainWindow
var console *widgets.QTextEdit

func runServer() {
	http.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte(mainPage))
	})

	http.ListenAndServe(":6868", nil)
}

func main() {
	var share1 = flag.String("share", "", "share path")
	var upload1 = flag.String("upload", "", "upload path")
	flag.Parse()
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	if *share1 == "" {
		setShareDir(home)
	} else {
		setShareDir(*share1)
	}
	if *upload1 == "" {
		setUploadDir(home)
	} else {
		setUploadDir(*upload1)
	}

	app := widgets.NewQApplication(len(os.Args), os.Args)
	window = widgets.NewQMainWindow(nil, core.Qt__Window)
	window.SetWindowTitle("Qt Lan Share")
	window.SetFixedHeight(400)
	createGui(window)
	app.SetActiveWindow(window)
	window.Show()
	app.Exec()
}

func createGui(parent *widgets.QMainWindow) {
	window := widgets.NewQWidget(parent, core.Qt__Widget)
	window.SetMinimumWidth(400)

	grid := widgets.NewQGridLayout(window)

	buttonShare := widgets.NewQPushButton2("Share Path", window)
	buttonShare.SetFixedWidth(100)
	grid.AddWidget2(buttonShare, 0, 0, 0)

	labelShare := widgets.NewQLabel2(share.Get(), window, core.Qt__Widget)
	grid.AddWidget2(labelShare, 0, 1, 0)

	buttonUpload := widgets.NewQPushButton2("Upload Path", window)
	buttonUpload.SetFixedWidth(100)
	grid.AddWidget2(buttonUpload, 1, 0, 0)

	labelUpload := widgets.NewQLabel2(uploadDir, window, core.Qt__Widget)
	grid.AddWidget2(labelUpload, 1, 1, 0)

	console = widgets.NewQTextEdit(window)
	console.SetReadOnly(true)
	grid.AddWidget3(console, 2, 0, 1, 2, 0)

	window.SetLayout(grid)
	parent.SetCentralWidget(window)

	window.ConnectShowEvent(func(e *gui.QShowEvent) {
		go runServer()
		showAddr()
	})

	buttonShare.ConnectClicked(func(b bool) {
		path1 := widgets.QFileDialog_GetExistingDirectory(window, "Select share path", share.Get(), widgets.QFileDialog__ShowDirsOnly)
		if len(path1) == 0 {
			return
		}
		share.Set(path1)
		labelShare.SetText(path1)
	})

	buttonUpload.ConnectClicked(func(b bool) {
		path1 := widgets.QFileDialog_GetExistingDirectory(window, "Select upload path", uploadDir, widgets.QFileDialog__ShowDirsOnly)
		if len(path1) == 0 {
			return
		}
		uploadDir = path1
		labelUpload.SetText(path1)
	})
}
