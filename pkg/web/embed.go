package web

import (
	"embed"
)

////go:embed index.html
//go:embed swagger-ui/*
var WebUI embed.FS
