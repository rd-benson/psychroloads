package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/rdbenson/psychroloads/backend/services"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

var (
	libdir, _  = os.UserConfigDir()
	basedir    = filepath.Join(libdir, "psychrolib")
	projectdir = filepath.Join(basedir, "projects")
)

// App struct
type App struct {
	ctx               context.Context
	FileSystemService *services.FileSystemService
	EPWService        *services.EPWService
}

// NewApp creates a new App application struct
// NewApp App
func NewApp() *App {
	return &App{
		FileSystemService: services.NewFileSystemService(),
		EPWService:        services.NewEPWService(),
	}
}

// startup is called at application startup
// startup
func (app *App) startup(ctx context.Context) {
	// Perform your setup here
	app.ctx = ctx
	app.FileSystemService.Ctx = ctx
	app.EPWService.Ctx = ctx
}

// domReady is called after the front-end dom has been loaded
// domReady
func (a *App) domReady(ctx context.Context) {
	// Add your action here
}

// beforeClose is called when the application is about to quit,
// either by clicking the window close button or calling runtime.Quit.
// Returning true will cause the application to continue,
// false will continue shutdown as normal.
// beforeCloseruntime.Quit
// true ，false。
func (a *App) beforeClose(ctx context.Context) (prevent bool) {
	return false
}

// shutdown is called at application termination
func (a *App) shutdown(ctx context.Context) {
	// Perform your teardown here
}

func (app *App) Title() string {
	return "psychroloads"
}

func (app *App) OpenDirectoryDialog(title string) string {
	path, _ := runtime.OpenDirectoryDialog(app.ctx, runtime.OpenDialogOptions{
		Title:                      title,
		CanCreateDirectories:       true,
		TreatPackagesAsDirectories: true,
	})

	return path
}

func (app *App) OpenFileDialog(title string) string {
	path, _ := runtime.OpenFileDialog(app.ctx, runtime.OpenDialogOptions{
		Title: title,
		Filters: []runtime.FileFilter{
			{
				DisplayName: "Weather files (*.epw)",
				Pattern:     "*.epw;",
			},
		},
	})

	return path
}

func (app *App) SaveFileDialog() string {
	path, _ := runtime.SaveFileDialog(app.ctx, runtime.SaveDialogOptions{})

	return path
}

func (app *App) SaveFile(file string, data string) bool {
	path := fmt.Sprintf("%s%s", projectdir, file)

	err := os.WriteFile(path, []byte(data), os.ModePerm)

	return err == nil
}
