package config

import (
    "html/template"
    "../models"
)

var Templates = models.Template {
    Public: template.Must(template.ParseGlob("./templates/public/*")),
}
