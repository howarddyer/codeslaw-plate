package config

import (
    "html/template"
    "../models"
)

var Templates = models.Template {
    Admin: template.Must(template.ParseGlob("./templates/admin/*")),
    Public: template.Must(template.ParseGlob("./templates/public/*")),
    User: template.Must(template.ParseGlob("./templates/user/*")),
}
