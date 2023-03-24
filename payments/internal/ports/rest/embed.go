package rest

import "embed"

//go:embed index.html
//go:embed api.swagger.json
var SwaggerUI embed.FS
