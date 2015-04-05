package environment

import (
  "fmt"
  "path"
  "runtime"
  
  "github.com/spf13/viper"
)

func ReadConfig() {
  _, filename, _, _ := runtime.Caller(1)  
  rootPath := path.Join(path.Dir(filename), "..")  
      
  viper.AutomaticEnv()
  
  viper.SetDefault("env", "development")
  viper.SetDefault("port", 9000)

  viper.Set("rootPath", rootPath)
  
    
  viper.AddConfigPath(path.Join(rootPath, "server", "config", "environment"))  
  viper.SetConfigName(viper.GetString("env"))    

  err := viper.ReadInConfig()
  if err != nil {
    panic(fmt.Errorf("Fatal error config file: %s \n", err))
  }
  
  
  // normalize paths for relative folders
  publicPaths := viper.GetStringSlice("publicPaths");
  for i, publicPath := range publicPaths {
    if !path.IsAbs(publicPath) {
      publicPaths[i] = path.Join(rootPath, publicPath)
    }
  }
  viper.Set("publicPaths", publicPaths);  
  
}

