package routes

import (
  "go-angular/server/api/sample"
  
	"fmt"
	"github.com/go-martini/martini"
)

func RegisterRoutes(app *martini.ClassicMartini) {  
  app.Group("/samples", sample.Routes)  
}
