package main

import (
	"io/fs"
	"os"
	"os/exec"
	"path"
)

type ConfigFiles struct {
	files []fs.DirEntry
}

func (cf *ConfigFiles) ReadFiles() error {
	var homedir, err = os.UserHomeDir()
	if err != nil {
		return err
	}

	var configPath = path.Join(homedir, ".config", "Paclarative")
	var files []fs.DirEntry

	var initialConfig string = path.Join(configPath, "pacman.conf")
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		err = cf.CreateFiles(configPath)
		if err != nil {
			return err
		}
	}

	err = exec.Command("touch", initialConfig).Run()

	if err != nil {
		return err
	}

	files, err = os.ReadDir(configPath)

  if err != nil {
    return err
  }

	cf.files = files
	return nil

}

func (cf *ConfigFiles) CreateFiles(configPath string) error {
	err := os.MkdirAll(configPath, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func InitFiles() *ConfigFiles {
	var config = ConfigFiles{}
	err := config.ReadFiles()
	if err != nil {
		panic(err)
	}
	return &config
}
