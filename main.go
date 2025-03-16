package main

import (
	"os"

	"github.com/glanceapp/glance/pkg/glance"
	"github.com/go-gl/glfw/v3.3/glfw"
	webview "github.com/webview/webview_go"
)

func main() {
	go func() {
		glance.Serve(os.Getenv("GLANCE_CONFIG_PATH"))
	}()

	w := webview.New(false)
	defer w.Destroy()

	w.SetTitle("Glance Dashboard")
	width, height := screenSize()
	w.SetSize(width-100, height-100, webview.HintMin)
	w.Navigate("http://localhost:1337")

	w.Run()
}

func screenSize() (int, int) {
	err := glfw.Init()
	if err != nil {
		panic(err)
	}
	defer glfw.Terminate()

	monitor := glfw.GetPrimaryMonitor()
	vidMode := monitor.GetVideoMode()

	width := vidMode.Width
	height := vidMode.Height

	return width, height
}
