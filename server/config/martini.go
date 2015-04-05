package config

import (
	"github.com/go-martini/martini"
	"github.com/spf13/viper"
)

func Martini(app *martini.ClassicMartini) {
  
	for _, publicPath := range viper.GetStringSlice("publicPaths") {
		app.Use(martini.Static(publicPath))
	}

}
