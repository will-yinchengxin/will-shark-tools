package entity

type {{ .FileName }} struct {
}

func ({{ .FileNameFirstChar }} *{{ .FileName }}) TableName() string {
	return "{{ .FileNameTitleLower }}"
}