package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"{{ .ProjectName }}/app/do/request"
	"{{ .ProjectName }}/app/service"
	"{{ .ProjectName }}/utils/validator"
	"{{ .ProjectName }}/utils"
)

type {{ .FileName }} struct {
	RequestValidate *validator.ValidatorX
	{{ .FileName }}            service.{{ .FileName }}
}

func ({{ .FileNameFirstChar }} *{{ .FileName }}) Index(ctx *gin.Context) {
    var (
        req           request.{{ .FileName }}
    	spanFatherCtx context.Context
    )
    if errMsg := {{ .FileNameFirstChar }}.RequestValidate.ParseQuery(ctx, &req); errMsg != "" {
        utils.MessageError(ctx, errMsg)
        return
    }
    spanFatherCtx = context.Background()
    res, codeType := {{ .FileNameFirstChar }}.{{ .FileName }}.Index(spanFatherCtx, req)
    if codeType.Code != 0 {
        utils.Error(ctx, codeType)
        return
    }
    utils.Out(ctx, res)
    return
}