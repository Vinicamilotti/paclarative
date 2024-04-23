package main

import "fmt"

func main() {
	var conf *ConfigFiles = CreateFileManager()
  err := conf.Init()

 
  if err != nil {
    handler := CreateErrorHandler(err);
    handler.ShowError();
  }

  fmt.Println(conf.files)
}
