package main

import (
	"os"
)

func main() {

	var conf *ConfigFiles = CreateFileManager()
	err := conf.Init()

	if err != nil {
		handler := CreateErrorHandler(err)
		handler.ShowError()
	}


	var cmdParse CommandParser = CreateCommandParser(os.Args)
  
  var cmd *Commands = CreateCommands(&cmdParse, conf)

  cmd.ExecCommand();

}


