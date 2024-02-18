package main

import (
	"context"
	"fmt"
	"os/exec"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	cmd := exec.Command("sh", "-c", "cat /proc/informacion_memoria")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Sprintf("Error: %s", err.Error())
	}
	return string(out)
}
