package tpl

import "embed"

//go:embed do/*.tpl
var CreateTemplateFS embed.FS
