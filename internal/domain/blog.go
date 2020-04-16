package domain

import "html/template"

type MyFile struct {
	FileName string
	Number   string
	Date     string
	Title    string
}

type BlogData struct {
	PageTitle string
	Blogs     []MyFile
	Content   template.HTML
}
