package main

import (
	"embed"
	"log"

	"github.com/rdbenson/psychroloads/backend/services"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/linux"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

//go:embed frontend/dist
var assets embed.FS

//go:embed build/appicon.png
var icon []byte

func main() {
	// Create an instance of the app structure
	app := NewApp()

	// Create an instance of the EPW structure
	epw := services.NewEPW()

	// Create application with options
	err := wails.Run(&options.App{
		Title:             "psychroloads",
		Width:             900,
		Height:            600,
		MinWidth:          900,
		MinHeight:         600,
		DisableResize:     false,
		Fullscreen:        false,
		Frameless:         false,
		StartHidden:       false,
		HideWindowOnClose: false,
		BackgroundColour:  &options.RGBA{R: 255, G: 255, B: 255, A: 0},
		Menu:              nil,
		Logger:            nil,
		LogLevel:          logger.DEBUG,
		OnStartup:         app.startup,
		OnDomReady:        app.domReady,
		OnBeforeClose:     app.beforeClose,
		OnShutdown:        app.shutdown,
		WindowStartState:  options.Normal,
		AssetServer: &assetserver.Options{
			Assets:     assets,
			Handler:    nil,
			Middleware: nil,
		},
		Bind: []interface{}{
			app,
			app.FileSystemService,
			epw,
		},
		// Windows platform specific options
		Windows: &windows.Options{
			WebviewIsTransparent:              true,
			WindowIsTranslucent:               false,
			DisableWindowIcon:                 false,
			DisableFramelessWindowDecorations: false,
			WebviewUserDataPath:               "",
			WebviewBrowserPath:                "",
			Theme:                             windows.SystemDefault,
		},
		// Mac platform specific options
		Mac: &mac.Options{
			TitleBar: &mac.TitleBar{
				TitlebarAppearsTransparent: true,
				HideTitle:                  true,
				HideTitleBar:               false,
				FullSizeContent:            true,
				UseToolbar:                 false,
				HideToolbarSeparator:       false,
			},
			Appearance:           mac.NSAppearanceNameDarkAqua,
			WebviewIsTransparent: true,
			WindowIsTranslucent:  true,
			About: &mac.AboutInfo{
				Title:   "Wails Template Vue",
				Message: "A Wails template based on Vue and Vue-Router",
				Icon:    icon,
			},
		},
		Linux: &linux.Options{
			Icon: icon,
		},
	})

	if err != nil {
		log.Fatal(err)
	}
}
