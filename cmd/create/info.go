package create

import (
	"fmt"
	"html/template"
	"log"
	"os"
	"strings"
	"ws/cmd/tpl"
)

type Create struct {
	ProjectName        string
	CreateType         string
	FilePath           string
	FileName           string
	FileNameTitleLower string
	FileNameFirstChar  string
	IsFull             bool
}

func NewCreate() *Create {
	return &Create{}
}

func (c Create) GenFile() {
	if c.FilePath == "" {
		c.FilePath = FilePath(c.CreateType)
	} else {
		c.FilePath = FilePath(c.CreateType) + strings.TrimLeft(c.FilePath, "/")
	}
	f := c.createFile()
	if f == nil {
		log.Printf("Err: %s%s Already Exist.", c.FilePath, c.FileName)
		return
	}
	defer func() {
		_ = f.Close()
	}()
	fs, err := template.ParseFS(tpl.CreateTemplateFS, fmt.Sprintf("do/%s.tpl", c.CreateType))
	if err != nil {
		log.Fatalf("Create Tpl 「%s」 Err: %s", c.CreateType, err.Error())
		return
	}
	err = fs.Execute(f, c)
	if err != nil {
		log.Fatalf("Create Tpl 「%s」 Err: %s", c.CreateType, err.Error())
		return
	}
	log.Printf("Create Tpl 「%s-%s」 Success !! ", c.FileName, c.CreateType)
}

func (c Create) createFile() *os.File {
	var (
		filePath = c.FilePath + strings.ToLower(c.FileName+".go")
		err      error
	)
	err = os.MkdirAll(c.FilePath, os.ModePerm)
	if err != nil {
		log.Fatalf("Fail To Create Dir Err: %s", err.Error())
	}
	stat, _ := os.Stat(filePath)
	if stat != nil {
		return nil
	}
	filePb, err := os.Create(filePath)
	if err != nil {
		log.Fatalf("Failed To Create File 「%s」: %s", filePath, err.Error())
	}
	return filePb
}

func FilePath(ft string) (path string) {
	switch ft {
	case "controller":
		path = "app/controller/"
	case "service":
		path = "app/service/"
	case "dao":
		path = "app/modules/mysql/dao/"
	case "entity":
		path = "app/modules/mysql/entity/"
	case "router":
		path = "app/router/"
	case "req":
		path = "app/do/request/"
	default:
		path = "app/"
	}
	return path
}
