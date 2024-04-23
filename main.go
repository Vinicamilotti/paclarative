package main

import "fmt"

func main() {
	var conf ConfigFiles = ConfigFiles{}

	conf.ReadFiles()

  fmt.Println(conf.files)
}
