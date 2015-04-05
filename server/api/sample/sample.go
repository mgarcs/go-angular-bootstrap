package sample

import (  
  "fmt"
  "github.com/go-martini/martini"
)

func Routes(router martini.Router) {
  
  fmt.Println("Registering Sample...")
  
  router.Get("", GetAll);
  router.Get("/:id", GetById)
  
}
