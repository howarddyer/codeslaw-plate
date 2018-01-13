package models

import "html/template"

type Template struct {
    Admin, Public, User *template.Template
}
