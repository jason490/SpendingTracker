package main

import (
	"embed"
	"log"

    "SpendingTracker/internal/app"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
    "github.com/wailsapp/wails/v2/pkg/options/linux"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

//go:embed frontend/dist
var assets embed.FS

//go:embed build/appicon.png
var icon []byte

func main() {
	// Create an instance of the app structure
	currApp := app.NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:             "SpendingTracker",
		Width:             1124,
		Height:            700,
		MinWidth:          1024,
		MinHeight:         648,
		// MaxWidth:          1280,
		// MaxHeight:         800,
		DisableResize:     false,
		Fullscreen:        false,
		Frameless:         false,
		StartHidden:       false,
		HideWindowOnClose: false,
		BackgroundColour:  &options.RGBA{R: 255, G: 255, B: 255, A: 255},
		Assets:            assets,
		Menu:              nil,
		Logger:            nil,
		LogLevel:          logger.DEBUG,
		OnStartup:         currApp.Startup,
		OnDomReady:        currApp.DomReady,
		OnBeforeClose:     currApp.BeforeClose,
		OnShutdown:        currApp.Shutdown,
		WindowStartState:  options.Normal,
		Bind: []interface{}{
			currApp,
		},
		// Windows platform specific options
		Windows: &windows.Options{
			WebviewIsTransparent: false,
			WindowIsTranslucent:  false,
			DisableWindowIcon:    false,
			// DisableFramelessWindowDecorations: false,
			WebviewUserDataPath: "",
			Theme:               windows.SystemDefault,
		},
		// Mac platform specific options
		Mac: &mac.Options{
			TitleBar: &mac.TitleBar{
				TitlebarAppearsTransparent: true,
				HideTitle:                  false,
				HideTitleBar:               false,
				FullSizeContent:            false,
				UseToolbar:                 true,
				HideToolbarSeparator:       true,
			},
			Appearance:           mac.NSAppearanceNameDarkAqua,
			WebviewIsTransparent: true,
			WindowIsTranslucent:  true,
			About: &mac.AboutInfo{
				Title:   "SpendingTracker",
				Message: "",
				Icon:    icon,
			},
		},
        Linux: &linux.Options{
            Icon: icon,
            WindowIsTranslucent: false,
            WebviewGpuPolicy: linux.WebviewGpuPolicyAlways,
            ProgramName: "wails",
        },
        Debug: options.Debug{
            OpenInspectorOnStartup: true,
        },
	})

	if err != nil {
		log.Fatal(err)
	}
}
