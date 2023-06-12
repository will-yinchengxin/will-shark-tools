package service

import (
    "context"
    "{{ .ProjectName }}/app/do/request"
    "{{ .ProjectName }}/utils"
    "{{ .ProjectName }}/app/modules/mysql/dao"
)

type {{ .FileName }} struct {
    {{ .FileName }} *dao.{{ .FileName }}
}

func ({{ .FileNameFirstChar }} *{{ .FileName }}) Index(ctx context.Context, req request.{{ .FileName }}) (interface{}, *utils.CodeType) {
	var (
		err error
	)
	defer func() {
		utils.ErrorLog(err)
	}()

    return nil, &utils.CodeType{}
}

