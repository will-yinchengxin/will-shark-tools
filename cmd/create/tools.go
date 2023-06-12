package create

import (
	"fmt"
	"log"
	"os"
)

func GetProjectName() string {
	file, err := os.Open("./go.mod")
	if err != nil {
		log.Fatal("Do Not Find The go.mod, Check It !! ")
		os.Exit(2)
	}
	defer file.Close()
	var (
		moduleName string
	)
	_, err = fmt.Fscanf(file, "module %s", &moduleName)
	if err != nil {
		log.Fatal("Fscanln Err: ", err.Error())
		return ""
	}
	return moduleName
}
