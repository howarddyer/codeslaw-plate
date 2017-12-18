package config

import "html/template"

var TemplatesPublic = template.Must(template.ParseGlob("./templates/public/*"))
var TemplatesAdmin = template.Must(template.ParseGlob("./templates/admin/*"))
var TemplatesUser = template.Must(template.ParseGlob("./templates/user/*"))
