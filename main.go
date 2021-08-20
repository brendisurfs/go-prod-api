package main

import (
	"fmt"
	"net/http"

	transportHTTP "github.com/brendisurfs/golang-prod-api/internal/transport/http"
)

// App - the struct which contains things like pointers
// to database connections
type App struct {
}

// extend App struct method
func (app *App) Run() error {
	fmt.Println("setting up our app.")
	handler := transportHTTP.NewHandler()
	handler.SetupRoutes()
	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		fmt.Println("failed to set up server")
		return err
	}
	return nil
}

func main() {
	fmt.Println("Go Rest API")
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error starting app.")
	}
}
