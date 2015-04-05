package sample

import (
  // "fmt"
  "github.com/go-martini/martini"
)

func GetAll() string {
  return "All Samples"
}


func GetById(params martini.Params) string {
  return "Samples from id:" + params["id"]
}

