package routes

import (
  "go-angular/server/api/sample"
  
	"fmt"
	"github.com/go-martini/martini"
)

func RegisterRoutes(app *martini.ClassicMartini) {
	fmt.Println("Registering routes...")
  
  app.Group("/samples", sample.Routes)  
}
