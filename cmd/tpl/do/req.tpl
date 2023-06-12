package request

type {{ .FileName }}List struct {
    Pagination
}

type {{ .FileName }} struct {
    // ApiName         string `form:"apiName" validate:"required,oneof=RunInstance StopInstance"`
    // ResourceId      string `form:"resourceId" validate:"omitempty,min=1"`
    // AppPackage      string `form:"appPackage" validate:"required,min=1"`
}