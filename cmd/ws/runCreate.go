package ws

import (
	"github.com/spf13/cobra"
	"log"
	"path/filepath"
	"strings"
	"ws/cmd/create"
)

func runCreate(cmd *cobra.Command, args []string) {
	// go run main.go create controller user
	// 这里允许传入 /controller/user/info.go
	c := create.NewCreate()
	c.ProjectName = create.GetProjectName()
	c.CreateType = cmd.Use
	c.FilePath, c.FileName = filepath.Split(args[0])
	c.FileName = strings.ReplaceAll(strings.ToUpper(string(c.FileName[0]))+c.FileName[1:], ".go", "")
	c.FileNameTitleLower = strings.ToLower(string(c.FileName[0])) + c.FileName[1:]
	c.FileNameFirstChar = string(c.FileNameTitleLower[0])

	switch c.CreateType {
	case "controller", "service", "dao", "entity", "router", "req":
		c.GenFile()
	case "all":
		c.CreateType = "controller"
		c.GenFile()

		c.CreateType = "service"
		c.GenFile()

		c.CreateType = "dao"
		c.GenFile()

		c.CreateType = "entity"
		c.GenFile()

		c.CreateType = "router"
		c.GenFile()

		c.CreateType = "req"
		c.GenFile()
	default:
		log.Fatalf("Do not Support This Type: 「%s」", c.CreateType)
	}
}
