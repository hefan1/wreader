//go:generate rsrc -ico icons/application.ico -manifest resource/wreader.exe.manifest -o wreader.syso
package main

import "github.com/webview/webview"

func main() {
	w := webview.New(true)
	defer w.Destroy()
	w.SetTitle("微信读书")
	w.SetSize(1920, 1200, webview.HintNone)
	w.Navigate("https://weread.qq.com/")
	w.Run()
}
