package main

import (
	"io/fs"
	"os"
	"os/exec"
	"path"
)

type ConfigFiles struct {
	configPath        string
	defaultConfigFile string
	files             []fs.DirEntry
}

func (cf *ConfigFiles) ReadFiles() error {

	var files []fs.DirEntry

	if !cf.VerifyDir(cf.configPath) {
		err := cf.CreateFolder(cf.configPath)
		if err != nil {
			return err
		}
	}

	cf.CreateDefaultFile(cf.defaultConfigFile)

	files, err := os.ReadDir(cf.configPath)

	if err != nil {
		return err
	}

	cf.files = files
	return nil

}

func (cf *ConfigFiles) SetConfigPath() error {
	var homeDir, err = os.UserHomeDir()
	if err != nil {
		return err
	}

	cf.configPath = path.Join(homeDir, "Paclarative")
	cf.defaultConfigFile = path.Join(cf.configPath, "pacman.conf")
	return nil
}

func (cf *ConfigFiles) VerifyDir(confiDir string) bool {
	_, err := os.Stat(confiDir)
	if err != nil {
		return false
	}
	return true
}

func (cf *ConfigFiles) CreateDefaultFile(filePath string) error {
	err := exec.Command("touch", filePath).Run()

	if err != nil {
		return err
	}

	return nil
}

func (cf *ConfigFiles) CreateFolder(configPath string) error {
	err := os.MkdirAll(configPath, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func (cf *ConfigFiles) Init() error {
	var err error
	err = cf.SetConfigPath()
	if err != nil {
		return err
	}

	err = cf.ReadFiles()
	return err
}

func CreateFileManager() *ConfigFiles {
	var config = ConfigFiles{}
	return &config
}
