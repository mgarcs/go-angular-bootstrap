package main


import (
  "go-angular/server/routes"
  "go-angular/server/config"
  "go-angular/server/config/environment"
  
  "fmt"
  "github.com/go-martini/martini"
  "github.com/spf13/viper"
)


func main() {
  app := martini.Classic();
  
  environment.ReadConfig();
  config.Martini(app);
  routes.RegisterRoutes(app);

  fmt.Println("Current Environment:", viper.GetString("env"), "Port:", viper.GetString("port"))
  
  app.Run();
}