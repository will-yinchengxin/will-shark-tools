package router

import (
	"github.com/gin-gonic/gin"
	"{{ .ProjectName }}/app/controller"
)

type {{ .FileName }}Router struct {
	{{ .FileName }}Api *controller.{{ .FileName }}
}

func ({{ .FileNameFirstChar }} *{{ .FileName }}Router) Initialize(app *gin.Engine) {
	v1 := app.Group("api")
	{
		apps := v1.Group("v1")
		{
            apps.GET("/index", {{ .FileNameFirstChar }}.{{ .FileName }}Api.Index)
		}
	}
}