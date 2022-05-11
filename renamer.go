package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

func main() {

	fmt.Println("Enter folder to print")
	var folder string

	fmt.Scanln(&folder)

	files, err := ioutil.ReadDir(folder)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Enter folder file extension")
	var fileExtension string
	fmt.Scanln(&fileExtension)

	for _, file := range files {
		var currFileLocation string = folder + "/" + file.Name()
		var uuid, err = exec.Command("uuidgen").Output()

		var uid = string(uuid)

		if err != nil {
			log.Fatal(err)
		}

		var newFileName string = folder + "/" + uid + "." + fileExtension
		fmt.Println("Current File location ", currFileLocation)
		fmt.Println("New file name ", newFileName)
		e := os.Rename(currFileLocation, newFileName)
		if e != nil {
			log.Fatal(e)
		}

	}

}
