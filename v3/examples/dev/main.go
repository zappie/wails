package main

import (
	"embed"
	_ "embed"
	"log"

	"github.com/zappie/wails/v3/pkg/application"
)

//go:embed frontend/dist
var assets embed.FS

func main() {
	app := application.New(application.Options{
		Name:        "dev",
		Description: "A demo of using raw HTML & CSS",
		Assets: application.AssetOptions{
			Handler: application.AssetFileServerFS(assets),
		},
		Mac: application.MacOptions{
			ApplicationShouldTerminateAfterLastWindowClosed: true,
		},
	})
	// Create window
	app.NewWebviewWindowWithOptions(application.WebviewWindowOptions{
		Title: "Plain Bundle",
		CSS:   `body { background-color: rgba(255, 255, 255, 0); } .main { color: white; margin: 20%; }`,
		Mac: application.MacWindow{
			InvisibleTitleBarHeight: 50,
			Backdrop:                application.MacBackdropTranslucent,
			TitleBar:                application.MacTitleBarHiddenInset,
		},

		URL: "/",
	})

	//app.On(events.Common.ThemeChanged, func(e *application.Event) {
	//	if app.IsDarkMode() {
	//		log.Println("Dark mode!")
	//	} else {
	//		log.Println("Light mode!")
	//	}
	//})

	err := app.Run()

	if err != nil {
		log.Fatal(err)
	}
}
