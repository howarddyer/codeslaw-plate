package models

type Page struct {
    Title string
    Description string
    Type string
    Content string
}

func (page *Page) UpdateContent(content string) {
    page.Content = content
}
