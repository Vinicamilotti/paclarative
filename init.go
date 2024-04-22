package main

import (
	"io/fs"
	"os"
	"path"
)

type ConfigFiles struct {
  files []fs.DirEntry
}

func (cf *ConfigFiles) ReadFiles() error {
   var homedir, err = os.UserHomeDir();
    if err != nil {
      return err
    }
    
    var configPath = path.Join(homedir, ".config", "Paclarative")
    var files []fs.DirEntry; 

    files, err = os.ReadDir(configPath)

    cf.files = files
    return nil
    
}

func InitFiles() *ConfigFiles {
  var config = ConfigFiles{}
  config.ReadFiles()
  return &config
}
