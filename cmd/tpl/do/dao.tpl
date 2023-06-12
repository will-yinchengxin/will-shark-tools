package dao

import (
	"context"
	"{{ .ProjectName }}/app/modules/mysql/entity"
)

type {{ .FileName }} struct {
	*MysqlPool
}

func ({{ .FileNameFirstChar }} *{{ .FileName }}) WithContext(ctx context.Context) *{{ .FileName }} {
	{{ .FileNameFirstChar }}.DB = {{ .FileNameFirstChar }}.DB.WithContext(ctx)
	return {{ .FileNameFirstChar }}
}

func ({{ .FileNameFirstChar }} *{{ .FileName }}) Create(ctx context.Context, req *entity.{{ .FileName }}) error {
	return {{ .FileNameFirstChar }}.DB.Model(entity.{{ .FileName }}{}).
		Create(&req).Error
}


func ({{ .FileNameFirstChar }} *{{ .FileName }}) List(req request.{{ .FileName }}List) ([]entity.{{ .FileName }}, int64, error) {
	var (
		list  = []entity.{{ .FileName }}{}
		total int64
	)
	db := {{ .FileNameFirstChar }}.DB.Model(entity.{{ .FileName }}{})
	// TODO：Like Or Other
	//if len(req.UserId) > 0 {
	//	db.Where("user_id LIKE ?", "%"+req.UserId+"%")
	//}

	if err := db.Count(&total).
		Offset((req.Page - 1) * req.Size).
		Limit(req.Size).
		Find(&list).Error; err != nil {
		return nil, total, err
	}
	return list, total, nil
}

func ({{ .FileNameFirstChar }} *{{ .FileName }}) Count() (int64, error) {
	var (
		count int64
	)
	err := {{ .FileNameFirstChar }}.DB.Model(entity.{{ .FileName }}{}).Count(&count).Error
	return count, err
}

func ({{ .FileNameFirstChar }} *{{ .FileName }}) Delete(id int64) error {
	return {{ .FileNameFirstChar }}.DB.Model(entity.{{ .FileName }}{}).
		Where("id = ?", id).
		Delete(&entity.{{ .FileName }}{}).
		Error
}

//func ({{ .FileNameFirstChar }} *{{ .FileName }}) Update(data entity.{{ .FileName }}, t int64) error {
//	return {{ .FileNameFirstChar }}.DB.Model(entity.{{ .FileName }}{}).
//		Where("id = ?", data.Id).
//		Update("create_time", t).
//		Error
//}