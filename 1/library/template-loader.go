package library

import(
  "html/template"
  "path/filepath"
  "os"
  "log"
  "strings"
)

var Templates = (func() (*template.Template) {
  templ := template.New("").Funcs(template.FuncMap{
    "hello"   : hello,
  })
  err := filepath.Walk("views", func (path string, info os.FileInfo, err error) (error) {
    if strings.Contains(path, ".html") {
      _, err = templ.ParseFiles(path)
      if err != nil {
        log.Println(err)
      }
    }
    return err
  })

  if err != nil {
    panic(err)
  }

  return templ
})()

func hello() (template.HTML) {
  return "HELLO"
}
