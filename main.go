package main

import "fmt"

// App - the struct which contains things like pointers
// to database connections
type App struct {
}

// extend App struct method
func (app *App) Run() error {
	fmt.Println("setting up our app.")
	return nil
}

func main() {
	fmt.Println("Go Rest API")
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error starting app.")
	}
}
